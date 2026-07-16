package handler

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"pairproject/internal/model"
	"pairproject/internal/repository"
	"strings"
)

func Run(db *sql.DB) {
	scanner := bufio.NewScanner(os.Stdin)
	// tampilan awal
	for {
		fmt.Println("\n============================")
		fmt.Println("WELCOME TO BEVERAGE CAFE (TEAM 5)")
		fmt.Println("============================")
		fmt.Println("Choose a number:")
		fmt.Println("1. Login")
		fmt.Println("2. Don't have an account yet? Register here")
		fmt.Println("0. Exit Application")
		fmt.Print("Choose (0-2): ")

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		switch input {
		case "1":
			handleLogin(scanner, db)
		case "2":
			customerRegister(scanner, db)
		case "0":
			fmt.Println("Exiting Application.. Thank You!")
			return
		default:
			fmt.Println("Wrong Input!")
		}
	}

}

func handleLogin(scanner *bufio.Scanner, db *sql.DB) {
	fmt.Println("\n=== LOGIN ===")
	fmt.Print("Email: ")
	scanner.Scan()
	email := strings.TrimSpace(scanner.Text())

	fmt.Print("Password: ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())

	// panggil userLogin dari repo
	userID, UserType, err := repository.UserLogin(db, email, password)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}

	// ngarahin ke dashboard berdasarkan role
	if UserType == "admin" {
		fmt.Println("Login Success! Welcome Back Admin!")
		adminMenu(scanner, db)
	} else if UserType == "customer" {
		fmt.Printf("Login Success! Welcome Back Customer (ID: %d)!\n", userID)
		customerMenu(scanner, db, userID)
	} else {
		fmt.Println("role unknown")
	}
}

// register customer
func customerRegister(scanner *bufio.Scanner, db *sql.DB) {
	fmt.Println("--- Register Customer ---")
	fmt.Print("Nama: ")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())
	fmt.Print("Email: ")
	scanner.Scan()
	email := strings.TrimSpace(scanner.Text())
	fmt.Print("Password: ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())
	fmt.Print("Phone: ")
	scanner.Scan()
	phone := strings.TrimSpace(scanner.Text())
	fmt.Print("Address: ")
	scanner.Scan()
	address := strings.TrimSpace(scanner.Text())
	// narik model, user.go
	newUser := model.User{
		Name:     name,
		Email:    email,
		Password: password,
		Phone:    phone,
		Address:  address,
	}

	err := repository.CustomerRegister(db, newUser)
	if err != nil {
		fmt.Println("Register Failed:", err)
		return
	}
}

// menu dashboard admin
func adminMenu(scanner *bufio.Scanner, db *sql.DB) {
	for {
		fmt.Println("\n=== ADMIN DASHBOARD ===")
		fmt.Println("--- User Management ---")
		fmt.Println("1. Create new User account")
		fmt.Println("2. List all User account")
		fmt.Println("3. Update user account data")
		fmt.Println("4. Delete user account")
		fmt.Println("--- Product Management ---")
		fmt.Println("5. Create new product")
		fmt.Println("6. List all product")
		fmt.Println("7. Update product data")
		fmt.Println("8. Delete Product")
		fmt.Println("--- Reports & Orders ---")
		fmt.Println("9. List all order")
		fmt.Println("10. Report revenue for all products")
		fmt.Println("11. List product with the highest sales")
		fmt.Println("12. Report user with most frequent order")
		fmt.Println("0. LogOut")
		fmt.Print("Choose a Number (0-7): ")

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		// hasilnya jika dipilih angka (belom gw bikin logicnya)
		switch input {
		case "1":
			fmt.Println("\nCreate User")
		case "2":
			fmt.Println("\nShows User")

		case "3":
			fmt.Println("\nReport Products")
		case "4":
			fmt.Println("\nShow Products")
		case "5":
			fmt.Println("\nShow Orders")
		case "6":
			fmt.Println("\nReport Revenue")
		case "7":
			fmt.Println("\nReport Product Terlaris")
		case "0":
			fmt.Println("\nLogging out...")
			return
		default:
			fmt.Println("\nWrong Input!")
		}
	}
}

// menu dashboard customer
func customerMenu(scanner *bufio.Scanner, db *sql.DB, customerID int) {
	for {
		fmt.Println("\n--- CUSTOMER DASHBOARD ---")
		fmt.Println("1. Create new order")
		fmt.Println("2. Display Menu")
		fmt.Println("3. Check Order History")
		fmt.Println("0. LogOut")
		fmt.Print("Choose Number (0-3): ")

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		// logic nya belom gw bikin juga :D
		switch input {
		case "1":
			fmt.Println("/nMake Orders")
		case "2":
			fmt.Println("/nCheck Menu")
		case "3":
			fmt.Println("/nCheck Order Status")
		case "0":
			fmt.Println("/nLoggingOut...")
			return
		default:
			fmt.Println("Wrong Input!")
		}
	}
}
