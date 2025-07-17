package models

import (
	"errors"
	"strings"

	"github.com/Niranjini-Kathiravan/go-rest-api-v2/db"
	"github.com/Niranjini-Kathiravan/go-rest-api-v2/utils"
)

type User struct {
	ID       int64
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	u.Email = strings.ToLower(u.Email) // Normalize email to lowercase

	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = userId
	return nil
}

func (u *User) ValidateCredentials() error {
	u.Email = strings.ToLower(u.Email) // Normalize email to lowercase

	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var storedHash string
	err := row.Scan(&u.ID, &storedHash)
	if err != nil {
		// return a generic error if user not found or other DB error
		return errors.New("credentials invalid")
	}

	if !utils.CheckPasswordHash(u.Password, storedHash) {
		return errors.New("credentials invalid")
	}

	return nil
}
