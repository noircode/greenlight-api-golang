package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Initialize a new httprouter router instance
	router := httprouter.New()

	// Register the relevant methods, URL patterns and handler function for our
	// endpoint
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	// Return the htpprouter instance
	return router

}
