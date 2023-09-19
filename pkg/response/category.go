package response

import (
	"net/http"

	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
)

type CategoryResponse struct {
	*dto.CategoryResponseDto
}

func (cr *CategoryResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewCategoryResponse(category *dto.CategoryResponseDto) *CategoryResponse {
	return &CategoryResponse{category}
}

type CategoriesResponse struct {
	Categories []*dto.CategoryResponseDto
}

func (cr *CategoriesResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewCategoriesResponse(categories []*dto.CategoryResponseDto) *CategoriesResponse {
	return &CategoriesResponse{categories}
}
