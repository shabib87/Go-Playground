package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHttp(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world!")
	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		log.Println("Error reading body ", error)
		http.Error(rw, "OOps!", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s\n", body)
}
