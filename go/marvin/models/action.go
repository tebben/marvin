package models

type MarvinAction interface {
	GetActionName() string
	GetName() string
	GetDescription() string
	Execute(msg map[string]interface{})
}

type Action struct {
	ActionName string `json:"actionName"`
	Name string `json:"name"`
	Description string `json:"description"`
}

func (a *Action) GetActionName() string {
	return a.ActionName
}

func (a *Action) GetName() string {
	return a.Name
}

func (a *Action) GetDescription() string {
	return a.Description
}
