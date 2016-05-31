package models

type MarvinEvent interface {
	GetName() string
	GetDescription() string
}

type Event struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (e *Event) GetName() string {
	return e.Name
}

func (e *Event) GetDescription() string {
	return e.Description
}
