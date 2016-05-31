package huemodule

import (
	"fmt"
	"github.com/eyesight-tech/go.hue"
	"github.com/tebben/marvin/go/marvin/models"
)

type HueModule struct {
	models.Module
	hueBridge *hue.Bridge
	hueLights []*hue.Light
}

func (mm *HueModule) Setup() {
	mm.Name = "Hue"
	mm.Description = "Control your Philips Hue lights"
	mm.Actions = []models.MarvinAction{
		CreateHueAllOff(mm),
		CreateHueAllOn(mm),
		CreateHueSetState(mm),
	}

	mm.hueBridge = hue.NewBridge("192.168.178.19", "6cfd742666978686a932f7824339135")

	lights, err := mm.hueBridge.GetAllLights()
	if err != nil {
		fmt.Printf("%v", err)
	}

	mm.hueLights = lights
}

func test() {
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
