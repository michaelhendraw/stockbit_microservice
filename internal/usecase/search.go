package usecase

import (
	"github.com/michaelhendraw/stockbit_microservice/internal/entity"
	"github.com/michaelhendraw/stockbit_microservice/internal/service"
)

// NewSearch usecase
func NewSearch(
	searchSvc service.Search,
) Search {
	return &searchUsecase{
		searchSvc: searchSvc,
	}
}

// Search func
func (u *searchUsecase) Search(searchRequest entity.SearchRequest) (entity.SearchResponse, error) {
	return u.searchSvc.Search(searchRequest)
}
