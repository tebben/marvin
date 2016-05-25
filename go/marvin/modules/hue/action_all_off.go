package huemodule

import (
	"log"
	"github.com/tebben/marvin/go/marvin/models"
)

type ActionHueAllOff struct {
	models.Action
	module *HueModule
}

func CreateHueAllOff(module *HueModule) models.MarvinAction{
	a := &ActionHueAllOff{}
	a.ActionName =  "allOff"
	a.Name = "All lights off"
	a.Description = "Turn off all your Hue lights"

	return a
}

func (ma *ActionHueAllOff) Execute(msg map[string]interface{}) {
	log.Println("TURNING HUE LIGHTS OFF")
	for _, light := range ma.module.hueLights {
		light.Off()
	}
}
