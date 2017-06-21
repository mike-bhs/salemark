package utils

import (
	"log"
)

func HandleError(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}

	return false
}

func PanicError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
