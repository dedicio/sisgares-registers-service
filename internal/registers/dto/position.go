package dto

type PositionDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	GroupId     string `json:"group_id"`
	CompanyId   string `json:"company_id"`
}

type PositionResponseDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	GroupId     string `json:"group_id"`
}

type GroupDto struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CompanyId string `json:"company_id"`
}

type GroupResponseDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
