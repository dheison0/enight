package models

type Location struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Distance int    `json:"distance"`
}
