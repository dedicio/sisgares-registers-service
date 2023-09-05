package controllers

import (
	"net/http"

	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type ProductController struct {
	Repository entity.ProductRepository
}

func NewProductController(productRepository entity.ProductRepository) *ProductController {
	return &ProductController{
		Repository: productRepository,
	}
}

func (pc *ProductController) FindById(w http.ResponseWriter, r *http.Request) {
	// products, err := usecase.NewListProductsUseCase(pc.Repository).Execute()

	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(products)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}
