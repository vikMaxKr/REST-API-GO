package models

import (
	"rest-api-go/db"
	"rest-api-go/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hasPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hasPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u User) ValidatePassword() (bool, error) {
	query := "SELECT password FROM users WHERE email = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	var hashedPassword string
	err = stmt.QueryRow(u.Email).Scan(&hashedPassword)

	if err != nil {
		return false, err
	}

	err = utils.CheckPasswordHash(u.Password, hashedPassword)

	if err != nil {
		return false, err
	}

	return true, nil
}
