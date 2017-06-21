package routes

import (
	"net/http"
)

type Route struct {
	Method  string
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
}

// base handler wich takes req and and root handler
// it will also parse dictionary params

func List() []Route {
	routesList := []Route{
		Route{Method: "GET", Path: "/", Handler: rootHandler},
		Route{Method: "POST", Path: "/search", Handler: rootHandler},
	}

	return routesList
}
