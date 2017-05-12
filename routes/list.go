package routes

import (
	"net/http"
)

type Route struct {
	Method  string
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
}

func List() []Route {
	routesList := []Route{
		Route{Method: "GET", Path: "/", Handler: rootHandler},
	}

	return routesList
}
