package router

import (
	"github.com/go-chi/chi/v5"
	"goBitly/internal/handler"
)

func SetUpRouter(urlHandler *handler.URLHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/goBitly/health", urlHandler.CheckHealth) // done
	r.Post("/goBitly/create", urlHandler.CreateShortURLHandler) // done
	r.Get("/goBitly/urls", urlHandler.GetByOriginalURLHandler) 
	r.Get("/goBitly/{shortURL}/clicks", urlHandler.GetClickCountHandler) // done
	r.Get("/goBitly/{shortURL}", urlHandler.GetByShortURLHandler)
	r.Delete("/goBitly/{shortURL}", urlHandler.DeleteURLHandler)
	r.Get("/{shortURL}", urlHandler.RedirectHandler) // done

	return r
}
