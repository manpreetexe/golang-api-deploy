package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go-api/internal/handler"
)

func main() {
	r := chi.NewRouter()

	// Routes
	r.Get("/", handler.HelloWorld)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message":"pong"}`))
	})

	log.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
