package models

import (
	"errors"
	"example.com/url-shortner/db"
	"example.com/url-shortner/utils"
)

type User struct {
	ID       int64
	Username string `json:"username" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Save() error {
	query := `INSERT INTO users(username,password) VALUES($1, $2) RETURNING id`
	_, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	err = db.DB.QueryRow(query, user.Username, hashedPassword).Scan(&user.ID)
	if err != nil {
		return err
	}
	return err
}

func (user *User) CheckUser() error {
	query := `SELECT id,password FROM users WHERE username=$1`
	row := db.DB.QueryRow(query, user.Username)

	var retrivedPassword string
	err := row.Scan(&user.ID, &retrivedPassword)
	if err != nil {
		return errors.New("User not found")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrivedPassword)
	if !passwordIsValid {
		return errors.New("InvalidPassword")
	}

	return nil
}
