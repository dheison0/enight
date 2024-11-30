package database

import (
	"crypto/sha256"
	"fmt"
	"server/models"
)

func GetSettings() (*models.Settings, error) {
	settings := models.Settings{}
	err := db.QueryRow("SELECT shipping_price, password_hash FROM settings;").
		Scan(&settings.ShippingPrice, &settings.PasswordHash)
	return &settings, err
}

func CheckPassword(password string) bool {
	settings, err := GetSettings()
	if err != nil {
		return false
	}
	hash := sha256.Sum256([]byte(password))
	return settings.PasswordHash == fmt.Sprintf("%x", hash)
}

func SetPassword(password string) error {
	hash := sha256.Sum256([]byte(password))
	_, err := db.Exec("UPDATE settings SET password_hash = ?;", fmt.Sprintf("%x", hash))
	return err
}

func SetShippingPrice(price float64) error {
	_, err := db.Exec("UPDATE settings SET shipping_price = ?;", price)
	return err
}
