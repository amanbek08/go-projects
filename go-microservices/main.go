package main

import (
	"go-projects/go-microservices/handlers"
	"log"
	"net/http"
	"os"
)

// "fmt"
// "go-projects/structures"

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":9090", sm)

}
