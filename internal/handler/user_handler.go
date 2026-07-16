package handler

import (
	"bufio"
	"database/sql"
	"fmt"
	"go/scanner"
	"os"
	"strings"
)

func Run(db *sql.DB) {
	scanner := bufio.NewScanner(os.Stdin)

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

	// bagian Query (sek nanti ya, wkwk)

	// simulasi login (buat ngetes)
	if email == "admin" && password == "admin" {
		fmt.Println("Login berhasil! Selamat Datang!")
		adminAuthMenu(scanner, db)
	} else {
		fmt.Println("Email atau Password salah! bukan Admin!")
	}

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

func customerRegister() {
	fmt.Println("--- Register Customer ---")
	fmt.Print("Nama: ")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())
	fmt.Print("Password: ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())
	fmt.Print("Phone: ")
	scanner.Scan()
	phone := strings.TrimSpace(scanner.Text())
	fmt.Print("Address: ")
	scanner.Scan()
	address := strings.TrimSpace(scanner.Text())

	// konek DB
}

func customerLogin() {
	fmt.Println("\n--- Login Customer ---")
	fmt.Print("Email: ")
	scanner.Scan()
	email := strings.TrimSpace(scanner.Text())
	fmt.Print("Password: ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())

	// konek DB

	// simulasi user
	customerID := 1
	customerLogin(scanner, db, customerID)
}

func adminMenu(scanner bufio.Scanner, db *sql.DB) {
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

	}
}

func customerMenu(scanner bufio.Scanner, db *sql.DB) {
	for {
		fmt.Println("\n--- CUSTOMER DASHBOARD ---")
		fmt.Println("1. Make Orders")
		fmt.Println("2. Check Menus")
		fmt.Println("3. Check Order Status")
		fmt.Println("0. Back")
		fmt.Print("Pilih menu Customer (0-3): ")

	}
}
