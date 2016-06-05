package huemodule

import (
	"encoding/json"
	"github.com/eyesight-tech/go.hue"
	"github.com/tebben/marvin/go/marvin/models"
)

type ActionHueSetState struct {
	models.Action
	module *HueModule
}

type stateMessage struct {
	States []lightState `json:"states"`
}

type lightState struct {
	Id  string `json:"id"`
	Hue string `json:"hue"`
	Bri string `json:"bri"`
	Sat string `json:"sat"`
}

func CreateHueSetState(module *HueModule) models.MarvinAction {
	a := &ActionHueSetState{}
	a.module = module
	a.ActionName = "hueSetState"
	a.Name = "Set state"
	a.Description = "Set the state for multiple lights"

	states := stateMessage{
		States: []lightState{
			lightState{
				Id:  "1",
				Hue: "50000",
				Bri: "254",
				Sat: "254",
			},
			lightState{
				Id:  "2",
				Hue: "80000",
				Bri: "254",
				Sat: "254",
			},
		},
	}

	a.Sample = models.ActionMessage{Action: a.ActionName, Payload: models.ToRawJson(states)}
	return a
}

func (a *ActionHueSetState) Execute(msg *json.RawMessage) {
	var sm stateMessage
	err := json.Unmarshal(*msg, &sm)
	if err != nil {
		return
	}

	for _, state := range sm.States {
		l, err := a.module.hueBridge.FindLightById(state.Id)
		if err != nil {
			return
		}

		l.SetState(hue.SetLightState{On: "true", Hue: state.Hue, Bri: state.Bri, Sat: state.Sat})
	}
}
