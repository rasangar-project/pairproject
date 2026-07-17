package handler

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"pairproject/internal/model"
	"pairproject/internal/repository"
	"pairproject/usecase"
	"strconv"
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
	// syarat khusus untuk pengisian email
	var email string
	for {
		fmt.Print("Email: ")
		scanner.Scan()
		email = strings.TrimSpace(scanner.Text())
		// syarat 1 butuh @
		if !strings.Contains(email, "@") {
			fmt.Println("Error: format isnt valid, email must contains '@'")
			continue //ulangi loop
		}
		// syarat 2 cek db apakah email sudah dipakai
		if repository.IsEmailExists(db, email) {
			fmt.Println("Error: Email has been registered! please use another email.")
			continue //loop
		}
		break //lolos 2 syarat, lanjut ke bawah (password)
	}

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
			fmt.Println("\nCreate a New User Account")
			fmt.Print("Nama: ")
			scanner.Scan()
			name := strings.TrimSpace(scanner.Text())

			var email string
			for {
				fmt.Print("Email: ")
				scanner.Scan()
				email = strings.TrimSpace(scanner.Text())
				// syarat 1 butuh @
				if !strings.Contains(email, "@") {
					fmt.Println("Error: format isnt valid, email must contains '@'")
					continue //ulangi loop
				}
				// syarat 2 cek db apakah email sudah dipakai
				if repository.IsEmailExists(db, email) {
					fmt.Println("Error: Email has been registered! please use another email.")
					continue //loop
				}
				break //lolos 2 syarat, lanjut ke bawah (password)
			}

			fmt.Print("Password: ")
			scanner.Scan()
			password := strings.TrimSpace(scanner.Text())
			fmt.Print("Phone: ")
			scanner.Scan()
			phone := strings.TrimSpace(scanner.Text())
			fmt.Print("Address: ")
			scanner.Scan()
			address := strings.TrimSpace(scanner.Text())
			// bikin validasi admin customer
			var userType string
			for {
				fmt.Print("Usertype (admin / customer): ")
				scanner.Scan()
				// convert ke lowercase
				userType = strings.ToLower(strings.TrimSpace(scanner.Text()))
				// cek valid admin/customer
				if userType == "admin" || userType == "customer" {
					break //jika benar break lanjut ke bawah
				}
				fmt.Println("only input 'admin' or 'customer'")
			}

			newUser := model.User{
				Name:     name,
				Email:    email,
				Password: password,
				Phone:    phone,
				Address:  address,
				UserType: userType,
			}
			err := usecase.CreateUserUsecase(db, newUser)
			if err != nil {
				fmt.Println("failed to create a new user:", err)
			} else {
				fmt.Println("Successfully Create a User!")
			}

		case "2":
			fmt.Println("\nShow List All User Accounts")
			// manggil func ShowUsers dari repo
			users, err := repository.ShowUsers(db)
			if err != nil {
				fmt.Println("Failed to retrieve user data", err)
				continue //ngabaikan code di bawah, ngulang menu
			}
			// cek db
			if len(users) == 0 {
				fmt.Println("No Database Users found.")
			} else {
				// bikin header tabel
				fmt.Printf("%-5s | %-20s | %-30s | %-15s | %-10s\n", "ID", "Nama", "Email", "Phone", "Role")
				// garis header nya
				fmt.Println(strings.Repeat("=", 90))
				// loop data untuk cetak setiap data users
				for _, u := range users {
					fmt.Printf("%-5d | %-20s | %-30s | %-15s | %-10s\n", u.ID, u.Name, u.Email, u.Phone, u.UserType)
				}
			}

		case "3":
			fmt.Println("\nUpdate user account data")
			fmt.Print("Enter the User ID to be updated: ")
			scanner.Scan()
			idStr := strings.TrimSpace(scanner.Text())

			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Error: ID must contain number!")
				continue
			}

			fmt.Print("Nama Baru: ")
			scanner.Scan()
			name := strings.TrimSpace(scanner.Text())

			fmt.Print("Phone Baru: ")
			scanner.Scan()
			phone := strings.TrimSpace(scanner.Text())

			fmt.Print("Address Baru: ")
			scanner.Scan()
			address := strings.TrimSpace(scanner.Text())
			// validasi usertype
			var userType string
			for {
				fmt.Print("Enter New Usertype (admin / customer): ")
				scanner.Scan()
				// convert ke lowercase
				userType = strings.ToLower(strings.TrimSpace(scanner.Text()))
				// cek valid admin/customer
				if userType == "admin" || userType == "customer" {
					break //jika benar break lanjut ke bawah
				}
				fmt.Println("only input 'admin' or 'customer'")
			}

			// memanggil func UpdateUser di repo
			err = repository.UpdateUser(db, id, name, phone, address, userType)
			if err != nil {
				fmt.Println("Failed to Update:", err)
			} else {
				fmt.Printf("User Account with ID %d has been updated!\n", id)
			}

		case "4":
			fmt.Println("\nDelete user account")
			fmt.Print("Enter User ID to be Deleted: ")
			scanner.Scan()
			idStr := strings.TrimSpace(scanner.Text())

			//  string id jadi int
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Error: ID must contain numbers!")
				continue //ngulang menu
			}
			// konfirmasi sebelum delete
			fmt.Printf("WARNING!: Are u sure want to permanently delete this ID %d? (y/n): ", id)
			scanner.Scan()
			confirm := strings.ToLower(strings.TrimSpace(scanner.Text()))

			if confirm != "y" {
				fmt.Println("Delletion Canceled. Return to menu...")
				continue //batal kembali ke menu awal
			}
			// panggil func DeleteUser di repo
			err = repository.DeleteUser(db, id)
			if err != nil {
				fmt.Println("failed to delete user:", err)
			} else {
				fmt.Printf("User with ID %d has been deleted permanently from DB!\n", id)
			}

		case "5":
			fmt.Println("\nCreate New Product")

			categories, err := repository.GetCategories(db)
			if err != nil {
				fmt.Println("Failed to load list category:", err)
				continue
			}
			if len(categories) == 0 {
				fmt.Println("no category on db")
				continue
			}
			fmt.Println("Category List:")
			for _, c := range categories {
				fmt.Printf("[%d] %s\n", c.ID, c.Name)
			}

			// input CLI
			fmt.Print("\nChoose Category ID: ")
			scanner.Scan()
			catIDstr := strings.TrimSpace(scanner.Text())
			catID, _ := strconv.Atoi(catIDstr)

			var isCategoryValid bool
			for _, c := range categories {
				if c.ID == catID {
					isCategoryValid = true
					break
				}
			}
			if !isCategoryValid {
				fmt.Println("Error: Category ID not listed! Please select from the list above.")
				continue //ngulangi menu tanpa nanya product name
			}

			fmt.Print("Product's Name: ")
			scanner.Scan()
			prodName := strings.TrimSpace(scanner.Text())

			fmt.Print("Product's Price (Rp): ")
			scanner.Scan()
			priceStr := strings.TrimSpace(scanner.Text())
			price, err := strconv.ParseInt(priceStr, 10, 64)
			if err != nil {
				fmt.Println("Error: The price must be a number without periods or commas!")
				continue
			}

			fmt.Print("Initial Stock Quantity: ")
			scanner.Scan()
			stockStr := strings.TrimSpace(scanner.Text())
			stock, _ := strconv.Atoi(stockStr)

			// bungkus ke model
			newProduct := model.Product{
				CategoryID: catID,
				Name:       prodName,
				Price:      price,
				Stock:      stock,
			}
			// lempar ke usecase
			err = usecase.CreateProductUsecase(db, newProduct)
			if err != nil {
				fmt.Println("Failed to Create Product:", err)
			} else {
				fmt.Printf("Succsessfully added '%s' Product on the Menu!\n", prodName)
			}

		case "6":
			fmt.Println("\nList all product")
			// manggil ListProducts repo
			products, err := repository.ListProducts(db)
			if err != nil {
				fmt.Println("Failed to take Product:", err)
				continue
			}
			// cek apakah isi tabel ksong
			if len(products) == 0 {
				fmt.Println("No product list on DB")
			} else {
				// bikin tabel
				fmt.Printf("%-5s | %-20s | %-30s | %-15s | %-5s\n", "ID", "Kategori", "Nama Produk", "Harga (Rp)", "Stok")
				fmt.Println(strings.Repeat("=", 90))
				// loop untuk cetak setiap product
				for _, p := range products {
					// Memberikan penanda khusus jika stok habis
					stockStatus := strconv.Itoa(p.Stock)
					if p.Stock <= 0 {
						stockStatus = "EMPTY"
					}

					// Mencetak baris data (%d digunakan untuk int/int64, %s untuk string)
					fmt.Printf("%-5d | %-20s | %-30s | %-15d | %-5s\n", p.ID, p.CategoryName, p.Name, p.Price, stockStatus)
				}
			}
		case "7":
			fmt.Println("\nUpdate product data")
		case "8":
			fmt.Println("\nDelete Product")
		case "9":
			fmt.Println("\nList all order")
		case "10":
			fmt.Println("\nReport revenue for all products")
		case "11":
			fmt.Println("\nList product with the highest sales")
		case "12":
			fmt.Println("\nReport user with most frequent order")
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
