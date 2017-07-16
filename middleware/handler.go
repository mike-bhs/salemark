package middleware

import (
	"net/http"

	routes "github.com/salemark/routes"
	s "github.com/salemark/services"
	u "github.com/salemark/utils"
)

const (
	notFound = 404
)

func ProcessRequest(res s.Response, req s.Request, routesList []routes.Route) {
	u.LogRequest(req)

	for _, route := range routesList {
		if req.MatchPatter("/search.*") && req.Method() == "GET" {
			routes.SearchHandler(res.Source, req.Source)
			return
		}

		if req.MatchRoute(route) {
			route.Handler(res.Source, req.Source)
			return
		}
	}

	res.JsonResponse(notFound, http.StatusText(notFound))
}

func HandleRequests() {
	handlerWrapper := func(w http.ResponseWriter, r *http.Request) {
		req := s.Request{Source: r}
		res := s.ResponseWriter{Source: w}

		ProcessRequest(res, req, routes.List())
	}

	http.HandleFunc("/", handlerWrapper)
}
