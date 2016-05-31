package huemodule

import (
	"github.com/tebben/marvin/go/marvin/models"
	"github.com/eyesight-tech/go.hue"
)

type ActionHueSetState struct {
	models.Action
	module *HueModule
}

func CreateHueSetState(module *HueModule) models.MarvinAction{
	a := &ActionHueSetState{}
	a.module = module
	a.ActionName =  "hueSetState"
	a.Name = "Set state"
	a.Description = "Set the state for multiple lights"

	states := make([]map[string]interface{}, 0)
	state1 := make(map[string]interface{}, 0)
	state1["id"] = "1"
	state1["hue"] = "50000"
	state1["bri"] = "254"
	state1["sat"] = "254"
	states = append(states, state1)
	payload := make(map[string]interface{})
	payload["states"] = states
	a.Sample = models.ActionMessage{ Action: a.ActionName, Payload: payload }
	return a

	/*
	Hue       int       `json:"hue"`
	On        bool      `json:"on"`
	Effect    string    `json:"effect"`
	Alert     string    `json:"effect"`
	Bri       int       `json:"bri"`
	Sat       int       `json:"sat"`
	Ct        int       `json:"ct"`
	Xy        []float32 `json:"xy"`
	Reachable bool      `json:"reachable"`
	ColorMode string    `json:"colormode"`
	 */
}


func (aao *ActionHueSetState) Execute(msg map[string]interface{}) {
	states := msg["states"].([]interface{})
	for _, state := range states {
		s := state.(map[string]interface{})
		l, err := aao.module.hueBridge.FindLightById(s["id"].(string))
		if(err != nil){
			return
		}

		l.SetState(hue.SetLightState{ On: "true", Hue: s["hue"].(string), Bri: s["bri"].(string), Sat: s["sat"].(string)})
	}
}
