package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Listening at http://localhost:8000")
	panic(http.ListenAndServe(":8000", http.FileServer(http.Dir("ui/dist"))))
}
