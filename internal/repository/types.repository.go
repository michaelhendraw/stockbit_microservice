package repository

import "github.com/michaelhendraw/stockbit_microservice/internal/entity"

// Repository Module
type (
	// agentRepo struct
	agentRepo struct {
		Endpoint string
		APIKey   string
	}
	// searchRepo struct
	searchRepo struct {
	}
)

// Repository interfaces
type (
	// Agent interface
	Agent interface {
		Search(searchWord string, pagination int64) (entity.SearchResponse, error)
	}

	// Search interface
	Search interface {
	}
)
