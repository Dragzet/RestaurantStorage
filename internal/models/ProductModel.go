package models

import "time"

type ProductModel struct {
	ID         int       `json:"id" swaggerignore:"true"`
	Name       string    `json:"name"`
	Amount     int       `json:"amount"`
	UpdateTime time.Time `json:"update"`
}
