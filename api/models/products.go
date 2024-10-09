package models

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	CoverURL    string `json:"cover_url,omitempty"`
}

type ProductSize struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	ProductID int     `json:"product_id"`
}

type ProductResponse struct {
	Product
	Sizes []ProductSize `json:"sizes"`
}
