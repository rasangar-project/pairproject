package model

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Phone    string
	Address  string
	UserType string
}

type Category struct {
	ID   int
	Name string
}

type Product struct {
	ID           int
	CategoryID   int
	CategoryName string
	Name         string
	Stock        int
	Price        int64
}

type UserOrderReport struct {
	UserID      int
	Name        string
	Email       string
	TotalOrders int
}
