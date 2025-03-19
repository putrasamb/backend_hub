package dto

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Status      string  `json:"status" validate:"required"`
	CreatedBy   string  `json:"created_by" validate:"required"`
}

type UpdateProductRequest struct {
	ID          int     `json:"id" validate:"required"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Status      string  `json:"status"`
	UpdatedBy   string  `json:"updated_by" validate:"required"`
}

type DeleteProductRequest struct {
	ID int `json:"id" validate:"required"`
}

type GetProductRequest struct {
	ID int `json:"id" validate:"required"`
}
