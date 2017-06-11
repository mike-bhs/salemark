package routes

import (
	"fmt"
	logger "github.com/salemark/httpLogger"
	"net/http"
)

func rootHandler(res http.ResponseWriter, req *http.Request) {
	// reusable response writer
	res.WriteHeader(200)
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	logger.LogResponse(200)
	fmt.Fprintln(res, "<h1>Welcome!</h1>")
}
