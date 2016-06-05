package system

import (
	"github.com/tebben/marvin/go/events"
	"github.com/tebben/marvin/go/marvin/models"
	"github.com/tebben/marvin/go/marvin/rest"
)

type Marvin struct {
	modules       []models.MarvinModule
	restEndpoints []models.MarvinEndpoint
}

// NewAPI Initialise a new SensorThings API
func CreateMarvin() models.Marvin {
	return &Marvin{}
}

// AddModule add a new module to Marvin
func (m *Marvin) AddModule(module models.MarvinModule) {
	m.modules = append(m.modules, module)
}

// GetModules retrieves all current models added to Marvin
func (m *Marvin) GetModules() []models.MarvinModule {
	return m.modules
}

// GetEndpoints retrieves all REST endpoints defined for Marvin including module endpoints
func (m *Marvin) GetEndpoints() []models.MarvinEndpoint {
	eps := make([]models.MarvinEndpoint, 0)
	eps = append(eps, m.restEndpoints...)

	for _, module := range m.GetModules() {
		mEps := module.GetEndpoints()
		if len(mEps) == 0 {
			continue
		}

		for _, ep := range mEps {
			eps = append(eps, ep)
		}
	}

	return eps
}

// Start Marvin, Start setups the modules and registers all module actions
func (m *Marvin) Start() {
	for _, module := range m.GetModules() {
		module.Setup()
		actions := module.GetMarvinActions()
		if len(actions) == 0 {
			continue
		}

		for _, action := range actions {
			events.On(action)
		}
	}

	m.restEndpoints = rest.CreateEndPoints()
}
