package model

import (
	"time"

	"github.com/shopspring/decimal"
)


type Customer struct {
	ID int
	Name string
	Email string
	Phone string
	Created_at time.Time
}

type Product struct {
	ID int 
	Title string
	Platform string
	Genre string
	Price decimal.Decimal
	Stock int
	Release_year int
	Created_at time.Time
	Updated_at time.Time
}

type Order struct {
	ID int
	Customer_id int
	Total_amount decimal.Decimal
	Status string
	Created_at time.Time
	Updated_at time.Time
}
type Payment struct {
	ID int
	Order_id int
	Amount decimal.Decimal
	Method string
	Status string
	Paid_at time.Time
	Created_at time.Time
}

type OrderItem struct {
	ID int
	Order_id int
	Product_id int
	quantity int
	unit_price decimal.Decimal
}