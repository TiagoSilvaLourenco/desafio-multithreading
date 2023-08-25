package main

import (
	"net/http"

	"github.com/TiagoSilvaLourenco/desafio-multithreading/configs"
	"github.com/TiagoSilvaLourenco/desafio-multithreading/internal/infra/webserver/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/address", func(r chi.Router) {
		r.Get("/{cep}", handlers.GetAddress) // This route should work
	})

	http.ListenAndServe(config.WebServerPort, r)
}
