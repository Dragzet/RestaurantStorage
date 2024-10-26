package repository

import (
	"RestaurantStorage/internal/models"
	"context"
	"fmt"
)

func (d *DB) AddProduct(ctx context.Context, name string, amount int) error {
	stmt := `INSERT INTO product (name, amount) VALUES (?, ?)`

	return nil
}

func (d *DB) ChangeProduct(ctx context.Context, name string, amount int) error {
	stmt := `UPDATE products SET amount = $1 WHERE name = $2`
	err := d.findProduct(ctx, name)
	if err != nil {
		return err
	}
	_, err = d.Client.Query(ctx, stmt, name, amount)
	if err != nil {
		return fmt.Errorf("Error in ProductDB: ChangeProduct: %s", err.Error())
	}
	return nil
}

func (d *DB) findProduct(ctx context.Context, name string) error {
	errorStatement := "ProductDB: findProduct"
	stmt := `SELECT * FROM products WHERE name = $1`

	var productItem models.ProductModel
	err := d.Client.QueryRow(ctx, stmt, name).Scan(&productItem.Name, &productItem.Amount, &productItem.UpdateTime)
	if err != nil {
		return fmt.Errorf("Error in %s: %s", errorStatement, err.Error())
	}
	return nil
}
