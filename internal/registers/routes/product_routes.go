package routes

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/controllers"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
	"github.com/go-chi/chi/v5"
)

type ProductRoutes struct {
	Controller controllers.ProductController
}

func NewProductRoutes(repository entity.ProductRepository) *ProductRoutes {
	return &ProductRoutes{
		Controller: *controllers.NewProductController(repository),
	}
}

func (pr ProductRoutes) Routes() chi.Router {
	router := chi.NewRouter()

	router.Route("/", func(router chi.Router) {
		router.Post("/", pr.Controller.Create)

		router.Route("/{id}", func(router chi.Router) {
			router.Get("/", pr.Controller.FindById)
			router.Delete("/", pr.Controller.Delete)
			router.Put("/", pr.Controller.Update)
		})
	})

	return router
}
