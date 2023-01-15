package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	log.Println("Server will start at http://localhost:8000/")

	e := AddApproutes()

	log.Fatal(http.ListenAndServe(":8000", e))
}
