package internal

import (
	"github.com/michaelhendraw/stockbit_microservice/internal/repository"
	"github.com/michaelhendraw/stockbit_microservice/internal/service"
	"github.com/michaelhendraw/stockbit_microservice/internal/usecase"
)

// GetUsecase func
func GetUsecase(config *Config) *Usecase {

	// REPOSITORY
	agentRepo := repository.NewAgent(config.Agent.Endpoint, config.Agent.APIKey)

	// SERVICE
	searchService := service.NewSearch(agentRepo)

	// USECASE
	searchUsecase := usecase.NewSearch(searchService)

	ucase := &Usecase{
		Search: searchUsecase,
	}

	return ucase
}
