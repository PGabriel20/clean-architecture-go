package repository

import (
	"database/sql"
	"fmt"

	"github.com/PGabriel20/expenses-go/internal/entity"
)

type UserRepositoryPostgres struct {
	Db *sql.DB
}

func NewUserRepositoryPostgres(db *sql.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{Db: db}
}

func(r *UserRepositoryPostgres) GetUser(id string) (*entity.User, error) {
	query := "SELECT id, username, email FROM users WHERE id = $1"
	row := r.Db.QueryRow(query, id)

	var user entity.User
	err := row.Scan(&user.ID, &user.Username, &user.Email)

	if err != nil {
		return &entity.User{}, err
	}

	return &user, nil
}

func (r *UserRepositoryPostgres) RegisterUser(user entity.User) error {
	
	//Running prepared statements, for concurrency
	stmt, err := r.Db.Prepare("INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4)")

	if err != nil {
		fmt.Println("err preparing query")
		return  err
	}

	_, err = stmt.Exec(user.ID, user.Username, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}