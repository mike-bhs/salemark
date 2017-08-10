package services

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Source http.ResponseWriter
}

func (res Response) writeResponse(statusCode int, contentType, body string) {
	res.Source.WriteHeader(statusCode)
	res.Source.Header().Set("Content-Type", contentType)

	log.Println("    Response: " + strconv.Itoa(statusCode) + " " + http.StatusText(statusCode))
	fmt.Fprintln(res.Source, body)
}

func (res Response) HtmlResponse(statusCode int, body string) {
	contentType := "text/html; charset=utf-8"
	res.writeResponse(statusCode, contentType, body)
}

func (res Response) JsonResponse(statusCode int, body string) {
	contentType := "application/json; charset=utf-8"
	res.writeResponse(statusCode, contentType, body)
}
