package models

import (
	"encoding/json"
	"net/http"
)

type Marvin interface {
	AddModule(module MarvinModule)
	GetModules() []MarvinModule
	GetEndpoints() []MarvinEndpoint
	Start()
}

type ActionMessage struct {
	Action  string           `json:"action"`
	Payload *json.RawMessage `json:"payload,omitempty"`
}

func ToRawJson(i interface{}) *json.RawMessage {
	p, _ := json.Marshal(i)
	raw := json.RawMessage(p)
	return &raw
}

// Server interface for starting and stopping the HTTP server
type HTTPServer interface {
	Start()
	Stop()
}

// HTTPOperation describes the HTTP operation such as GET POST DELETE.
type HTTPOperation string

// HTTPOperation is a "enumeration" of the HTTP operations needed for all endpoints.
const (
	HTTPOperationGet    HTTPOperation = "GET"
	HTTPOperationPost   HTTPOperation = "POST"
	HTTPOperationPatch  HTTPOperation = "PATCH"
	HTTPOperationDelete HTTPOperation = "DELETE"
)

// HTTPHandler func defines the format of the handler to process the incoming request
type HTTPHandler func(w http.ResponseWriter, r *http.Request, m *Marvin)

// EndpointOperation contains the needed information to create an endpoint in the HTTP.Router
type EndpointOperation struct {
	OperationType HTTPOperation `json:"operation"`
	Path          string        `json:"path"` //relative path to the endpoint for example: /v1.0/myendpoint/
	Handler       HTTPHandler   `json:"-"`
}

// Endpoint defines the rest endpoint options
type MarvinEndpoint interface {
	GetName() string
	GetOperations() []EndpointOperation
}

type Endpoint struct {
	Name       string
	Operations []EndpointOperation
}

func (e *Endpoint) GetName() string {
	return e.Name
}

func (e *Endpoint) GetOperations() []EndpointOperation {
	return e.Operations
}

// ErrorResponse is the default response format for sending errors back
type ErrorResponse struct {
	Error ErrorContent `json:"error"`
}

// ErrorContent holds information on the error that occurred
type ErrorContent struct {
	StatusText string `json:"status"`
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}
