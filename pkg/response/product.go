package response

import (
	"net/http"

	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
)

type ProductResponse struct {
	*dto.ProductResponseDto
}

func (pr *ProductResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewProductResponse(product *dto.ProductResponseDto) *ProductResponse {
	return &ProductResponse{product}
}
