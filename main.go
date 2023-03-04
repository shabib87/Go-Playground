package main

import (
	"Go-Playground/handlers"
	"log"
	"net/http"
	"os"

	"github.com/shabib87/Go-Playground/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	log.Println("Starting Server")
	error := http.ListenAndServe(":9090", sm)
	if error != nil {
		log.Fatal(error)
	}
}
