package modules

import (
	"github.com/tebben/marvin/go/models"
	"fmt"
	"github.com/eyesight-tech/go.hue"
	"log"
)

type HueModule struct {
	name        string
	description string
	events      []models.MarvinEvent
	actions     []models.MarvinAction

	hueBridge *hue.Bridge
 hueLights []*hue.Light
}

func (mm *HueModule) Setup() {
	mm.name = "Hue"
	mm.description = "Control you Philips Hue lights"

	mm.actions = []models.MarvinAction{
		&ActionHueAllOn{module: mm , actionName: "allOn", name: "All lights on", description: "Turn on al your Hue lights"},
		&ActionHueAllOff{module: mm , actionName: "allOff", name: "All lights off", description: "Turn off all your Hue lights"},
	}

	mm.hueBridge = hue.NewBridge("192.168.178.19", "6cfd742666978686a932f7824339135")
	lights, err := mm.hueBridge.GetAllLights()

	if(err != nil){
		fmt.Printf("%v", err)
	}

	mm.hueLights = lights
}

func (mm *HueModule) GetName() string {
	return mm.name
}

func (mm *HueModule) GetDescription() string {
	return mm.description
}

func (mm *HueModule) GetMarvinEvents() []models.MarvinEvent {
	return mm.events
}

func (mm *HueModule) GetMarvinActions() []models.MarvinAction {
	return mm.actions
}

type ActionHueAllOn struct {
	module *HueModule
	actionName  string
	name        string
	description string
}

func (ma *ActionHueAllOn) GetActionName() string {
	return ma.actionName
}

func (ma *ActionHueAllOn) GetName() string {
	return ma.name
}

func (ma *ActionHueAllOn) GetDescription() string {
	return ma.description
}

func (ma *ActionHueAllOn) Execute(msg map[string]interface{}) {
	log.Println("TURNING HUE LIGHTS ON")
	for _, light := range ma.module.hueLights {
		light.ColorLoop()
	}
}

type ActionHueAllOff struct {
	module *HueModule
	actionName  string
	name        string
	description string
}

func (ma *ActionHueAllOff) GetActionName() string {
	return ma.actionName
}

func (ma *ActionHueAllOff) GetName() string {
	return ma.name
}

func (ma *ActionHueAllOff) GetDescription() string {
	return ma.description
}

func (ma *ActionHueAllOff) Execute(msg map[string]interface{}) {
	log.Println("TURNING HUE LIGHTS OFF")
	for _, light := range ma.module.hueLights {
		light.Off()
	}
}

func test (){
	/*locators, _ := hue.DiscoverBridges(false)
locator := locators[0] // find the first locator
deviceType := "my nifty app"

// remember to push the button on your hue first
bridge, _ := locator.CreateUser(deviceType)
fmt.Printf("registered new device => %+v\n", bridge)*/

	/*
		bridge := hue.NewBridge("192.168.178.19", "6cfd742666978686a932f7824339135")
		lights, err := bridge.GetAllLights()

		if(err != nil){
			fmt.Printf("%v", err)
		}

		for _, light := range lights {
			light.ColorLoop()
		}
	*/
}
// {"action": "print", "payload":{"msg": "BAMS!"}}
