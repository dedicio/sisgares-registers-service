package routes

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/controllers"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
	"github.com/go-chi/chi/v5"
)

type GroupRoutes struct {
	Controller controllers.GroupController
}

func NewGroupRoutes(repository entity.GroupRepository) *GroupRoutes {
	return &GroupRoutes{
		Controller: *controllers.NewGroupController(repository),
	}
}

func (pr GroupRoutes) Routes() chi.Router {
	router := chi.NewRouter()

	router.Route("/", func(router chi.Router) {
		router.Get("/", pr.Controller.FindAll)
		router.Post("/", pr.Controller.Create)

		router.Route("/{id}", func(router chi.Router) {
			router.Get("/", pr.Controller.FindById)
			router.Delete("/", pr.Controller.Delete)
			router.Put("/", pr.Controller.Update)
		})
	})

	return router
}
