package router

import (
	"github.com/go-chi/chi/v5"
)

type Router struct{}

func (router Router) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		r.Route("/products", func(r chi.Router) {
			r.Get("/", ProductService.ListAll)
		})
	})

	return r
}
