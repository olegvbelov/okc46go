package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/olegvbelov/okc46go/internal/config"
	"github.com/olegvbelov/okc46go/internal/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/services", handlers.Repo.Services)
	mux.Post("/sendEmail", handlers.Repo.SendEmail)
	mux.Get("/sitemap.xml", handlers.Repo.Sitemap)
	mux.Get("/robot.txt", handlers.Repo.Robot)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
