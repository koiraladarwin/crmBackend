package postgres

import "github.com/koiraladarwin/crmbackend/models"

func (pg *Postgres) AddClient(id, companyID, name, gmail, phone string) error {
	query := `INSERT INTO clients (id, company_id, name, gmail, phone) VALUES ($1, $2, $3, $4, $5)`
	_, err := pg.db.Exec(query, id, companyID, name, gmail, phone)
	return err
}

func (pg *Postgres) GetClientByID(id string) (models.Client, error) {
	var client models.Client
	query := `SELECT id, company_id, name, gmail, phone FROM clients WHERE id = $1`
	err := pg.db.QueryRow(query, id).Scan(&client.ID, &client.CompanyID, &client.Name, &client.Gmail, &client.Phone)
	return client, err
}

func (pg *Postgres) GetClientsByCompanyID(companyID string) ([]models.Client, error) {
	query := `SELECT id, company_id, name, gmail, phone FROM clients WHERE company_id = $1`
	rows, err := pg.db.Query(query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ID, &client.CompanyID, &client.Name, &client.Gmail, &client.Phone); err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}

