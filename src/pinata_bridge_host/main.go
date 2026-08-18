package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/6022-labs/agentic-pinata-bridge/src/common/host_configurations"
	common_settings "github.com/6022-labs/agentic-pinata-bridge/src/common/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_host/configurations"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_host/settings"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	_ = godotenv.Load(".env")

	config, err := host_configurations.LoadKoanfConfig()
	if err != nil {
		panic(err)
	}

	container := configurations.ConfigureDI(config)
	host_configurations.ConfigureLogging(container)

	shutdownTelemetry := host_configurations.ConfigureTelemetry(container, configurations.AppName)

	configurations.ConfigureServer(container)

	err = container.Invoke(func(logger *zap.Logger, hostSettings *settings.HostFeaturesSettings) {
		logger.Info("Starting Pinata Bridge Host")

		if !hostSettings.UseApi && !hostSettings.UseListeners {
			logger.Info(
				"Both API and listeners are disabled. Waiting 10 seconds before exiting to spare the RPC nodes.",
			)
			time.Sleep(10 * time.Second)
		}
	})
	if err != nil {
		panic(err)
	}

	err = container.Invoke(func(logger *zap.Logger, pushMissingImageCids *use_cases.PushMissingImageCids) error {
		logger.Info("Pushing missing image cids...")
		_, err := pushMissingImageCids.Execute(context.Background())
		return err
	})
	if err != nil {
		panic(err)
	}

	err = container.Invoke(func(
		app *fiber.App,
		logger *zap.Logger,
		hostSettings *settings.HostFeaturesSettings,
		commonHostSettings *common_settings.HostSettings,
	) error {
		if !hostSettings.UseApi && !hostSettings.UseListeners {
			return nil
		}

		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

		// Listener-only mode has no server to wait on; the signal is the only thing that ends the process.
		serverErrCh := make(chan error, 1)
		if hostSettings.UseApi {
			bind := fmt.Sprintf("%s:%d", commonHostSettings.ListenAddress, commonHostSettings.ApiPort)
			logger.Info("Starting server", zap.String("bind", bind))

			go func() {
				serverErrCh <- app.Listen(bind)
			}()
		} else {
			logger.Info("Listeners are enabled, but API is disabled. Waiting for a shutdown signal.")
		}

		select {
		case sig := <-sigCh:
			logger.Info("Shutdown signal received", zap.String("signal", sig.String()))

			if hostSettings.UseApi {
				if err := app.Shutdown(); err != nil {
					logger.Warn("Fiber shutdown returned an error", zap.Error(err))
				}

				if err := <-serverErrCh; err != nil {
					logger.Error("Server stopped with error", zap.Error(err))
				} else {
					logger.Info("Server stopped")
				}
			}
			if err := shutdownTelemetry(context.Background()); err != nil {
				logger.Warn("Telemetry shutdown returned an error", zap.Error(err))
			}

			logger.Info("Shutdown complete")
			_ = logger.Sync()
			return nil

		case err := <-serverErrCh:
			if err != nil {
				return fmt.Errorf("server exited unexpectedly: %w", err)
			}
			return nil
		}
	})
	if err != nil {
		panic(err)
	}
}
