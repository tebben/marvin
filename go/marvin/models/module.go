package models

type MarvinModule interface {
	GetName() string
	GetDescription() string
	GetMarvinEvents() []MarvinEvent
	GetMarvinActions() []MarvinAction

	Setup()
}

type Module struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Events      []MarvinEvent  `json:"events"`
	Actions     []MarvinAction `json:"actions"`
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
