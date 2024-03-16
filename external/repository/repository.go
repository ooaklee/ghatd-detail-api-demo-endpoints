package repository

import (
	"github.com/ooaklee/ghatd-detail-api-demo-endpoints/external/dictionary"
)

// InMememoryRepository holds the methods for managing words
type InMememoryRepository struct {
	// Store holds "persistent" words
	Store []dictionary.Word
}

// NewInMememoryRepository returns a new in-memory repository
func NewInMememoryRepository() *InMememoryRepository {

	return &InMememoryRepository{
		Store: []dictionary.Word{
			{
				Id:        "8b0abdfa-1eaa-48c6-a7b6-e0011c195d67",
				Name:      "fire truck",
				CreatedAt: "2024-02-29T04:35:13.612482",
			},
			{
				Id:        "6b0abdfa-1eaa-48c6-a7b6-e0011c195d12",
				Name:      "police car",
				CreatedAt: "2024-01-29T04:35:13.612482",
			},
			{
				Id:        "1b2abdfa-1eaa-38c6-a7b7-c1011c195a10",
				Name:      "ambulance",
				CreatedAt: "2024-02-14T00:35:13.612482",
			},
		},
	}
}
