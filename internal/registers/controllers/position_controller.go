package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
	usecase "github.com/dedicio/sisgares-registers-service/internal/registers/usecase/position"
	httpResponsePkg "github.com/dedicio/sisgares-registers-service/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type PositionController struct {
	Repository entity.PositionRepository
}

func NewPositionController(positionRepository entity.PositionRepository) *PositionController {
	return &PositionController{
		Repository: positionRepository,
	}
}

func (pc *PositionController) FindAll(w http.ResponseWriter, r *http.Request) {
	companyID := r.Header.Get("X-Company-ID")
	positions, err := usecase.NewListPositionsUseCase(pc.Repository).Execute(companyID)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, httpResponsePkg.NewPositionsResponse(positions))
}

func (pc *PositionController) FindById(w http.ResponseWriter, r *http.Request) {
	positionId := chi.URLParam(r, "id")
	position, err := usecase.NewFindPositionByIdUseCase(pc.Repository).Execute(positionId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrNotFound(err, "Cargo"))
		return
	}

	render.Render(w, r, httpResponsePkg.NewPositionResponse(position))
}

func (pc *PositionController) Create(w http.ResponseWriter, r *http.Request) {
	companyID := r.Header.Get("X-Company-ID")
	payload := json.NewDecoder(r.Body)
	position := dto.PositionDto{}
	err := payload.Decode(&position)
	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	position.CompanyId = companyID
	positionSaved, err := usecase.NewCreatePositionUseCase(pc.Repository).Execute(position)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.PositionResponseDto{
		ID:          positionSaved.ID,
		Name:        positionSaved.Name,
		Description: positionSaved.Description,
		GroupId:     positionSaved.GroupId,
	}

	render.Render(w, r, httpResponsePkg.NewPositionResponse(output))
}

func (pc *PositionController) Update(w http.ResponseWriter, r *http.Request) {
	payload := json.NewDecoder(r.Body)
	position := dto.PositionDto{}
	err := payload.Decode(&position)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	err = usecase.NewUpdatePositionUseCase(pc.Repository).Execute(position)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.PositionResponseDto{
		ID:          position.ID,
		Name:        position.Name,
		Description: position.Description,
		GroupId:     position.GroupId,
	}

	render.Render(w, r, httpResponsePkg.NewPositionResponse(output))
}

func (pc *PositionController) Delete(w http.ResponseWriter, r *http.Request) {
	positionId := chi.URLParam(r, "id")
	err := usecase.NewDeletePositionUseCase(pc.Repository).Execute(positionId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, nil)
}
