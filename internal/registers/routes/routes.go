package routes

import (
	"database/sql"

	"github.com/dedicio/sisgares-registers-service/internal/registers/infra/repository"
	"github.com/go-chi/chi/v5"
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
	productRepository := repository.NewProductRepositoryMysql(routes.DB)

	router.Route("/v1", func(router chi.Router) {
		router.Mount("/products", NewProductRoutes(productRepository).Routes())
	})

	return router
}
