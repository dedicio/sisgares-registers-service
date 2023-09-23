package response

import (
	"net/http"

	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
)

type PositionResponse struct {
	*dto.PositionResponseDto
}

func (pr *PositionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewPositionResponse(product *dto.PositionResponseDto) *PositionResponse {
	return &PositionResponse{product}
}

type PositionsResponse struct {
	Positions []*dto.PositionResponseDto
}

func (pr *PositionsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewPositionsResponse(products []*dto.PositionResponseDto) *PositionsResponse {
	return &PositionsResponse{products}
}
