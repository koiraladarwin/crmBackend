package postgres

import "github.com/koiraladarwin/crmbackend/models"


func (pg *Postgres) AddClientProcess(clientID, employeeID string, revenue float64, priority, status string) error {
	query := `INSERT INTO client_process 
		(client_id, assigned_employee_id, expected_revenue, priority, status) 
		VALUES ($1, $2, $3, $4, $5)`
	_, err := pg.db.Exec(query, clientID, employeeID, revenue, priority, status)
	return err
}

func (pg *Postgres) GetClientProcessByClientID(clientID string) (models.ClientProcess, error) {
	var cp models.ClientProcess
	query := `SELECT client_id, assigned_employee_id, expected_revenue, priority, status 
			  FROM client_process WHERE client_id = $1`
	err := pg.db.QueryRow(query, clientID).Scan(
		&cp.ClientID,
		&cp.AssignedEmployeeID,
		&cp.ExpectedRevenue,
		&cp.Priority,
		&cp.Status,
	)
	return cp, err
}
