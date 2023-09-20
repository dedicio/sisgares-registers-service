package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
	usecase "github.com/dedicio/sisgares-registers-service/internal/registers/usecase/product"
	httpResponsePkg "github.com/dedicio/sisgares-registers-service/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ProductController struct {
	Repository entity.ProductRepository
}

func NewProductController(productRepository entity.ProductRepository) *ProductController {
	return &ProductController{
		Repository: productRepository,
	}
}

func (pc *ProductController) FindAll(w http.ResponseWriter, r *http.Request) {
	products, err := usecase.NewListProductsUseCase(pc.Repository).Execute()

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, httpResponsePkg.NewProductsResponse(products))
}

func (pc *ProductController) FindById(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "id")
	product, err := usecase.NewFindProductByIdUseCase(pc.Repository).Execute(productId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrNotFound(err, "Produto"))
		return
	}

	render.Render(w, r, httpResponsePkg.NewProductResponse(product))
}

func (pc *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	payload := json.NewDecoder(r.Body)
	product := dto.ProductDto{}
	err := payload.Decode(&product)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	productSaved, err := usecase.NewCreateProductUseCase(pc.Repository).Execute(product)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.ProductResponseDto{
		ID:          productSaved.ID,
		Name:        productSaved.Name,
		Description: productSaved.Description,
		Price:       productSaved.Price,
		Image:       productSaved.Image,
		CategoryId:  productSaved.CategoryId,
	}
	render.Render(w, r, httpResponsePkg.NewProductResponse(output))
}

func (pc *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "id")
	err := usecase.NewDeleteProductUseCase(pc.Repository).Execute(productId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, nil)
}

func (pc *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	payload := json.NewDecoder(r.Body)
	product := dto.ProductDto{}
	err := payload.Decode(&product)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	err = usecase.NewUpdateProductUseCase(pc.Repository).Execute(product)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.ProductResponseDto{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Image:       product.Image,
		CategoryId:  product.CategoryId,
	}

	render.Render(w, r, httpResponsePkg.NewProductResponse(output))
}
