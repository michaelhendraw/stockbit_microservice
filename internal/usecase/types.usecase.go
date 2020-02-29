package usecase

import (
	"github.com/michaelhendraw/stockbit_microservice/internal/entity"
	"github.com/michaelhendraw/stockbit_microservice/internal/service"
)

// Usecase module
type (
	// searchUsecase struct
	searchUsecase struct {
		searchSvc service.Search
	}
)

// Usecase interfaces
type (
	// Search interface
	Search interface {
		Search(searchRequest entity.SearchRequest) (entity.SearchResponse, error)
	}
)
