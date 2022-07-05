package model

type Product struct {
	ProductId   string   `json:"id" binding:"required"`
	ProductName string   `json:"name" binding:"required"`
	Category    Category `json:"category" binding:"required"`
	IsActive    bool     `json:"is_active"`
}
