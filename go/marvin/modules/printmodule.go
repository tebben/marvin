package modules

import (
	"github.com/tebben/marvin/go/marvin/models"
	"log"
)

type PrintModule struct {
	models.Module
}

func (pm *PrintModule) Setup() {
	pm.Name = "Printer"
	pm.Description = "print a message"
	pm.Actions = []models.MarvinAction{CreatePrintAction()}
}

type PrintAction struct {
	models.Action
}

func CreatePrintAction() models.MarvinAction{
	a := &PrintAction{}
	a.ActionName =  "print"
	a.Name = "Default Print"
	a.Description = "Print a message to the log"

	return a
}

func (p *PrintAction) Execute(msg map[string]interface{}) {
	log.Printf("%v", msg["msg"])
}
