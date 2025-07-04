package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostGres(dbName string, connStr string) (*Postgres, error) {
	db, err := sql.Open(dbName, connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	users := `CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    password TEXT NOT NULL,
		name TEXT NOT NULL,
		gmail TEXT NOT NULL UNIQUE,
		phone TEXT
	);`
	if _, err := db.Exec(users); err != nil {
		return nil, err
	}

	companies := `CREATE TABLE IF NOT EXISTS companies (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name TEXT NOT NULL
	);`
	if _, err := db.Exec(companies); err != nil {
		return nil, err
	}

	employees := `CREATE TABLE IF NOT EXISTS employees (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    company_id UUID NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    role TEXT NOT NULL,
    UNIQUE (user_id, company_id)
   );`
	if _, err := db.Exec(employees); err != nil {
		return nil, err
	}

	clients := `
    CREATE TABLE IF NOT EXISTS clients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    company_id UUID NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    gmail TEXT NOT NULL,
    phone TEXT,
    UNIQUE (company_id, gmail) 
  );
`
	if _, err := db.Exec(clients); err != nil {
		return nil, err
	}

	clientProcess := `CREATE TABLE IF NOT EXISTS client_process (
  	client_id UUID NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
  	assigned_employee_id UUID NOT NULL REFERENCES employees(id) ON DELETE SET NULL,
  	expected_revenue NUMERIC(12,2) DEFAULT 0.00,
  	priority TEXT CHECK (priority IN ('Low', 'Medium', 'High')) DEFAULT 'Medium',
  	status TEXT DEFAULT 'Pending',
  	PRIMARY KEY (client_id, assigned_employee_id)
  );`

	if _, err := db.Exec(clientProcess); err != nil {
		return nil, err
	}

	schedule := `CREATE TABLE IF NOT EXISTS schedules (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	process_client_id UUID NOT NULL,
	process_assigned_employee_id UUID NOT NULL,
	schedule TEXT NOT NULL,
	FOREIGN KEY (process_client_id, process_assigned_employee_id)
		REFERENCES client_process(client_id, assigned_employee_id)
		ON DELETE CASCADE
  );`
	if _, err := db.Exec(schedule); err != nil {
		return nil, err
	}

	return &Postgres{
		db: db,
	}, nil
}
