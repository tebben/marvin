package events

import (
	"encoding/json"
	"github.com/tebben/marvin/go/marvin/models"
)

func On(action models.MarvinAction) error {
	return add(action)
}

func Fire(actionName string, msg *json.RawMessage) {
	invoke(actionName, msg)
}

func Clear(event string) error {
	return clear(event)
}

func ClearEvents() error {
	return deleteAll()
}

func HasEvent(event string) bool {
	return hasEvent(event)
}

func Events() []string {
	return eventList()
}

func EventCount() int {
	return eventCount()
}
