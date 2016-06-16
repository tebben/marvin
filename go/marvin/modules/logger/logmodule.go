package logmodule

import (
	"encoding/json"
	"log"

	"github.com/tebben/marvin/go/marvin/models"
)

type LogModule struct {
	models.Module
}

type printMessage struct {
	Msg string `json:"msg"`
}

func (lm *LogModule) Setup() {
	lm.Name = "Logger"
	lm.Description = "Log a message"
	lm.Actions = []models.MarvinAction{CreateLogAction()}
}

type PrintAction struct {
	models.Action
}

func CreateLogAction() models.MarvinAction {
	a := &PrintAction{}
	a.ActionName = "log"
	a.Name = "Default action to log a message"
	a.Description = "Send a message to the log"
	a.Sample = models.ActionMessage{Action: a.ActionName, Payload: models.ToRawJson(printMessage{Msg: "Hello logger!"})}

	return a
}

func (p *PrintAction) Execute(msg *json.RawMessage) {
	var pm printMessage
	json.Unmarshal(*msg, &pm)
	log.Printf("%v", pm.Msg)
}
