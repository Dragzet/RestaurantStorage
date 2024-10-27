package repository

import (
	"RestaurantStorage/internal/models"
	"context"
	"fmt"
)

func (d *DB) AddProduct(ctx context.Context, product models.ProductModel) error {
	stmt := `
		INSERT INTO products
			(name, amount, image)
		VALUES
			($1, $2, $3)`
	err := d.findProduct(ctx, product.Name)
	if err == nil {
		return fmt.Errorf("ProductDB: AddProduct: Product already exists")
	}

	_, err = d.Client.Exec(ctx, stmt, product.Name, product.Amount, product.Image)
	if err != nil {
		return fmt.Errorf("Error in ProductDB: AddProduct: %s", err.Error())
	}

	return nil
}

func (d *DB) GetProduct(ctx context.Context, name string) (*models.ProductModel, error) {
	stmt := `SELECT * FROM products WHERE name = $1`

	err := d.findProduct(ctx, name)
	if err != nil {
		return nil, err
	}

	var productItem models.ProductModel
	err = d.Client.QueryRow(ctx, stmt, name).Scan(&productItem.ID, &productItem.Name, &productItem.Amount, &productItem.Image)
	if err != nil {
		return nil, fmt.Errorf("Error in ProductDB: GetProduct: %s", err.Error())
	}
	return &productItem, nil
}

func (d *DB) ChangeProduct(ctx context.Context, name string, amount int) error {
	stmt := `UPDATE products SET amount = $2 WHERE name = $1`
	err := d.findProduct(ctx, name)
	if err != nil {
		return err
	}

	//lastProduct, err := d.GetProduct(ctx, name)
	//if err != nil {
	//	return fmt.Errorf("Error in ProductDB: ChangeProduct: %s", err.Error())
	//}
	//
	//amount = lastProduct.Amount + amount
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
	err := d.Client.QueryRow(ctx, stmt, name).Scan(&productItem.ID, &productItem.Name, &productItem.Amount, &productItem.Image)
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("Error in %s: %s", errorStatement, err.Error())
	}
	return nil
}

func (d *DB) DeleteProduct(ctx context.Context, name string) error {
	stmt := `DELETE FROM products WHERE name = $1`
	err := d.findProduct(ctx, name)
	if err != nil {
		return err
	}

	_, err = d.Client.Query(ctx, stmt, name)
	if err != nil {
		return fmt.Errorf("Error in ProductDB: DeleteProduct: %s", err.Error())
	}
	return nil
}

func (d *DB) GetProductList(ctx context.Context) ([]*models.ProductModel, error) {
	stmt := `SELECT * FROM products`
	rows, err := d.Client.Query(ctx, stmt)
	if err != nil {
		return nil, fmt.Errorf("Error in ProductDB: GetProductList: %s", err.Error())
	}
	defer rows.Close()

	var products []*models.ProductModel
	for rows.Next() {
		var productItem models.ProductModel
		err = rows.Scan(&productItem.ID, &productItem.Name, &productItem.Amount, &productItem.Image)
		if err != nil {
			return nil, fmt.Errorf("Error in ProductDB: GetProductList: %s", err.Error())
		}
		products = append(products, &productItem)
	}
	return products, nil
}
