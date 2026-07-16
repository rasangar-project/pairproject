package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"pairproject/internal/model"
)

func AdminLogin(db *sql.DB, email, password string) error {
	// DB querry (belom nyambung)

	// tes dulu
	if email == "admin" && password == "admin" {
		return nil
	}
	return errors.New("email atau password admin salah")
}

func CustomerRegister(db *sql.DB, user model.User) error {
	// The ? placeholders protect against SQL injection
	query := "INSERT INTO users(name, email, password, phone, address,user_type) VALUES(?, ?, ?,?,?,'customer')"

	// Query pakai model User
	_, err := db.Exec(query, user.Name, user.Password, user.Password, user.Phone, user.Address)
	if err != nil {
		return fmt.Errorf("failed to insert customer:%w", err)
	}
	// simulasi sukses
	fmt.Printf("Successfully added customer: %s\n", user.Name)
	return nil
}

func CustomerLogin(db *sql.DB, email, password string) (int, error) {
	// select Querry konek DB (belum bikin)

	if email == "customer" && password == "customer" {
		return 1, nil
	}
	return 0, errors.New("email / pass customer salah")
}
