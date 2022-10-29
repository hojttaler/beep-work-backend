package database

import (
	"beep-work-backend/config"
	database "beep-work-backend/database/queries"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Queries struct {
	*database.UserQueries
	*database.SessionQueries
}

func OpenConnection() (*Queries, error) {
	connectionLink := config.GetEnv("DB_CONNECTION_LINK")

	db, err := sqlx.Connect("pgx", connectionLink)
	if err != nil {
		return nil, fmt.Errorf("can't connect to database, %w", err)
	}

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("can't connect to database, %w", err)
	}

	return &Queries{
		UserQueries:    &database.UserQueries{DB: db},
		SessionQueries: &database.SessionQueries{DB: db},
	}, nil
}
