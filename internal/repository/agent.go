package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/michaelhendraw/stockbit_microservice/internal/entity"
)

type (
	// AgentResponse struct
	AgentResponse struct {
		Search []struct {
			Title  string `json:"Title"`
			Year   string `json:"Year"`
			ImdbID string `json:"imdbID"`
			Type   string `json:"Type"`
			Poster string `json:"Poster"`
		} `json:"Search"`
		TotalResults string `json:"totalResults"`
		Error        string `json:"Error"`
		Response     string `json:"Response"`
	}
)

// NewAgent service
func NewAgent(endpoint string, apiKey string) Agent {
	rp := &agentRepo{
		Endpoint: endpoint,
		APIKey:   apiKey,
	}
	return rp
}

// Search func
func (r *agentRepo) Search(searchWord string, pagination int64) (entity.SearchResponse, error) {
	url := fmt.Sprintf("%s/?apikey=%s&s=%s&page=%d", r.Endpoint, r.APIKey, searchWord, pagination)

	result := entity.SearchResponse{}
	agentResponse, err := agentAPICall(url, "GET", nil, nil)
	if err != nil {
		return result, err
	}

	for _, v := range agentResponse.Search {
		result.SearchResponseData = append(result.SearchResponseData, entity.SearchResponseData{
			Title:  v.Title,
			Year:   v.Year,
			ImdbID: v.ImdbID,
			Type:   v.Type,
			Poster: v.Poster,
		})
	}

	result.TotalResults = agentResponse.TotalResults

	return result, nil
}

// agentAPICall func
func agentAPICall(url string, method string, request []byte, additionalHeaders map[string]string) (AgentResponse, error) {
	result := AgentResponse{}

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(request))
	req.Close = true

	for k, v := range additionalHeaders {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("agent repository func agentAPICall error client Do, request: %v, err: %s\n", req, err)
		return result, errors.New("Could not connect to agent")
	}

	defer resp.Body.Close()
	contents, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(contents, &result)
	if err != nil {
		log.Printf("agent repository func agentAPICall error Unmarshal, request: %v, err: %s\n", req, err)
		return result, err
	}
	if len(result.Error) > 0 {
		err = errors.New(result.Error)
		log.Printf("agent repository func agentAPICall error response, request: %v, err: %s\n", req, result.Error)
		return result, err
	}

	return result, err
}
