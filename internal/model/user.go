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
type Add_on struct {
	ID    int
	Name  string
	Price int
}

type Product_add_ons struct {
	ID         int
	Product_id int
	Add_on_id  int
}

type Order struct {
	ID           int
	User_id      int
	Status       string
	Total_amount int
}

type Payment struct {
	ID                int
	Order_id          int
	Amount            int
	Payment_method_id int
	status            string
	paid_at           string
}

type Payments_method struct {
	ID   int
	Name string
}
type Order_items struct {
	ID         int
	Order_id   int
	Product_id int
	Add_on_id  int
	Quantity   int
	Unit_price int
	Subtotal   int
	Note       string
}

type UserHistory struct {
	OrderID     int
	ProductName string
	Quantity    int
	Subtotal    int
}

type TotalRevenue struct {
	Revenue int
}

type ProductSales struct {
	Name  string
	sales int
}