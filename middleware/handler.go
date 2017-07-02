package middleware

import (
	routes "github.com/salemark/routes"
	u "github.com/salemark/utils"
	"net/http"
)

func ProcessRequest(res http.ResponseWriter, req *http.Request, routesList []routes.Route) {
	u.LogRequest(req)

	for _, route := range routesList {
		if req.Method == route.Method && req.URL.Path == route.Path {
			route.Handler(res, req)
			return
		}
	}

	code := 404
	u.JsonResponse(res, code, http.StatusText(code))
}

func HandleRequests() {
	routesList := routes.List()

	handlerWrapper := func(w http.ResponseWriter, r *http.Request) {
		ProcessRequest(w, r, routesList)
	}

	http.HandleFunc("/", handlerWrapper)
}
