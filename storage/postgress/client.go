package postgres

import "github.com/koiraladarwin/crmbackend/storage/models"

func (pg *Postgres) AddClient(id, name, gmail, phone string) error {
	query := `INSERT INTO clients (id, name, gmail, phone) VALUES ($1, $2, $3, $4)`
	_, err := pg.db.Exec(query, id, name, gmail, phone)
	return err
}

func (pg *Postgres) GetClientByID(id string) (models.Client, error) {
	var client models.Client
	query := `SELECT id, name, gmail, phone FROM clients WHERE id = $1`
	err := pg.db.QueryRow(query, id).Scan(&client.ID, &client.Name, &client.Gmail, &client.Phone)
	return client, err
}
