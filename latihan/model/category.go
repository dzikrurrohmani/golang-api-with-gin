package model

type Category struct {
	CategoryId   string `json:"id" binding:"required"`
	CategoryName string `json:"name" binding:"required"`
}
