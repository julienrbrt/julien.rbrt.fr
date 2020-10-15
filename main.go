package main

import (
	"log"
	"net/http"
	"os"

	"julien.rbrt.fr/router"
)

func main() {
	// app port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Listening on port", port)

	err := http.ListenAndServe(":"+port, router.NewRouter())
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
