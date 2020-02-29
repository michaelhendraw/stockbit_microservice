package grpc

import (
	"errors"

	"github.com/michaelhendraw/stockbit_microservice/internal/entity"
	"github.com/michaelhendraw/stockbit_microservice/internal/usecase"
)

// SearchGRPC struct
type SearchGRPC struct {
	searchUC usecase.Search
}

// NewSearchGRPC ...
func NewSearchGRPC(searchUseCase usecase.Search) *SearchGRPC {
	return &SearchGRPC{
		searchUC: searchUseCase,
	}
}

// Search ...
func (d *SearchGRPC) Search(pbRequest *SearchRequest) (*SearchResponse, error) {
	pbResponse := &SearchResponse{}

	if pbRequest == nil {
		pbResponse.Error = "Request should not be nil"
		return pbResponse, errors.New("Search request should not be nil")
	}

	var searchWord string
	var pagination int64
	var err error

	searchWord, pagination, pbResponse.Error, err = d.getAndValidateSearchParam(pbRequest)
	if err != nil {
		return pbResponse, err
	}

	searchRequest := entity.SearchRequest{
		SearchWord: searchWord,
		Pagination: pagination,
	}

	searchResponse, err := d.searchUC.Search(searchRequest)
	if err != nil {
		return pbResponse, err
	}

	pbResponse = d.mapToPbResponse(searchResponse)

	return pbResponse, nil
}

func (d *SearchGRPC) mapToPbResponse(searchResponse entity.SearchResponse) *SearchResponse {
	var pbResponse SearchResponse

	var pbResponseDatas []*SearchResponseData
	for _, data := range searchResponse.SearchResponseData {
		pbResponseDatas = append(pbResponseDatas, &SearchResponseData{
			Title:  data.Title,
			Year:   data.Year,
			ImdbID: data.ImdbID,
			Type:   data.Type,
			Poster: data.Poster,
		})
	}

	pbResponse.TotalResults = searchResponse.TotalResults
	pbResponse.Error = searchResponse.Error

	return &pbResponse
}

func (d *SearchGRPC) getAndValidateSearchParam(pbRequest *SearchRequest) (string, int64, string, error) {
	var searchWord string
	var pagination int64
	var searchResponseError string

	if pbRequest.SearchWord == "" {
		searchResponseError = "Empty search word"
		return searchWord, pagination, searchResponseError, errors.New(searchResponseError)
	}

	if pbRequest.Pagination <= 0 {
		searchResponseError = "Invalid pagination"
		return searchWord, pagination, searchResponseError, errors.New(searchResponseError)
	}

	searchWord = pbRequest.SearchWord
	pagination = pbRequest.Pagination

	return searchWord, pagination, searchResponseError, nil
}
