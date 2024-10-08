package database

import (
	"api/models"
	"encoding/json"
	"errors"
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

func GetAllPurchases(offset, limit int) (items []models.PurchaseListItem, err error) {
	rows, err := db.Query("SELECT id, client, price, stage FROM purchases;")
	if err != nil {
		return items, err
	}
	item := models.PurchaseListItem{}
	var client string
	for rows.Next() {
		if err = rows.Scan(&item.ID, &client, &item.Price, &item.Stage); err != nil {
			break
		} else if err = json.Unmarshal([]byte(client), &item.Client); err != nil {
			break
		}
		items = append(items, item)
	}
	return items, err
}

func GetPurchase(id int) (purchase models.PurchaseResponse, err error) {
	purchase.ID = id
	var client, products string
	err = db.QueryRow("SELECT client, products, price, stage FROM purchases WHERE id = ?;", id).
		Scan(&client, &products, &purchase.Price, &purchase.Stage)
	if err != nil {
		err = errors.New("failed to find item! " + err.Error())
	} else if err = json.Unmarshal([]byte(client), &purchase.Client); err != nil {
		err = errors.New("failed to parse client structure! " + err.Error())
	} else if err = json.Unmarshal([]byte(products), &purchase.Products); err != nil {
		err = errors.New("failed to parse products structure! " + err.Error())
	}
	return purchase, err
}
