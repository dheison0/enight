package models

type Settings struct {
	ShippingPrice float64 `json:"shipping_price"`
	PasswordHash  string  `json:"password_hash"`
}
