package rest

import (
	"github.com/tebben/marvin/go/marvin/models"
)

// CreateEndPoints creates the pre-defined endpoint config, the config contains all endpoint info
func CreateEndPoints() []models.Endpoint {
	endpoints := []models.Endpoint{
		createModules(),
		//triggers, config
	}

	return endpoints
}

func createModules() *Endpoint {
	return &Endpoint{
		Name: "Modules",
		Operations: []models.EndpointOperation{
			{models.HTTPOperationGet, "/Modules", HandleGetModules},
		},
	}
}