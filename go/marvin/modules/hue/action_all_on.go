package huemodule

import (
	"encoding/json"
	"github.com/tebben/marvin/go/marvin/models"
)

type ActionHueAllOn struct {
	models.Action
	module *HueModule
}

func CreateHueAllOn(module *HueModule) models.MarvinAction {
	a := &ActionHueAllOn{}
	a.module = module
	a.ActionName = "hueAllOn"
	a.Name = "All lights on"
	a.Description = "Turn on all your Hue lights"
	a.Sample = models.ActionMessage{Action: a.ActionName}
	return a
}

func (a *ActionHueAllOn) Execute(msg *json.RawMessage) {
	for _, light := range a.module.hueLights {
		light.On()
	}
}
