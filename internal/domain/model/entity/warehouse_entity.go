package entity

type Product struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	UpdatedBy   string  `json:"updated_by"`
	CreatedBy   string  `json:"created_by"`
}

func (p *Product) TableName() string {
	return "products"
}
