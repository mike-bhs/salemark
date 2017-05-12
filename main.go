package main

import (
	"log"
	"net/http"
	middleware "salemark/middleware"
)

func main() {
	middleware.HandleRequests()

	log.Println("Server is running at http://localhost:8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
