package response

import (
	"net/http"

	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
)

type GroupResponse struct {
	*dto.GroupResponseDto
}

func (cr *GroupResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewGroupResponse(category *dto.GroupResponseDto) *GroupResponse {
	return &GroupResponse{category}
}

type GroupsResponse struct {
	Groups []*dto.GroupResponseDto `json:"items"`
}

func (cr *GroupsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewGroupsResponse(categories []*dto.GroupResponseDto) *GroupsResponse {
	return &GroupsResponse{categories}
}
