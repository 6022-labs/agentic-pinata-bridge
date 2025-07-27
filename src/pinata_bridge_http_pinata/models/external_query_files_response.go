package models

import (
	"time"

	"github.com/google/uuid"
)

type ExternalQueryFilesResponse struct {
	Data struct {
		Files []struct {
			ID            uuid.UUID  `json:"id"`
			Name          *string    `json:"name"`
			Cid           string     `json:"cid"`
			Size          int        `json:"size"`
			NumberOfFiles int        `json:"number_of_files"`
			MimeType      string     `json:"mime_type"`
			GroupID       *uuid.UUID `json:"group_id"`
			CreatedAt     time.Time  `json:"created_at"`
		} `json:"files"`
		NextPageToken *string `json:"next_page_token"`
	} `json:"data"`
}
