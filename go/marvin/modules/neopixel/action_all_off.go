package neopixel

import (
	"encoding/json"
	"github.com/tebben/marvin/go/marvin/models"
)

const (
	aOffName        = "Neopixels off"
	aOffActionName  = "NeopixelsOff"
	aOffDescription = "Turn off all your Neopixels"
)

type NeopixelOffAction struct {
	models.Action
	mod *NeopixelModule
}

func CreateNeopixelOffAction(mod *NeopixelModule) models.MarvinAction {
	a := &NeopixelOffAction{}
	a.ActionName = aOffActionName
	a.Name = aOffName
	a.Description = aOffDescription
	a.Sample = models.ActionMessage{Action: a.ActionName}
	a.mod = mod

	return a
}

func (p *NeopixelOffAction) Execute(msg *json.RawMessage) {
	p.mod.Port.Write([]byte("allOff"))
}
