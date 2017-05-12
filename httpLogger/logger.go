package middleware

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func LogRequest(req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		log.Panicln(err)
		return
	}

	var reqParams []string

	for key, values := range req.Form {
		v := strings.Join(values[:], ", ")
		reqParams = append(reqParams, "\""+key+"\": \""+v+"\"")
	}

	log.Println(req.Method + " " + req.URL.String())
	log.Println("    Parameters: {" + strings.Join(reqParams[:], ", ") + "}")
}

func LogResponse(code int) {
	log.Println("    Response: " + strconv.Itoa(code) + " " + http.StatusText(code))
}