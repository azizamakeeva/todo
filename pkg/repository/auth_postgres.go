package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo_app "todo-app"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo_app.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name,username,password_hash) values ($1,$2,$3) RETURNING id", usersListsTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
