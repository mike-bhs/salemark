package utils

import (
	"fmt"
	"net/http"
)

func Response(res http.ResponseWriter, statusCode int, contentType, body string) {
	res.WriteHeader(statusCode)
	res.Header().Set("Content-Type", contentType)
	LogResponse(statusCode)
	fmt.Fprintln(res, body)
}

func HtmlResponse(res http.ResponseWriter, statusCode int, body string) {
	contentType := "text/html; charset=utf-8"
	Response(res, statusCode, contentType, body)
}

func JsonResponse(res http.ResponseWriter, statusCode int, body string) {
	contentType := "application/json; charset=utf-8"
	Response(res, statusCode, contentType, body)
}
