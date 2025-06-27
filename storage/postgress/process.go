package postgres

import "github.com/koiraladarwin/crmbackend/storage/models"

func (pg *Postgres) AddClientProcess(clientID, employeeID string, revenue float64, priority string) error {
	query := `INSERT INTO client_process (client_id, assigned_employee_id, expected_revenue, priority) 
			  VALUES ($1, $2, $3, $4)`
	_, err := pg.db.Exec(query, clientID, employeeID, revenue, priority)
	return err
}

func (pg *Postgres) GetClientProcessByClientID(clientID string) (models.ClientProcess, error) {
	var cp models.ClientProcess
	query := `SELECT client_id, assigned_employee_id, expected_revenue, priority 
			  FROM client_process WHERE client_id = $1`
	err := pg.db.QueryRow(query, clientID).Scan(&cp.ClientID, &cp.AssignedEmployeeID, &cp.ExpectedRevenue, &cp.Priority)
	return cp, err
}
