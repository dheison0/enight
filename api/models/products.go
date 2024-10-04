package models

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CoverURL string `json:"cover_url"`
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
