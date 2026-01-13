package main

import (
	"net/http"
)

// internalServerError logs and sends a 500 response.
func (app *application) internalServerError(w http.ResponseWriter, req *http.Request, err error) {
	app.config.logger.Errorf("internal server error: %s path: %s error: %s", req.Method, req.URL.Path, err)
	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

// badRequestError logs and sends a 400 response.
func (app *application) badRequestError(w http.ResponseWriter, req *http.Request, err error) {
	app.config.logger.Errorf("bad request error: %s path: %s error: %s", req.Method, req.URL.Path, err)
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

// notFoundError logs and sends a 404 response.
func (app *application) notFoundError(w http.ResponseWriter, req *http.Request, err error) {
	app.config.logger.Errorf("not found error: %s path: %s error: %s", req.Method, req.URL.Path, err)
	writeJSONError(w, http.StatusNotFound, err.Error())
}

// conflictResponseError logs and sends a 409 response.
func (app *application) conflictResponseError(w http.ResponseWriter, req *http.Request, err error) {
	app.config.logger.Errorf("conflict server error: %s path: %s error: %s", req.Method, req.URL.Path, err)
	writeJSONError(w, http.StatusConflict, err.Error())
}
