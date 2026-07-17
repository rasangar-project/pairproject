package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"pairproject/internal/model"
)

// ============================================================================
// USER MANAGEMENT
// ============================================================================
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

// ============================================================================
// PRODUCT MANAGEMENT
// ============================================================================
// GetCategories mengambil daftar kategori untuk ditampilkan ke Admin
func GetCategories(db *sql.DB) ([]model.Category, error) {
	rows, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

// InsertProduct mengeksekusi query pembuatan produk baru
func InsertProduct(db *sql.DB, p model.Product) error {
	query := "INSERT INTO products (category_id, name, stock, price) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, p.CategoryID, p.Name, p.Stock, p.Price)
	return err
}

// ListProducts mengambil semua data produk beserta nama kategorinya
func ListProducts(db *sql.DB) ([]model.Product, error) {
	query := `
		SELECT p.id, c.name AS category_name, p.name, p.price, p.stock 
		FROM products p 
		JOIN categories c ON p.category_id = c.id
		ORDER BY p.id ASC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("gagal mengeksekusi query list produk: %w", err)
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		// Urutan Scan harus sama persis dengan urutan kolom di query SELECT
		err := rows.Scan(&p.ID, &p.CategoryName, &p.Name, &p.Price, &p.Stock)
		if err != nil {
			return nil, fmt.Errorf("gagal membaca baris data produk: %w", err)
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// UpdateProduct memperbarui data produk di database berdasarkan ID
func UpdateProduct(db *sql.DB, id, categoryID int, name string, price int64, stock int) error {
	query := "UPDATE products SET category_id = ?, name = ?, price = ?, stock = ? WHERE id = ?"

	result, err := db.Exec(query, categoryID, name, price, stock, id)
	if err != nil {
		return fmt.Errorf("failed to execute the product update query: %w", err)
	}

	// Mengecek apakah ada produk yang ter-update (ID valid)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check status update: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("The product with that ID was not found.")
	}

	return nil
}
