package neopixel

import (
	"encoding/json"
	"github.com/tebben/marvin/go/marvin/models"
)

const (
	aOnName        = "Neopixels on"
	aOnActionName  = "NeopixelsOn"
	aOnDescription = "Turn on all your Neopixels"
)

type NeopixelOnAction struct {
	models.Action
	mod *NeopixelModule
}

func CreateNeopixelOnAction(mod *NeopixelModule) models.MarvinAction {
	a := &NeopixelOnAction{}
	a.ActionName = aOnActionName
	a.Name = aOnName
	a.Description = aOnDescription
	a.Sample = models.ActionMessage{Action: a.ActionName}
	a.mod = mod

	return a
}

func (p *NeopixelOnAction) Execute(msg *json.RawMessage) {
	p.mod.Port.Write([]byte("allOn"))
}
