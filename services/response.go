package services

import (
	"fmt"
	"net/http"

	utils "github.com/salemark/utils"
)

type Response struct {
	Source http.ResponseWriter
}

func (res Response) WriteResponse(statusCode int, contentType, body string) {
	res.Source.WriteHeader(statusCode)
	res.Source.Header().Set("Content-Type", contentType)

	utils.LogResponse(statusCode)
	fmt.Fprintln(res.Source, body)
}

func (res Response) HtmlResponse(statusCode int, body string) {
	contentType := "text/html; charset=utf-8"
	res.WriteResponse(statusCode, contentType, body)
}

func (res Response) JsonResponse(statusCode int, body string) {
	contentType := "application/json; charset=utf-8"
	res.WriteResponse(statusCode, contentType, body)
}
