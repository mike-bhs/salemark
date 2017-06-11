package middleware

import (
	"fmt"
	logger "github.com/salemark/httpLogger"
	routes "github.com/salemark/routes"
	"net/http"
)

func ProcessRequest(res http.ResponseWriter, req *http.Request, routesList []routes.Route) {
	logger.LogRequest(req)

	for _, route := range routesList {
		if req.Method == route.Method && req.URL.Path == route.Path {
			route.Handler(res, req)
			return
		}
	}

	DefaultResponse(res, 404)
}

func DefaultResponse(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	logger.LogResponse(code)
	fmt.Fprintln(w, http.StatusText(code))
}

func HandleRequests() {
	routesList := routes.List()

	handlerWrapper := func(w http.ResponseWriter, r *http.Request) {
		ProcessRequest(w, r, routesList)
	}

	http.HandleFunc("/", handlerWrapper)
}
