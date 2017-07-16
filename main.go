package main

import (
	"log"
	"net/http"

	middleware "github.com/salemark/middleware"
)

func main() {
	log.Println("Starting elasticsearch ...")

	middleware.HandleRequests()

	log.Println("Server is running at http://localhost:8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
