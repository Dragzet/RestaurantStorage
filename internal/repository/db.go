package repository

import "RestaurantStorage/internal/storage/PostgreSQL"

type DB struct {
	Client PostgreSQL.Client
}

func NewDB(client PostgreSQL.Client) *DB {
	return &DB{Client: client}
}
