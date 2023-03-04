package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World!")
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")

		body, error := ioutil.ReadAll(r.Body)
		if error != nil {
			log.Println("Error reading body ", error)
			http.Error(rw, "OOps!", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s\n", body)
	})

	log.Println("Starting Server")
	error := http.ListenAndServe(":9090", nil)
	if error != nil {
		log.Fatal(error)
	}
}
