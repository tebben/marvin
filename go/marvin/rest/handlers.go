package rest

import (
"net/http"
	"github.com/tebben/marvin/go/marvin/models"
	"encoding/json"
	marvinErrors "github.com/tebben/marvin/go/errors"
)

// HandleGetModules retrieves all modules from Marvin
func HandleGetModules(w http.ResponseWriter, r *http.Request, m *models.Marvin) {
	marvin := * m
	handle := func() (interface{}) { return marvin.GetModules() }
	handleGetRequest(w, r, &handle)
}

// handleGetRequest is the default function to handle incoming GET requests
func handleGetRequest(w http.ResponseWriter, r *http.Request, h *func() (interface{})) {
	// Run the handler func such as Api.GetThingById
	handler := *h
	data := handler()
	sendJSONResponse(w, http.StatusOK, data)
}

// sendJSONResponse sends the desired message to the user
// the message will be marshalled into an indented JSON format
func sendJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	b, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		panic(err)
	}

	w.Write(b)
}

// sendError creates an ErrorResponse message and sets it to the user
// using SendJSONResponse
func sendError(w http.ResponseWriter, error []error) {
	//errors cannot be marshalled, create strings
	errors := make([]string, len(error))
	for idx, value := range error {
		errors[idx] = value.Error()
	}

	// Set te status code, default 500 for error, check if there is an ApiError an get
	// the status code
	var statusCode = http.StatusInternalServerError
	if error != nil && len(error) > 0 {
		switch e := error[0].(type) {
		case marvinErrors.APIError:
			statusCode = e.GetHTTPErrorStatusCode()
			break
		}
	}

	statusText := http.StatusText(statusCode)
	errorResponse := models.ErrorResponse{
		Error: models.ErrorContent{
			StatusText: statusText,
			StatusCode: statusCode,
			Messages:   errors,
		},
	}

	sendJSONResponse(w, statusCode, errorResponse)
}

