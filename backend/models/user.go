package models

import (
	"backend/db"
	"backend/utils"
	"errors"
	"fmt"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPass)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = userId
	return err
}

func (user *User) ValidateCredentials() error {
	fmt.Printf("fn called")
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)
	fmt.Println("validpass: ", passwordIsValid)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}
	return nil
}
