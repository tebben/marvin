package modules

import (
	"encoding/json"
	"log"

	"github.com/tebben/marvin/go/marvin/models"
)

type PrintModule struct {
	models.Module
}

type printMessage struct {
	Msg string `json:"msg"`
}

func (pm *PrintModule) Setup() {
	pm.Name = "Printer"
	pm.Description = "print a message"
	pm.Actions = []models.MarvinAction{CreatePrintAction()}
}

type PrintAction struct {
	models.Action
}

func CreatePrintAction() models.MarvinAction {
	a := &PrintAction{}
	a.ActionName = "print"
	a.Name = "Default Print"
	a.Description = "Print a message to the log"

	payload := printMessage{Msg: "Hello Printer!"}
	a.Sample = models.ActionMessage{Action: a.ActionName, Payload: models.ToRawJson(payload)}

	return a
}

func (p *PrintAction) Execute(msg *json.RawMessage) {
	var pm printMessage
	err := json.Unmarshal(*msg, &pm)
	if err != nil {
		log.Printf("%v", err.Error())
	}

	log.Printf("PRINT: %v", pm.Msg)
}
