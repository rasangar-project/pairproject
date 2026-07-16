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
		fmt.Println("BEVERAGE (TEAM 5)")
		fmt.Println("============================")
		fmt.Println("Pilih untuk masuk sebagai:")
		fmt.Println("1. ADMIN")
		fmt.Println("2. CUSTOMER")
		fmt.Println("0. Exit Application")
		fmt.Print("Pilih (0-2): ")

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		switch input {
		case "1":
			adminAuthMenu(scanner, db)
		case "2":
			customerAuthMenu(scanner, db)
		case "0":
			fmt.Println("Menutup Aplikasi.. Terima kasih!")
			return
		default:
			fmt.Println("Input tidak valid! silahkan pilih 0, 1, atau 2")
		}
	}

}

// Logic CLI Dashboard Admin
func adminAuthMenu(scanner *bufio.Scanner, db *sql.DB) {
	fmt.Println("=== LOGIN ADMIN ===")
	fmt.Println("LOGIN")
	fmt.Print("Email: ")
	scanner.Scan()
	email := strings.TrimSpace(scanner.Text())

	fmt.Print("Password: ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())

	// Panggil Repo
	err := repository.AdminLogin(db, email, password)
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	fmt.Println("Login Berhasil ! Selamat Datang Admin!")
	adminMenu(scanner, db) //Masuk ke Dashboard Admin

}

// Dashboard Customer
func customerAuthMenu(scanner *bufio.Scanner, db *sql.DB) {
	for {
		fmt.Println("\n --- Customer Module ---")
		fmt.Println("1. Login Customer")
		fmt.Println("2. Register Akun Baru")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih (0-2): ")

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		switch input {
		case "1":
			customerLogin(scanner, db)
		case "2":
			customerRegister(scanner, db)
		case "0":
			return
		default:
			fmt.Println("Input tidak valid! masukkan angka, pilih angka 0 sampai 2")
		}
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
		fmt.Println("Gagal register:", err)
		return
	}
}

// login customer
func customerLogin(scanner *bufio.Scanner, db *sql.DB) {
	fmt.Println("\n--- Login Customer ---")
	fmt.Print("Email: ")
	scanner.Scan()
	email := strings.TrimSpace(scanner.Text())
	fmt.Print("Password: ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())

	// panggil repo
	customerID, err := repository.CustomerLogin(db, email, password)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Printf("Login berhasil! Selamat Datang (ID: %d)\n", customerID)
	customerMenu(scanner, db, customerID)
}

// menu dashboard admin
func adminMenu(scanner *bufio.Scanner, db *sql.DB) {
	for {
		fmt.Println("\n=== ADMIN DASHBOARD ===")
		fmt.Println("1. Create Users")
		fmt.Println("2. Shows Users")
		fmt.Println("3. Report Products")
		fmt.Println("4. Show Products")
		fmt.Println("5. Show Orders")
		fmt.Println("6. Report Revenue")
		fmt.Println("7. Report Product Terlaris")
		fmt.Println("0. LogOut")
		fmt.Print("Pilih menu Admin (0-7): ")

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
			fmt.Println("\nLogout")
			return
		default:
			fmt.Println("\nInput Salah!")
		}
	}
}

// menu dashboard customer
func customerMenu(scanner *bufio.Scanner, db *sql.DB, customerID int) {
	for {
		fmt.Println("\n--- CUSTOMER DASHBOARD ---")
		fmt.Println("1. Make Orders")
		fmt.Println("2. Check Menu")
		fmt.Println("3. Check Order Status")
		fmt.Println("0. Back")
		fmt.Print("Pilih menu Customer (0-3): ")

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
			fmt.Println("/nBack")
			return
		default:
			fmt.Println("Input Salah!")
		}
	}
}
