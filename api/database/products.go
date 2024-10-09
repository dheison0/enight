package database

import (
	"api/models"
)

func CreateProduct(product *models.Product) error {
	return db.QueryRow(
		"INSERT INTO products(name, description, cover_url) VALUES (?, ?, ?) RETURNING id;",
		product.Name, product.Description, product.CoverURL,
	).Scan(&product.ID)
}

func DeleteProduct(id int) error {
	_, err := db.Exec("DELETE FROM products WHERE id = ?;", id)
	return err
}

func AddProductSize(productSize *models.ProductSize) error {
	return db.QueryRow(
		"INSERT INTO product_sizes(name, price, product_id) VALUES (?, ?, ?) RETURNING id;",
		productSize.Name, productSize.Price, productSize.ProductID,
	).Scan(&productSize.ID)
}

func DeleteProductSize(productSize *models.ProductSize) error {
	_, err := db.Exec("DELETE FROM product_sizes WHERE id = ?;", productSize.ID)
	return err
}

func GetAllProducts() ([]models.Product, error) {
	rows, err := db.Query("SELECT id, name, cover_url FROM products;")
	if err != nil {
		return nil, err
	}
	var product models.Product
	products := []models.Product{}
	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.CoverURL)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func GetProduct(id int) (*models.ProductResponse, error) {
	product := models.ProductResponse{}
	product.ID = id
	err := db.QueryRow("SELECT name, description, cover_url FROM products WHERE id = ?", id).
		Scan(&product.Name, &product.Description, &product.CoverURL)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(
		"SELECT id, name, price FROM product_sizes WHERE product_id = ?;",
		product.ID,
	)
	if err != nil {
		return nil, err
	}
	var size models.ProductSize
	for rows.Next() {
		err = rows.Scan(&size.ID, &size.Name, &size.Price)
		if err != nil {
			return nil, err
		}
		product.Sizes = append(product.Sizes, size)
	}
	return &product, nil
}
