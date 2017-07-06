package main

import (
	middleware "github.com/salemark/middleware"
	// search "github.com/salemark/search"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting elasticsearch ...")
	// search.Start()

	middleware.HandleRequests()

	log.Println("Server is running at http://localhost:8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
