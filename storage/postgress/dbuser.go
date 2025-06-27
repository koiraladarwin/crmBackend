package postgres

import "github.com/koiraladarwin/crmbackend/models"


func (pg *Postgres) AddUser(id, name, gmail, phone string) error {
	query := `INSERT INTO users (id, name, gmail, phone) VALUES ($1, $2, $3, $4)`
	_, err := pg.db.Exec(query, id, name, gmail, phone)
	return err
}

func (pg *Postgres) GetUserByID(id string) (models.User, error) {
	var user models.User
	query := `SELECT id, name, gmail, phone FROM users WHERE id = $1`
	err := pg.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Gmail, &user.Phone)
	return user, err
}

func (pg *Postgres) GetAllUsers() ([]models.User, error) {
	query := `SELECT id, name, gmail, phone FROM users`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Gmail, &user.Phone); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
