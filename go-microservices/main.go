package main

import (
	"github.com/gofiber/fiber/v2"
	// "context"
	"go-projects/go-microservices/data"
	"log"

	"github.com/gofiber/fiber/v2/middleware/logger"
	// "net/http"
	// "os"
	// "os/signal"
	// "time"
	// "github.com/gorilla/mux"
)

func setupRoutes(app *fiber.App) {
	app.Get("/books", data.GetBooks)
	app.Get("/book/:id", data.GetBook)
	app.Post("/book", data.AddBook)
}

func main() {

	app := fiber.New()

	app.Use(
		logger.New(logger.Config{
			Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		}), // add Logger middleware
	)

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

	// l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	// ph := handlers.NewProducts(l)

	// sm := mux.NewRouter()

	// getRouter := sm.Methods(http.MethodGet).Subrouter()
	// getRouter.HandleFunc("/", ph.GetProducts)

	// putRouter := sm.Methods(http.MethodPut).Subrouter()
	// putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	// putRouter.Use(ph.MiddlewareProductValidation)

	// postRouter := sm.Methods(http.MethodPost).Subrouter()
	// postRouter.HandleFunc("/", ph.AddProduct)
	// postRouter.Use(ph.MiddlewareProductValidation)

	// s := &http.Server{
	// 	Addr:         ":9090",
	// 	Handler:      sm,
	// 	IdleTimeout:  120 * time.Second,
	// 	ReadTimeout:  1 * time.Second,
	// 	WriteTimeout: 1 * time.Second,
	// }

	// go func() {
	// 	err := s.ListenAndServe()
	// 	if err != nil {
	// 		l.Fatal(err)
	// 	}
	// }()

	// sigChan := make(chan os.Signal)
	// signal.Notify(sigChan, os.Interrupt)
	// signal.Notify(sigChan, os.Kill)

	// sig := <-sigChan
	// l.Println("Recieved terminate, graceful shutdown", sig)

	// tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// s.Shutdown(tc)

}
