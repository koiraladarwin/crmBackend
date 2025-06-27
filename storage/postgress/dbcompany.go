package postgres

import "github.com/koiraladarwin/crmbackend/models"

func (pg *Postgres) AddCompany(id, name string) error {
	query := `INSERT INTO companies (id, name) VALUES ($1, $2)`
	_, err := pg.db.Exec(query, id, name)
	return err
}

func (pg *Postgres) GetCompanyByID(id string) (models.Company, error) {
	var company models.Company
	query := `SELECT id, name FROM companies WHERE id = $1`
	err := pg.db.QueryRow(query, id).Scan(&company.ID, &company.Name)
	return company, err
}

func (pg *Postgres) GetCompaniesByUserID(userID string) ([]models.Company, error) {
	query := `
		SELECT DISTINCT c.id, c.name
		FROM companies c
		JOIN employees e ON e.company_id = c.id
		WHERE e.user_id = $1
	`
	rows, err := pg.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []models.Company
	for rows.Next() {
		var company models.Company
		if err := rows.Scan(&company.ID, &company.Name); err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}
	return companies, nil
}
