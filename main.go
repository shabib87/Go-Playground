package main

import (
	"Go-Playground/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	log.Println("Starting Server")
	error := http.ListenAndServe(":9090", sm)
	if error != nil {
		log.Fatal(error)
	}
}
