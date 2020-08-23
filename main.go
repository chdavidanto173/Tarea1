package main

import (
	"log"
	"net/http"
	"os"

	sw "github.com/arce/gowebservices/go"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
