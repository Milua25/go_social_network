package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

// init configures the request validator with required struct support.
func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

// writeJSON marshals data as JSON with the given status code.
func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

// readJSON decodes a JSON body into data with size and unknown-field checks.
func readJSON(w http.ResponseWriter, req *http.Request, data any) error {
	maxBytes := 1_048_578 // 1mb

	req.Body = http.MaxBytesReader(w, req.Body, int64(maxBytes))

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func writeJSONError(w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Error string `json:"error"`
	}

	return writeJSON(w, status, &envelope{Error: message})
}

// jsonResponse wraps data in a consistent envelope and writes JSON.
func (app *application) jsonResponse(w http.ResponseWriter, status int, data any) error {
	type envelope struct {
		Data any `json:"data"`
	}

	if status == http.StatusNoContent || status == http.StatusResetContent {
		w.WriteHeader(status)
		return nil
	}

	return writeJSON(w, status, &envelope{Data: data})
}
