package routes

import (
	"database/sql"

	"github.com/dedicio/sisgares-registers-service/internal/registers/infra/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Routes struct {
	DB *sql.DB
}

func NewRoutes(db *sql.DB) *Routes {
	return &Routes{
		DB: db,
	}
}

func (routes Routes) Routes() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	productRepository := repository.NewProductRepositoryPostgresql(routes.DB)
	categoryRepository := repository.NewCategoryRepositoryPostgresql(routes.DB)
	positionRepository := repository.NewPositionRepositoryPostgresql(routes.DB)
	groupRepository := repository.NewGroupRepositoryPostgresql(routes.DB)

	router.Route("/v1", func(router chi.Router) {
		router.Mount("/products", NewProductRoutes(productRepository).Routes())
		router.Mount("/categories", NewCategoryRoutes(categoryRepository).Routes())
		router.Mount("/positions", NewPositionRoutes(positionRepository).Routes())
		router.Mount("/groups", NewGroupRoutes(groupRepository).Routes())
	})

	return router
}
