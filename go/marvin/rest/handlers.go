package rest

import (
	"encoding/json"
	marvinErrors "github.com/tebben/marvin/go/errors"
	"github.com/tebben/marvin/go/marvin/models"
	"log"
	"net/http"
)

// HandleGetModules retrieves all modules from Marvin
func HandleGetModules(w http.ResponseWriter, r *http.Request, m *models.Marvin) {
	marvin := *m
	handle := func() interface{} { return marvin.GetModules() }
	HandleGetRequest(w, r, &handle)
}

// handleGetRequest is the default function to handle incoming GET requests
func HandleGetRequest(w http.ResponseWriter, r *http.Request, h *func() interface{}) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
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
		log.Printf("%v", err.Error())
		//panic(err)
	}

	w.Write(b)
}

// sendError creates an ErrorResponse message and sets it to the user
// using SendJSONResponse
func SendError(w http.ResponseWriter, error error) {
	// Set te status code, default 500 for error, check if there is an ApiError an get
	// the status code
	var statusCode = http.StatusInternalServerError
	if error != nil {
		switch e := error.(type) {
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
			Message:    error.Error(),
		},
	}

	sendJSONResponse(w, statusCode, errorResponse)
}
