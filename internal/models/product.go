package models

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

type ProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
}
