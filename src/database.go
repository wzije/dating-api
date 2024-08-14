package src

import (
	"context"
	"database/sql"
)

type Database interface {
	GetUserByEmail(email string) (*User, error)
	StoreUser(user User) (*int64, error)
}

type database struct {
	db *sql.DB
}

func (r database) GetUserByEmail(email string) (*User, error) {
	var user User
	query := "SELECT * FROM users WHERE email = ?"

	err := r.db.QueryRow(query, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Gender, &user.Address)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r database) StoreUser(user User) (*int64, error) {

	query := "INSERT INTO `users` (`name`, `password`, `email`, `gender`, `address`) VALUES (?, ?, ?, ?, ?)"

	result, err := r.db.ExecContext(
		context.Background(),
		query, &user.Name, &user.Password, &user.Email, &user.Gender, user.Address)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func NewDatabase(db *sql.DB) Database {
	return &database{
		db: db,
	}
}
