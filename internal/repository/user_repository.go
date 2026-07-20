package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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

func DeleteProduct(db *sql.DB, id int) error {
	query := "DELETE FROM products WHERE id = ?"

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

// get most frequent order case 12
func GetMostFrequentUsers(db *sql.DB) ([]model.UserOrderReport, error) {
	query := `
		SELECT u.id, u.name, u.email, COUNT(o.id) as total_orders
		FROM users u
		JOIN orders o ON u.id = o.user_id
		GROUP BY u.id, u.name, u.email
		ORDER BY total_orders DESC
		LIMIT 5
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query report: %w", err)
	}
	defer rows.Close()

	var reports []model.UserOrderReport
	for rows.Next() {
		var r model.UserOrderReport
		err := rows.Scan(&r.UserID, &r.Name, &r.Email, &r.TotalOrders)
		if err != nil {
			return nil, fmt.Errorf("failed to scan report: %w", err)
		}
		reports = append(reports, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func FetchingProducts(db *sql.DB) ([]model.Product, error) {
	query := "SELECT id,category_id, name, stock, price from products"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	defer rows.Close()
	var products []model.Product

	for rows.Next() {
		var p model.Product

		err := rows.Scan(&p.ID, &p.CategoryID, &p.Name, &p.Stock, &p.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		products = append(products, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func CheckValidProductId(db *sql.DB, id int) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM products WHERE id = ?)`

	err := db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return exists
}

func FetchProductAddOn(db *sql.DB, product_id int) ([]model.Add_on, error) {
	query := "select a.id as add_ons_id, a.name as add_ons_name, a.price from product_add_ons as pa Join add_ons as a ON a.id = pa.add_on_id Join products as p on p.id = pa.product_id where pa.product_id = ?; "

	rows, err := db.Query(query, product_id)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	defer rows.Close()
	var add_ons []model.Add_on

	for rows.Next() {
		var a model.Add_on

		err := rows.Scan(&a.ID, &a.Name, &a.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		add_ons = append(add_ons, a)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return add_ons, nil
}

func ChecksValidAddOn(db *sql.DB, product_id int, add_on_id int) bool {
	var connected bool
	query := `SELECT EXISTS(SELECT 1 FROM product_add_ons WHERE product_id = ? AND add_on_id = ?)`

	err := db.QueryRow(query, product_id, add_on_id).Scan(&connected)
	if err != nil {
		log.Fatal(err)
		return connected
	}
	return connected
}

func GetProductPrice(db *sql.DB, product_id int) (int, error) {
	query := "SELECT price FROM products WHERE id = ?"
	var price int
	err := db.QueryRow(query, product_id).Scan(&price)
	if err != nil {
		return 0, err
	}
	return price, nil
}
func GetAddOnPrice(db *sql.DB, add_on_id int) (int, error) {
	query := "SELECT price FROM add_ons WHERE id = ?"
	var price int
	err := db.QueryRow(query, add_on_id).Scan(&price)
	if err != nil {
		return 0, err
	}
	return price, nil
}

func FetchPaymentMethod(db *sql.DB) ([]model.Payments_method, error) {
	query := "SELECT * FROM payments_method"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	defer rows.Close()
	var payments_method []model.Payments_method

	for rows.Next() {
		var p model.Payments_method

		err := rows.Scan(&p.ID, &p.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		payments_method = append(payments_method, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return payments_method, nil
}

func CheckPayMethodValid(db *sql.DB, pay_id int) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM payments_method WHERE id = ?)`

	err := db.QueryRow(query, pay_id).Scan(&exists)
	if err != nil {
		log.Fatal(err)
		return exists
	}
	return exists
}

func GetProductName(db *sql.DB, product_id int) (string, error) {
	query := "SELECT name FROM products WHERE id = ?"
	var product_name string
	err := db.QueryRow(query, product_id).Scan(&product_name)
	if err != nil {
		return "", err
	}
	return product_name, nil
}

func GetAddOnName(db *sql.DB, add_onId int) (string, error) {
	query := "SELECT name FROM add_ons WHERE id = ?"
	var add_onName string
	err := db.QueryRow(query, add_onId).Scan(&add_onName)
	if err != nil {
		return "", err
	}
	return add_onName, nil
}

func CreatingOrdersDB(db *sql.DB, customerID int, totalAmount int, paymentMethod int, order_items []model.Order_items) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer tx.Rollback()

	orderQuery := `INSERT INTO orders (user_id, status, total_amount) VALUES (?, ?, ?)`
	orderResult, err := tx.Exec(orderQuery, customerID, "completed", totalAmount)
	if err != nil {
		return fmt.Errorf("failed to insert order: %w", err)
	}

	orderID64, err := orderResult.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %w", err)
	}
	orderID := int(orderID64)

	itemQuery := `INSERT INTO order_items (order_id, product_id, add_on_id, quantity, unit_price, subtotal, note) 
	              VALUES (?, ?, ?, ?, ?, ?, ?)`
	updateStockQuery := `UPDATE products Set stock = stock - ? WHERE id = ? AND stock >= ?`

	for _, item := range order_items {
		if item.Add_on_id == 0 {
			_, err := tx.Exec(itemQuery, orderID, item.Product_id, nil, item.Quantity, item.Unit_price, item.Subtotal, item.Note)
			if err != nil {
				return fmt.Errorf("failed to insert order item for product %d: %w", item.Product_id, err)
			}
		} else {
			_, err := tx.Exec(itemQuery, orderID, item.Product_id, item.Add_on_id, item.Quantity, item.Unit_price, item.Subtotal, item.Note)
			if err != nil {
				return fmt.Errorf("failed to insert order item for product %d: %w", item.Product_id, err)
			}
		}
		_, err := tx.Exec(updateStockQuery, item.Quantity, item.Product_id, item.Quantity)
		if err != nil {
			return fmt.Errorf("failed to execute stock update: %w", err)
		}
	}

	paymentQuery := `INSERT INTO payments (order_id, amount, payment_method_id, status, paid_at) VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP)`
	_, err = tx.Exec(paymentQuery, orderID, totalAmount, paymentMethod, "paid")
	if err != nil {
		return fmt.Errorf("failed to create payment record: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("Order created successfully!")
	return nil
}

func GetUserHistory(db *sql.DB, customerID int) ([]model.UserHistory, error) {
	query := "select o.id as order_id, p.name as product_name, oi.quantity as quantity, oi.subtotal as subtotal from order_items as oi join orders o ON o.id = oi.order_id join products p on p.id = oi.product_id left join add_ons a ON oi.add_on_id = a.id where o.user_id = ?;"

	rows, err := db.Query(query, customerID)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	defer rows.Close()
	var userHistory []model.UserHistory

	for rows.Next() {
		var u model.UserHistory

		err := rows.Scan(&u.OrderID, &u.ProductName, &u.Quantity, &u.Subtotal)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		userHistory = append(userHistory, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return userHistory, nil
}

func GenerateTotalRevenue(db *sql.DB) (model.TotalRevenue, error) {
	query := "SELECT SUM(o.total_amount) AS total_revenue FROM orders o JOIN payments p ON o.id = p.order_id WHERE p.status = 'Paid'"
	var revenue model.TotalRevenue
	err := db.QueryRow(query).Scan(&revenue.Revenue) //minor bug tampilan {}
	if err != nil {
		return revenue, err
	}
	return revenue, nil
}

// ListAllOrders mengambil semua daftar order digabung dengan nama user (case 9 adminMenu)
func ListAllOrders(db *sql.DB) ([]model.OrderDetail, error) {
	query := `
		SELECT o.id, u.name, o.status, o.total_amount 
		FROM orders o 
		JOIN users u ON o.user_id = u.id 
		ORDER BY o.id DESC
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch orders: %w", err)
	}
	defer rows.Close()

	var orders []model.OrderDetail
	for rows.Next() {
		var o model.OrderDetail
		if err := rows.Scan(&o.ID, &o.CustomerName, &o.Status, &o.TotalAmount); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

// GetHighestSalesProducts menampilkan top 5 produk paling laris (case 11 adminMenu)
func GetHighestSalesProducts(db *sql.DB) ([]model.ProductSales, error) {
	query := `
		SELECT p.name, SUM(oi.quantity) as sales
		FROM order_items oi
		JOIN products p ON oi.product_id = p.id
		GROUP BY p.id, p.name
		ORDER BY sales DESC
		LIMIT 5
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch highest sales: %w", err)
	}
	defer rows.Close()

	var sales []model.ProductSales
	for rows.Next() {
		var s model.ProductSales
		if err := rows.Scan(&s.Name, &s.Sales); err != nil {
			return nil, err
		}
		sales = append(sales, s)
	}
	return sales, nil
}
