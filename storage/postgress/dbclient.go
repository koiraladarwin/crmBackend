package postgres

import "github.com/koiraladarwin/crmbackend/models"


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

func (pg *Postgres) GetClientsByCompanyID(companyID string) ([]models.Client, error) {
	query := `
		SELECT DISTINCT c.id, c.name, c.gmail, c.phone
		FROM clients c
		JOIN client_process cp ON cp.client_id = c.id
		JOIN employees e ON cp.assigned_employee_id = e.id
		WHERE e.company_id = $1
	`
	rows, err := pg.db.Query(query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ID, &client.Name, &client.Gmail, &client.Phone); err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}

