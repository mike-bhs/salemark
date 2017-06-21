package routes

import (
	"fmt"
	logger "github.com/salemark/httpLogger"
	search "github.com/salemark/search"
	"net/http"
)

func SearchHandler(res http.ResponseWriter, req *http.Request, s search.Search) {
	// reusable response writer
	search.Find(s)
	res.WriteHeader(200)
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	logger.LogResponse(200)
	fmt.Fprintln(res, "<h1>Search!</h1>")
}
