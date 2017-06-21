package utils

import (
	"log"
)

func HandleError(err error) bool {
	if err != nil {
		log.Panicln(err)
		return true
	}

	return false
}
