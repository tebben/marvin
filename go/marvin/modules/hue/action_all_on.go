package huemodule

import (
	"log"
	"github.com/tebben/marvin/go/marvin/models"
)

type ActionHueAllOn struct {
	models.Action
	module *HueModule
}

func CreateHueAllOn(module *HueModule) models.MarvinAction{
	a := &ActionHueAllOn{}
	a.ActionName =  "allOn"
	a.Name = "All lights on"
	a.Description = "Turn on all your Hue lights"

	return a
}

func (ma *ActionHueAllOn) Execute(msg map[string]interface{}) {
	log.Println("TURNING HUE LIGHTS ON")
	for _, light := range ma.module.hueLights {
		light.ColorLoop()
	}
}

