package services

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

// subscribeToCollectionEvent opens one log subscription covering every collection address at once,
// so a chain costs a single filter instead of one per collection. It is a package-level func because
// Go methods cannot carry their own type parameter.
func subscribeToCollectionEvent[T any](
	ctx context.Context,
	client *ethclient.Client,
	contractAbi *abi.ABI,
	eventName string,
	addresses []common.Address,
	parse func(types.Log) (*T, error),
) (<-chan *T, ethereum.Subscription, error) {
	eventDef, ok := contractAbi.Events[eventName]
	if !ok {
		return nil, nil, fmt.Errorf("event %s not found in the contract ABI", eventName)
	}

	query := ethereum.FilterQuery{Topics: [][]common.Hash{{eventDef.ID}}, Addresses: addresses}

	rawLogs := make(chan types.Log)
	subscription, err := client.SubscribeFilterLogs(ctx, query, rawLogs)
	if err != nil {
		return nil, nil, err
	}

	parsedLogs := make(chan *T, 64)

	wrapped := event.NewSubscription(func(quit <-chan struct{}) error {
		defer close(parsedLogs)
		defer subscription.Unsubscribe()

		for {
			select {
			case rawLog := <-rawLogs:
				parsed, parseErr := parse(rawLog)
				if parseErr != nil {
					return fmt.Errorf("failed to parse %s log: %w", eventName, parseErr)
				}

				select {
				case parsedLogs <- parsed:
				case subscriptionErr := <-subscription.Err():
					return subscriptionErr
				case <-quit:
					return nil
				}
			case subscriptionErr := <-subscription.Err():
				return subscriptionErr
			case <-quit:
				return nil
			}
		}
	})

	return parsedLogs, wrapped, nil
}
