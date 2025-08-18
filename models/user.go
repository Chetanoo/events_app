package models

import (
	"errors"
	"events_app/db"
	"events_app/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users (email, password)
		VALUES (?, ?)
	`
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := db.DB.Exec(query, u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := `
		SELECT id, password FROM users WHERE email = ?
	`
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return err
	}

	passwordIsValid := utils.ComparePasswords(retrievedPassword, u.Password)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}
	return nil
}
