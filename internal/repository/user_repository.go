package repository

import (
	"GamingStore/internal/model"
	"database/sql"
	"fmt"
)


func CreateCustomer(db *sql.DB, name string, email string, phone string) error {
		// The ? placeholders protect against SQL injection
		query := "INSERT INTO customers(name, email, phone) VALUES(?, ?, ?)"

		// db.Exec is used for queries that don't return rows (INSERT, UPDATE, DELETE)
		_, err := db.Exec(query, name, email, phone)
		if err != nil {
			return fmt.Errorf("failed to insert customer:%w",err)
		}
		fmt.Printf("Successfully added customer: %s\n", name)
		return nil
	}
func ListCustomer(db *sql.DB) ([]model.Customer,error){
	query := "SELECT id, name, email, phone, created_at from customers"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()
	var customers []model.Customer
	
	for rows.Next() {
		var c model.Customer

		
		err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.Created_at)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		customers = append(customers, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return customers, nil

}
