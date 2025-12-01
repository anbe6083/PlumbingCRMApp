package main

import (
	"log"
	"net/http"
)

func main() {
	server := &CustomerServer{}

	log.Fatal(http.ListenAndServe(":5000", server))
}
