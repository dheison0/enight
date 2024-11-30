package database

import "server/models"

func CreateLocation(location *models.Location) error {
	return db.QueryRow(
		"INSERT INTO locations(name, distance) VALUES (?, ?) RETURNING id;",
		location.Name, location.Distance,
	).Scan(&location.ID)
}

func DeleteLocation(location *models.Location) error {
	_, err := db.Exec("DELETE FROM locations WHERE id = ?;", location.ID)
	return err
}

func GetAllLocations() ([]models.Location, error) {
	rows, err := db.Query("SELECT id, name, distance FROM locations;")
	if err != nil {
		return nil, err
	}
	var location models.Location
	locations := []models.Location{}
	for rows.Next() {
		if err := rows.Scan(&location.ID, &location.Name, &location.Distance); err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}
	return locations, nil
}

func GetLocation(id int) (*models.Location, error) {
	location := models.Location{ID: id}
	err := db.QueryRow(
		"SELECT name, distance FROM locations WHERE id = ?;",
		id,
	).Scan(&location.Name, &location.Distance)
	return &location, err
}
