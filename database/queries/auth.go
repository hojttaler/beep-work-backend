package database

import (
	"beep-work-backend/models"
	"github.com/jmoiron/sqlx"
)

type SessionQueries struct {
	*sqlx.DB
}

func (q *SessionQueries) GetSessionById(id string) (models.Session, error) {
	session := models.Session{}

	query := `SELECT * FROM "sessions" WHERE id = $1`

	err := q.Get(&session, query, id)
	if err != nil {
		return session, err
	}

	return session, nil
}

func (q *SessionQueries) GetSessionWithUser(sessionId string, userId string) (models.SessionWithUser, error) {
	session := models.SessionWithUser{}

	query := `
		SELECT * FROM "sessions" LEFT JOIN "users" 
		    ON "sessions"."user_id" = "users"."id" 
		WHERE "sessions"."id" = $1 AND "sessions"."user_id" = $2;
	`

	err := q.Get(&session, query, sessionId, userId)
	if err != nil {
		return session, err
	}

	session.User.Password = ""

	return session, nil
}

func (q *SessionQueries) CreateSession(session *models.Session) error {
	query := `INSERT INTO "sessions" VALUES ($1, $2, $3, $4)`

	_, err := q.Exec(
		query,
		session.ID,
		session.UserID,
		session.CreatedAt,
		session.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
