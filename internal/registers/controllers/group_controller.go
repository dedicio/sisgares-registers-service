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

type GroupController struct {
	Repository entity.GroupRepository
}

func NewGroupController(groupRepository entity.GroupRepository) *GroupController {
	return &GroupController{
		Repository: groupRepository,
	}
}

func (gc *GroupController) FindAll(w http.ResponseWriter, r *http.Request) {
	companyID := r.Header.Get("X-Company-ID")
	groups, err := usecase.NewListGroupsUseCase(gc.Repository).Execute(companyID)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, httpResponsePkg.NewGroupsResponse(groups))
}

func (gc *GroupController) FindById(w http.ResponseWriter, r *http.Request) {
	groupId := chi.URLParam(r, "id")
	group, err := usecase.NewFindGroupByIdUseCase(gc.Repository).Execute(groupId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrNotFound(err, "Grupo"))
		return
	}

	render.Render(w, r, httpResponsePkg.NewGroupResponse(group))
}

func (gc *GroupController) Create(w http.ResponseWriter, r *http.Request) {
	companyID := r.Header.Get("X-Company-ID")
	payload := json.NewDecoder(r.Body)
	group := dto.GroupDto{}
	err := payload.Decode(&group)
	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	group.CompanyId = companyID
	groupSaved, err := usecase.NewCreateGroupUseCase(gc.Repository).Execute(group)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.GroupResponseDto{
		ID:   groupSaved.ID,
		Name: groupSaved.Name,
	}

	render.Render(w, r, httpResponsePkg.NewGroupResponse(output))
}

func (gc *GroupController) Update(w http.ResponseWriter, r *http.Request) {
	payload := json.NewDecoder(r.Body)
	group := dto.GroupDto{}
	err := payload.Decode(&group)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	err = usecase.NewUpdateGroupUseCase(gc.Repository).Execute(group)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.GroupResponseDto{
		ID:   group.ID,
		Name: group.Name,
	}

	render.Render(w, r, httpResponsePkg.NewGroupResponse(output))
}

func (gc *GroupController) Delete(w http.ResponseWriter, r *http.Request) {
	groupId := chi.URLParam(r, "id")
	err := usecase.NewDeleteGroupUseCase(gc.Repository).Execute(groupId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, nil)
}
