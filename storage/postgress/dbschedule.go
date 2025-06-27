package postgres

import "github.com/koiraladarwin/crmbackend/models"


func (pg *Postgres) AddSchedule(id, clientID, employeeID, schedule string) error {
	query := `INSERT INTO schedules (id, process_client_id, process_assigned_employee_id, schedule) 
			  VALUES ($1, $2, $3, $4)`
	_, err := pg.db.Exec(query, id, clientID, employeeID, schedule)
	return err
}


func (pg *Postgres) GetSchedulesByProcess(clientID, employeeID string) ([]models.Schedule, error) {
	query := `
		SELECT id, process_client_id, process_assigned_employee_id, schedule
		FROM schedules
		WHERE process_client_id = $1 AND process_assigned_employee_id = $2
	`
	rows, err := pg.db.Query(query, clientID, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []models.Schedule
	for rows.Next() {
		var sch models.Schedule
		if err := rows.Scan(&sch.ID, &sch.ProcessClientID, &sch.ProcessAssignedEmpID, &sch.Schedule); err != nil {
			return nil, err
		}
		schedules = append(schedules, sch)
	}
	return schedules, nil
}

func (pg *Postgres) GetSchedulesByCompanyID(companyID string) ([]models.Schedule, error) {
	query := `
		SELECT s.id, s.process_client_id, s.process_assigned_employee_id, s.schedule
		FROM schedules s
		JOIN employees e ON s.process_assigned_employee_id = e.id
		WHERE e.company_id = $1
	`
	rows, err := pg.db.Query(query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []models.Schedule
	for rows.Next() {
		var sch models.Schedule
		if err := rows.Scan(&sch.ID, &sch.ProcessClientID, &sch.ProcessAssignedEmployeeID, &sch.Schedule); err != nil {
			return nil, err
		}
		schedules = append(schedules, sch)
	}
	return schedules, nil
}
