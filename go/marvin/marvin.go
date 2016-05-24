package marvin

import (
	"github.com/tebben/marvin/go/events"
	"github.com/tebben/marvin/go/models"
)

type Marvin struct {
	modules []models.MarvinModule
}

// AddModule add a new module to Marvin
func (m *Marvin) AddModule(module models.MarvinModule) {
	m.modules = append(m.modules, module)
}

// GetModules retrieves all current models added to Marvin
func (m *Marvin) GetModules() []models.MarvinModule {
	return m.modules
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
}

func (m *Marvin) Trigger(actionName string, msg map[string]interface{}) {
	events.Fire(actionName, msg)
}
