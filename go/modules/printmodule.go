package modules

import (
	"github.com/tebben/marvin/go/models"
	"log"
)

type PrintModule struct {
	name        string
	description string
	events      []models.MarvinEvent
	actions     []models.MarvinAction
}

func (pm *PrintModule) Setup() {
	pm.name = "Printer"
	pm.description = "print a message"
	pa := &PrintAction{actionName: "print", name: "Default Print", description: "Print a message to the log"}
	pm.actions = []models.MarvinAction{pa}
}

func (pm *PrintModule) GetName() string {
	return pm.name
}

func (pm *PrintModule) GetDescription() string {
	return pm.description
}

func (pm *PrintModule) GetMarvinEvents() []models.MarvinEvent {
	return pm.events
}

func (pm *PrintModule) GetMarvinActions() []models.MarvinAction {
	return pm.actions
}

type PrintAction struct {
	actionName  string
	name        string
	description string
}

func (pa *PrintAction) GetActionName() string {
	return pa.actionName
}

func (pa *PrintAction) GetName() string {
	return pa.name
}

func (pa *PrintAction) GetDescription() string {
	return pa.description
}

func (p *PrintAction) Execute(msg map[string]interface{}) {
	log.Printf("%v", msg["msg"])
}
