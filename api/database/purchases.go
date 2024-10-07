package database

import (
	"api/models"
	"encoding/json"
)

func CreatePurchase(purchase *models.PurchaseRequest) (models.PurchaseResponse, error) {
	result := models.PurchaseResponse{}
	purchaseItems, err := getPurchaseItems(purchase.Products)
	if err != nil {
		return result, err
	}
	result.Products = purchaseItems
	client, err := GetClient(purchase.ClientPhone)
	if err != nil {
		return result, err
	}
	result.Client = *client
	clientJSON, _ := json.Marshal(result.Client)
	productsJSON, _ := json.Marshal(result.Products)
	result.Price = calculateTotalPrice(result.Products)
	err = db.QueryRow(
		"INSERT INTO purchases(client, products, price) VALUES (?, ?, ?) RETURNING id, stage;",
		string(clientJSON), string(productsJSON), result.Price,
	).Scan(&result.ID, &result.Stage)
	return result, err
}

func calculateTotalPrice(items []models.PurchaseItemResponse) float64 {
	total := 0.0
	for _, item := range items {
		total += item.UnityPrice * float64(item.Quantity)
	}
	return total
}

func getPurchaseItems(items []models.PurchaseItemRequest) ([]models.PurchaseItemResponse, error) {
	purchaseItems := []models.PurchaseItemResponse{}
	var toInsert models.PurchaseItemResponse
	for _, item := range items {
		toInsert.ID = item.ItemID
		toInsert.Quantity = item.Quantity
		err := db.QueryRow(
			`SELECT products.name, products.cover_url, product_sizes.name, product_sizes.price
			 FROM products INNER JOIN product_sizes ON products.id = product_sizes.product_id
			 WHERE products.id = ? AND product_sizes.id = ?;`,
			item.ItemID, item.SizeID,
		).Scan(&toInsert.Name, &toInsert.CoverURL, &toInsert.SizeName, &toInsert.UnityPrice)
		if err != nil {
			return nil, err
		}
		purchaseItems = append(purchaseItems, toInsert)
	}
	return purchaseItems, nil
}

func GetAllPurchases(offset, limit int) ([]models.PurchaseResponse, error) {
	result := []models.PurchaseResponse{}
	rows, err := db.Query("SELECT id, client, products, price, stage FROM purchases;")
	if err != nil {
		return result, err
	}
	item := models.PurchaseResponse{}
	for rows.Next() {
		var client string
		var products string
		if err = rows.Scan(&item.ID, &client, &products, &item.Price, &item.Stage); err != nil {
			return result, err
		}
		if err = json.Unmarshal([]byte(client), &item.Client); err != nil {
			break
		}
		if err = json.Unmarshal([]byte(products), &item.Products); err != nil {
			break
		}
		result = append(result, item)
	}
	return result, nil
}
