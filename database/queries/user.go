package database

import (
	"beep-work-backend/models"
	"github.com/jmoiron/sqlx"
)

type UserQueries struct {
	*sqlx.DB
}

func (q *UserQueries) GetUserById(id string) (models.User, error) {
	user := models.User{}

	query := `SELECT * FROM "users" WHERE id = $1`

	err := q.Get(&user, query, id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (q *UserQueries) GetUserByEmail(email string) (models.User, error) {
	user := models.User{}

	query := `SELECT * FROM "users" WHERE email = $1`

	err := q.Get(&user, query, email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (q *UserQueries) GetUserByUsername(username string) (models.User, error) {
	user := models.User{}

	query := `SELECT * FROM "users" WHERE username = $1`

	err := q.Get(&user, query, username)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (q *UserQueries) CreateUser(user *models.User) error {
	query := `INSERT INTO "users" VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := q.Exec(
		query,
		user.ID,
		user.Email,
		user.Username,
		user.Role,
		user.Status,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
