package postgres

import "github.com/koiraladarwin/crmbackend/storage/models"

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
