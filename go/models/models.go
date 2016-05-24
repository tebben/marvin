package models

import ()

type ActionMessage struct {
	Action  string                 `json:"action"`
	Payload map[string]interface{} `json:"payload"`
}

type MarvinEvent interface {
	GetName() string
	GetDescription() string
}

type MarvinAction interface {
	GetActionName() string
	GetName() string
	GetDescription() string
	Execute(msg map[string]interface{})
}

type MarvinModule interface {
	GetName() string
	GetDescription() string
	GetMarvinEvents() []MarvinEvent
	GetMarvinActions() []MarvinAction

	Setup()
}
