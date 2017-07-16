package routes

import (
	s "github.com/salemark/services"
)

// base handler wich takes req and and root handler
// it will also parse dictionary params

func List() []s.Route {
	routesList := []s.Route{
		s.Route{Method: "GET", Path: "/", Handler: rootHandler},
		// Route{Method: "POST", Path: "/search", Handler: SearchHandler},
	}

	return routesList
}
