package routes

import (
	"database/sql"

	"github.com/dedicio/sisgares-registers-service/internal/registers/infra/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

var tokenAuth *jwtauth.JWTAuth

type Routes struct {
	DB *sql.DB
}

func NewRoutes(db *sql.DB) *Routes {
	return &Routes{
		DB: db,
	}
}

func (routes Routes) Routes() chi.Router {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(jwtauth.Verifier(tokenAuth))
	router.Use(jwtauth.Authenticator)

	productRepository := repository.NewProductRepositoryMysql(routes.DB)
	categoryRepository := repository.NewCategoryRepositoryMysql(routes.DB)
	positionRepository := repository.NewPositionRepositoryMysql(routes.DB)
	groupRepository := repository.NewGroupRepositoryMysql(routes.DB)

	router.Route("/v1", func(router chi.Router) {
		router.Mount("/products", NewProductRoutes(productRepository).Routes())
		router.Mount("/categories", NewCategoryRoutes(categoryRepository).Routes())
		router.Mount("/positions", NewPositionRoutes(positionRepository).Routes())
		router.Mount("/groups", NewGroupRoutes(groupRepository).Routes())
	})

	return router
}
