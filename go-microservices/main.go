package main

import (
	"log"
	"net/http"
	"os"
)

// "fmt"
// "go-projects/structures"

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.newHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", sm)

}
