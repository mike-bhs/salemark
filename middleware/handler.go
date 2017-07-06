package middleware

import (
	routes "github.com/salemark/routes"
	u "github.com/salemark/utils"
	"net/http"
	"regexp"
)

func ProcessRequest(res http.ResponseWriter, req *http.Request, routesList []routes.Route) {
	u.LogRequest(req)

	for _, route := range routesList {
		if isSearch(req) {
			routes.SearchHandler(res, req)
			return
		}

		if req.Method == route.Method && req.URL.Path == route.Path {
			route.Handler(res, req)
			return
		}
	}

	code := 404
	u.JsonResponse(res, code, http.StatusText(code))
}

func HandleRequests() {
	handlerWrapper := func(w http.ResponseWriter, r *http.Request) {
		ProcessRequest(w, r, routes.List())
	}

	http.HandleFunc("/", handlerWrapper)
}

func isSearch(req *http.Request) bool {
	pattern := "/search.*"

	matched, err := regexp.Match(pattern, []byte(req.URL.Path))
	hasErr := u.HandleError(err)

	if req.Method != "GET" || hasErr {
		return false
	}

	if matched {
		return true
	}

	return false
}
