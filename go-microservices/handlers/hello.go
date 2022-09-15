package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println("Hello World!")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Error", http.StatusBadRequest)
	}

	fmt.Fprintf(rw, "Hello %s", d)
}
