package router

import (
	"github.com/go-chi/chi/v5"
)

type ProductRoutes struct{}

func (pr ProductRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	return r
}
