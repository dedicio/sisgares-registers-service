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

type CategoryController struct {
	Repository entity.CategoryRepository
}

func NewCategoryController(categoryRepository entity.CategoryRepository) *CategoryController {
	return &CategoryController{
		Repository: categoryRepository,
	}
}

func (cc *CategoryController) FindAll(w http.ResponseWriter, r *http.Request) {
	categories, err := usecase.NewListCategoriesUseCase(cc.Repository).Execute()

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, httpResponsePkg.NewCategoriesResponse(categories))
}

func (cc *CategoryController) FindById(w http.ResponseWriter, r *http.Request) {
	categoryId := chi.URLParam(r, "id")
	category, err := usecase.NewFindCategoryByIdUseCase(cc.Repository).Execute(categoryId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrNotFound(err, "Categoria"))
		return
	}

	render.Render(w, r, httpResponsePkg.NewCategoryResponse(category))
}

func (cc *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	payload := json.NewDecoder(r.Body)
	category := dto.CategoryDto{}
	err := payload.Decode(&category)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	categorySaved, err := usecase.NewCreateCategoryUseCase(cc.Repository).Execute(category)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.CategoryResponseDto{
		ID:   categorySaved.ID,
		Name: categorySaved.Name,
	}
	render.Render(w, r, httpResponsePkg.NewCategoryResponse(output))
}

func (cc *CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	payload := json.NewDecoder(r.Body)
	category := dto.CategoryDto{}
	err := payload.Decode(&category)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	err = usecase.NewUpdateCategoryUseCase(cc.Repository).Execute(category)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.CategoryResponseDto{
		ID:   category.ID,
		Name: category.Name,
	}
	render.Render(w, r, httpResponsePkg.NewCategoryResponse(output))
}

func (cc *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	categoryId := chi.URLParam(r, "id")
	err := usecase.NewDeleteCategoryUseCase(cc.Repository).Execute(categoryId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, nil)
}

func (cc *CategoryController) FindProductsByCategoryId(w http.ResponseWriter, r *http.Request) {
	categoryId := chi.URLParam(r, "id")
	products, err := usecase.NewListProductsByCategoryUseCase(cc.Repository).Execute(categoryId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, httpResponsePkg.NewProductsResponse(products))
}
