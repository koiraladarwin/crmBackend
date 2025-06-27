package postgres

import "github.com/koiraladarwin/crmbackend/storage/models"

func (pg *Postgres) AddEmployee(id, userID, companyID, role string) error {
	query := `INSERT INTO employees (id, user_id, company_id, role) VALUES ($1, $2, $3, $4)`
	_, err := pg.db.Exec(query, id, userID, companyID, role)
	return err
}

func (pg *Postgres) GetEmployeeByID(id string) (models.Employee, error) {
	var emp models.Employee
	query := `SELECT id, user_id, company_id, role FROM employees WHERE id = $1`
	err := pg.db.QueryRow(query, id).Scan(&emp.ID, &emp.UserID, &emp.CompanyID, &emp.Role)
	return emp, err
}
