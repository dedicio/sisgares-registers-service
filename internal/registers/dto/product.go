package dto

type ProductDto struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	CategoryId  string  `json:"category_id"`
	CompanyId   string  `json:"company_id"`
}

type ProductResponseDto struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	CategoryId  string  `json:"category_id"`
}

type CategoryDto struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CompanyId string `json:"company_id"`
}

type CategoryResponseDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
