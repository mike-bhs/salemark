package main

import (
	middleware "github.com/salemark/middleware"
	"log"
	"net/http"
)

func main() {
	middleware.HandleRequests()

	log.Println("Server is running at http://localhost:8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
