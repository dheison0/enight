package models

type ClientDatabase struct {
	Phone      string `json:"phone"`
	Name       string `json:"name"`
	LocationID int    `json:"location_id"`
}

type ClientResponse struct {
	Phone    string   `json:"phone"`
	Name     string   `json:"name"`
	Location Location `json:"location"`
}
