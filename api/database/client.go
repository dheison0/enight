package database

import "api/models"

func CreateClient(client *models.ClientDatabase) error {
	_, err := db.Exec(
		"INSERT INTO clients(phone, name, location_id) VALUES (?, ?, ?);",
		client.Phone, client.Name, client.LocationID,
	)
	return err
}

func GetAllClients() ([]models.ClientDatabase, error) {
	rows, err := db.Query("SELECT phone, name, location_id FROM clients;")
	if err != nil {
		return nil, err
	}
	var client models.ClientDatabase
	clients := []models.ClientDatabase{}
	for rows.Next() {
		err := rows.Scan(&client.Phone, &client.Name, &client.LocationID)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func GetClient(phone string) (*models.ClientResponse, error) {
	clientResponse := models.ClientResponse{Phone: phone}
	err := db.QueryRow(
		`SELECT clients.name, locations.id, locations.name, locations.distance
		 FROM clients INNER JOIN locations ON locations.id = clients.location_id
		 WHERE clients.phone = ?;`,
		phone,
	).Scan(
		&clientResponse.Name, &clientResponse.Location.ID,
		&clientResponse.Location.Name, &clientResponse.Location.Distance,
	)
	if err != nil {
		return nil, err
	}
	return &clientResponse, nil
}

func DeleteClient(client *models.ClientDatabase) error {
	_, err := db.Exec("DELETE FROM clients WHERE phone = ?;", client.Phone)
	return err
}
