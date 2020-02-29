package service

import (
	"github.com/michaelhendraw/stockbit_microservice/internal/entity"
	"github.com/michaelhendraw/stockbit_microservice/internal/repository"
)

// NewSearch service
func NewSearch(
	agentRp repository.Agent,
	// searchRp repository.Search,
) Search {
	return &searchSvc{
		agentRp: agentRp,
	}
}

// Search func
func (u *searchSvc) Search(searchRequest entity.SearchRequest) (entity.SearchResponse, error) {
	return u.agentRp.Search(searchRequest.SearchWord, searchRequest.Pagination)
}
