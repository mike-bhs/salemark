package utils

import (
	"encoding/json"
	"log"

	s "github.com/salemark/services"
)

func LogRequest(req s.Request) {
	simpleParams, err := req.SimpleParams()

	if err != nil {
		return
	}

	jsonParams, err := json.Marshal(simpleParams)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(req.Method() + " " + req.Url())
	log.Println("    Parameters: {" + string(jsonParams[:]) + "}")
}
