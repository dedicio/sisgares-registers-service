package routes

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/controllers"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
	"github.com/go-chi/chi/v5"
)

type CategoryRoutes struct {
	Controller controllers.CategoryController
}

func NewCategoryRoutes(repository entity.CategoryRepository) *CategoryRoutes {
	return &CategoryRoutes{
		Controller: *controllers.NewCategoryController(repository),
	}
}

func (cr CategoryRoutes) Routes() chi.Router {
	router := chi.NewRouter()

	router.Route("/", func(router chi.Router) {
		router.Get("/", cr.Controller.FindAll)
		router.Post("/", cr.Controller.Create)

		router.Route("/{id}", func(router chi.Router) {
			router.Get("/", cr.Controller.FindById)
			router.Delete("/", cr.Controller.Delete)
			router.Put("/", cr.Controller.Update)
			router.Get("/products", cr.Controller.FindProductsByCategoryId)
		})
	})

	return router
}
