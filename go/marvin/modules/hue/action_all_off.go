package huemodule

import (
	"github.com/tebben/marvin/go/marvin/models"
)

type ActionHueAllOff struct {
	models.Action
	module *HueModule
}

func CreateHueAllOff(module *HueModule) models.MarvinAction{
	a := &ActionHueAllOff{}
	a.module = module
	a.ActionName =  "hueAllOff"
	a.Name = "All lights off"
	a.Description = "Turn off all your Hue lights"
	a.Sample = models.ActionMessage{ Action: a.ActionName }
	return a
}

func (aao *ActionHueAllOff) Execute(msg map[string]interface{}) {
	for _, light := range aao.module.hueLights {
		light.Off()
	}
}
