package router

import (
	"github.com/go-chi/chi/v5"
	"goBitly/internal/handler"
)

func SetUpRouter(urlHandler *handler.URLHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/goBitly/health", urlHandler.CheckHealth)
	r.Post("/goBitly/create", urlHandler.CreateShortURLHandler)
	r.Get("/goBitly/urls", urlHandler.GetByOriginalURLHandler)
	r.Get("/goBitly/{shortURL}/clicks", urlHandler.GetClickCountHandler)
	r.Get("/goBitly/{shortURL}", urlHandler.GetByShortURLHandler)
	r.Delete("/goBitly/{shortURL}", urlHandler.DeleteURLHandler)
	r.Get("/{shortURL}", urlHandler.RedirectHandler)

	return r
}
