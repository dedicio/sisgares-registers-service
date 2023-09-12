package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/dedicio/sisgares-registers-service/internal/registers/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_NAME = "registers"
	DB_HOST = "mysql"
	DB_USER = "root"
	DB_PASS = "root"
	DB_PORT = "3306"
)

func main() {
	dbUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_USER,
		DB_PASS,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)
	db, err := sql.Open("mysql", dbUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection is been established succesfully")
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	router := chi.NewRouter()
	routes := routes.NewRoutes(db)
	router.Use(middleware.Logger)
	router.Mount("/", routes.Routes())

	http.ListenAndServe(":3000", router)
}

// func apiVersionCtx(version string) func(next http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			r = r.WithContext(context.WithValue(r.Context(), "api.version", version))
// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }
