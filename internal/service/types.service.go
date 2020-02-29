package service

import (
	"github.com/michaelhendraw/stockbit_microservice/internal/entity"
	"github.com/michaelhendraw/stockbit_microservice/internal/repository"
)

// Service Module
type (
	// searchSvc struct
	searchSvc struct {
		agentRp repository.Agent
		// searchRp repository.Search
	}
)

// Service interfaces
type (
	// Search interface
	Search interface {
		Search(searchRequest entity.SearchRequest) (entity.SearchResponse, error)
	}
)
