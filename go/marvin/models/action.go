package models

type MarvinAction interface {
	GetActionName() string
	GetName() string
	GetDescription() string
	Execute(msg map[string]interface{})
}

type Action struct {
	Name        string `json:"name"`
	ActionName  string `json:"actionName"`
	Description string `json:"description"`
	Sample	    ActionMessage `json:"sample,omitempty"`
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
