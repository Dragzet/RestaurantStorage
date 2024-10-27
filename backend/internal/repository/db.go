package repository

import "RestaurantStorage/internal/storage/PostgreSQL"

type DB struct {
	Client PostgreSQL.Client
	ProductStorage
}

func NewDB(client PostgreSQL.Client) *DB {
	return &DB{Client: client}
}
