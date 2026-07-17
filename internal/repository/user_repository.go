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

func InsertUser(db *sql.DB, user model.User) error {
	// The ? placeholders protect against SQL injection
	query := "INSERT INTO users(name, email, password, phone, address,user_type) VALUES(?, ?, ?,?,?,?)"

	// Query pakai model User
	_, err := db.Exec(query, user.Name, user.Email, user.Password, user.Phone, user.Address, user.UserType)
	if err != nil {
		return fmt.Errorf("failed to insert customer:%w", err)
	}
	// sukses insert user
	fmt.Printf("Successfully added customer: %s\n", user.Name)
	return nil
}

func IsEmailExists(db *sql.DB, email string) bool {
	var exists bool
	// Query ini akan mengembalikan true (1) jika email ada, dan false (0) jika tidak ada
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"

	err := db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false // Jika error database, asumsikan belum ada agar tidak memblokir sistem
	}

	return exists
}

// ShowUsers mengambil semua data user dari database
func ShowUsers(db *sql.DB) ([]model.User, error) {
	query := "SELECT id, name, email, phone, user_type FROM users"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to load query: %w", err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		// Ingat: Urutan Scan harus sama persis dengan urutan SELECT di query
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.UserType)
		if err != nil {
			return nil, fmt.Errorf("failed to scan data: %w", err)
		}
		users = append(users, u)
	}

	return users, nil
}

// UpdateUser memperbarui data user berdasarkan ID
func UpdateUser(db *sql.DB, id int, name, phone, address, userType string) error {
	query := "UPDATE users SET name = ?, phone = ?, address = ?, user_type = ? WHERE id = ?"

	result, err := db.Exec(query, name, phone, address, userType, id)
	if err != nil {
		return fmt.Errorf("failed to do query update: %w", err)
	}

	// Mengecek apakah ada baris data yang berubah (ID ditemukan)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check status update: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("user dengan ID tersebut tidak ditemukan")
	}

	return nil
}

// DeleteUser menghapus data user dari database berdasarkan ID
func DeleteUser(db *sql.DB, id int) error {
	query := "DELETE FROM users WHERE id = ?"

	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete query: %w", err)
	}

	// Mengecek apakah ada baris data yang berhasil dihapus
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check delete status: %w", err)
	}

	// Jika rowsAffected == 0, artinya ID tersebut tidak ada di database
	if rowsAffected == 0 {
		return errors.New("user ID not found")
	}

	return nil
}
