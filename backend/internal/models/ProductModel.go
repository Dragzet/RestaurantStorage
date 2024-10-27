package models

type ProductModel struct {
	ID     int    `json:"id" swaggerignore:"true"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Image  string `json:"image"`
}
