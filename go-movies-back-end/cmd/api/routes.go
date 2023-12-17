package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {
	// создаем маршрутизатор
	mux := chi.NewRouter()
	// обрабатываем возможную панику
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)
	mux.Get("/", app.Home)
	mux.Get("/movies", app.AllMovies)
	return mux
}
