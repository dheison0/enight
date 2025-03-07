package models

type PurchaseItemRequest struct {
	ItemID   int `json:"id"`
	SizeID   int `json:"size_id"`
	Quantity int `json:"quantity"`
}

type PurchaseItemResponse struct {
	Product
	SizeName   string  `json:"size_name"`
	UnityPrice float64 `json:"unity_price"`
	Quantity   int     `json:"quantity"`
}

type PurchaseRequest struct {
	ID          int                   `json:"id"`
	ClientPhone string                `json:"client_phone"`
	Products    []PurchaseItemRequest `json:"products"`
	Stage       string                `json:"stage"`
}

type PurchaseListItem struct {
	ID        int            `json:"id"`
	Client    ClientResponse `json:"client"`
	Price     float64        `json:"price"`
	Stage     string         `json:"stage"`
	CreatedAt int64          `json:"created_at"`
}

type PurchaseResponse struct {
	PurchaseListItem
	Products []PurchaseItemResponse `json:"products"`
}
