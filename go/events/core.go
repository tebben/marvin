package events

import (
	"encoding/json"
	"errors"
	"github.com/tebben/marvin/go/marvin/models"
	"log"
	"runtime"
)

type EventType interface {
	Parse(msg string)
}

var functionMap map[string]models.MarvinAction

func init() {
	functionMap = make(map[string]models.MarvinAction)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func add(action models.MarvinAction) error {
	if _, ok := functionMap[action.GetActionName()]; ok {
		return errors.New("Action Already Defined")
	}
	functionMap[action.GetActionName()] = action
	return nil
}

func invoke(actionName string, msg *json.RawMessage) {
	if _, ok := functionMap[actionName]; ok {
		functionMap[actionName].Execute(msg)
		return
	}

	log.Printf("Action %s called but not implemented or hooked up", actionName)
}

func clear(event string) error {
	if _, ok := functionMap[event]; !ok {
		return errors.New("Event Not Defined")
	}
	delete(functionMap, event)
	return nil
}

func deleteAll() error {
	functionMap = make(map[string]models.MarvinAction)
	return nil
}

func eventList() []string {
	events := make([]string, 0)
	for k := range functionMap {
		events = append(events, k)
	}
	return events
}

func eventCount() int {
	return len(functionMap)
}

func hasEvent(event string) bool {
	_, ok := functionMap[event]
	return ok
}
