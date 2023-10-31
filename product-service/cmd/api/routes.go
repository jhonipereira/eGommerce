package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.Logger)

	mux.Route("/products", func(r chi.Router) {
		r.Get("/", app.GetAllProducts)
		r.Get("/{name}", app.GetProductByName)
		r.Get("/{id:[0-9]+}", app.GetProductByID)
		r.Put("/", app.UpdateProduct)
		r.Delete("/{id:[0-9]+}", app.DeleteProductByID)
		r.Post("/", app.InsertProduct)
	})

	return mux
}
