package utils

import (
	"encoding/json"
	s "github.com/salemark/services"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func LogRequest(req s.Request) {
	simpleParams, err := req.SimpleParams()

	if err != nil {
		return
	}

	jsonParams, err := json.Marshal(simpleParams)

	if HandleError(err) {
		return
	}

	log.Println(req.Method() + " " + req.Url())
	log.Println("    Parameters: {" + string(jsonParams[:]) + "}")
}

func LogResponse(code int) {
	log.Println("    Response: " + strconv.Itoa(code) + " " + http.StatusText(code))
}
