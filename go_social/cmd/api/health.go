package main

import (
	"log"
	"net/http"
)

// healthCheckHandler reports service availability and metadata.
//
//	@Summary		Health check
//	@Description	Returns service status, environment, and version.
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/v1/health" {
		log.Println("404 not Found")
		writeJSONError(w, http.StatusNotFound, "404 not Found")
		return
	}

	if req.Method != "GET" {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// map of health message
	data := map[string]string{
		"status":  "up",
		"env":     app.config.env,
		"version": version,
	}

	if err := writeJSON(w, http.StatusOK, data); err != nil {
		log.Println(err)
		writeJSONError(w, http.StatusInternalServerError, err.Error())
	}
}
