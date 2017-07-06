package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func LogRequest(req *http.Request) {
	err := req.ParseForm()

	if HandleError(err) {
		return
	}

	jsonParams, err := json.Marshal(RequestParams(req))

	if HandleError(err) {
		return
	}

	log.Println(req.Method + " " + req.URL.String())
	log.Println("    Parameters: {" + string(jsonParams[:]) + "}")
}

func LogResponse(code int) {
	log.Println("    Response: " + strconv.Itoa(code) + " " + http.StatusText(code))
}

func GetParamValue(req *http.Request, param string) string {
	err := req.ParseForm()

	if HandleError(err) {
		return ""
	}

	for key, values := range req.Form {
		if key == param {
			v := strings.Join(values[:], ", ")
			return v
		}
	}

	return ""
}
