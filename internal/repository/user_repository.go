package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"pairproject/internal/model"
)

func UserLogin(db *sql.DB, email, password string) (int, string, error) {
	var id int
	var UserType string

	query := "SELECT id, user_type FROM users WHERE email = ? AND password = ?"
	err := db.QueryRow(query, email, password).Scan(&id, &UserType)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, "", errors.New("email atau password salah!")
		}
		return 0, "", fmt.Errorf("database error: %w", err)
	}
	return id, UserType, nil
}

func CustomerRegister(db *sql.DB, user model.User) error {
	// The ? placeholders protect against SQL injection
	query := "INSERT INTO users(name, email, password, phone, address,user_type) VALUES(?, ?, ?,?,?,'customer')"

	// Query pakai model User
	_, err := db.Exec(query, user.Name, user.Email, user.Password, user.Phone, user.Address)
	if err != nil {
		return fmt.Errorf("failed to insert customer:%w", err)
	}
	// simulasi sukses
	fmt.Printf("Successfully added customer: %s\n", user.Name)
	return nil
}
