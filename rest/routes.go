package rest

import (
	"net/http"
)

var routes = ConfiguredRoutes{
	Route{
		Name:        "GetNearEvents",
		Method:      "GET",
		Pattern:     "/near",
		HandlerFunc: HandleGetNear},
}

type ConfiguredRoutes []Route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc AuthHandler
}

type AuthHandler func(http.ResponseWriter, *http.Request)

func (fn AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// TODO: implement authentication

	fn(w, r)
}