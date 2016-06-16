package models

type MarvinModule interface {
	GetName() string
	GetDescription() string
	GetMarvinEvents() []MarvinEvent
	GetMarvinActions() []MarvinAction
	GetEndpoints() []MarvinEndpoint
	GetSettings() interface{}

	Setup()
	SettingsChanged(newSettings interface{})
}

type Module struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Events      []MarvinEvent    `json:"events"`
	Actions     []MarvinAction   `json:"actions"`
	Endpoints   []MarvinEndpoint `json:"endpoints,omitempty"`
}

func (mm *Module) GetName() string {
	return mm.Name
}

func (mm *Module) GetDescription() string {
	return mm.Description
}

func (mm *Module) GetMarvinEvents() []MarvinEvent {
	return mm.Events
}

func (mm *Module) GetMarvinActions() []MarvinAction {
	return mm.Actions
}

func (mm *Module) GetEndpoints() []MarvinEndpoint {
	return mm.Endpoints
}

func (mm *Module) GetSettings() interface{} {
	return nil
}

func (mm *Module) SettingsChanged(newSettings interface{}) {

}
