package entity

type (
	// SearchRequest struct
	SearchRequest struct {
		SearchWord string
		Pagination int64
	}

	// SearchResponse struct
	SearchResponse struct {
		SearchResponseData []SearchResponseData `json:"search,omitempty"`
		TotalResults       string               `json:"totalResults,omitempty"`
		Error              string               `json:"error,omitempty"`
	}

	// SearchResponseData struct
	SearchResponseData struct {
		Title  string `json:"title,omitempty"`
		Year   string `json:"year,omitempty"`
		ImdbID string `json:"imdb_id,omitempty"`
		Type   string `json:"type,omitempty"`
		Poster string `json:"poster,omitempty"`
	}
)
