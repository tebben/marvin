package huemodule

import (
	"fmt"
	"github.com/eyesight-tech/go.hue"
	"github.com/tebben/marvin/go/marvin/models"
	"github.com/tebben/marvin/go/marvin/rest"
	"net/http"
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
	mm.Endpoints = mm.createEndpoints()
	mm.hueBridge = hue.NewBridge("192.168.178.19", "6cfd742666978686a932f7824339135")

	go mm.startup()
}

func (mm *HueModule) startup() {
	lights, err := mm.hueBridge.GetAllLights()
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	mm.hueLights = lights
}

func (mm *HueModule) createEndpoints() []models.MarvinEndpoint {
	eps := []models.MarvinEndpoint{
		&models.Endpoint{
			Name: "Get Lights",
			Operations: []models.EndpointOperation{
				{
					OperationType: models.HTTPOperationGet,
					Path:          "/Hue/Lights",
					Handler:       mm.HandletGetHueLights,
				},
			},
		},
	}

	return eps
}

func (h *HueModule) HandletGetHueLights(w http.ResponseWriter, r *http.Request, m *models.Marvin) {
	l, err := h.hueBridge.GetAllLights()
	if err != nil {
		rest.SendError(w, err)
	}

	handle := func() interface{} { return l }
	rest.HandleGetRequest(w, r, &handle)
}

//func test() {
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
//}

// {"action": "print", "payload":{"msg": "BAMS!"}}
