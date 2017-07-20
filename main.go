package main

import (
	"log"
	"net/http"

	middleware "github.com/salemark/middleware"
	s "github.com/salemark/services"
)

func main() {
	log.Println("Starting elasticsearch ...")

	middleware.HandleRequests()

	log.Println("Server is running at http://localhost:8080")
	s.StartParsing()
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
