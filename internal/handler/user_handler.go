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
	"text/tabwriter"
	"time"
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
			fmt.Print("Enter the Product ID to be updated: ")
			scanner.Scan()
			idStr := strings.TrimSpace(scanner.Text())
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Error: Product ID must contain number!")
				continue
			}

			// Nampilin daftar kategori lagi sebagai panduan untuk Admin
			categories, err := repository.GetCategories(db)
			if err != nil || len(categories) == 0 {
				fmt.Println("failed to load category.")
				continue
			}
			fmt.Println("\nList Available Category:")
			for _, c := range categories {
				fmt.Printf("[%d] %s\n", c.ID, c.Name)
			}

			fmt.Print("\nChoose new ID Category: ")
			scanner.Scan()
			catIDStr := strings.TrimSpace(scanner.Text())
			catID, err := strconv.Atoi(catIDStr)
			if err != nil {
				fmt.Println("Error: Category ID must contain number!")
				continue
			}

			// Validasi apakah kategori yang diinput ada di database (seperti di Case 5)
			var isCategoryValid bool
			for _, c := range categories {
				if c.ID == catID {
					isCategoryValid = true
					break
				}
			}
			if !isCategoryValid {
				fmt.Println("Error: Category ID not registered! Update cancelled.")
				continue
			}

			fmt.Print("Name a New Product: ")
			scanner.Scan()
			prodName := strings.TrimSpace(scanner.Text())

			fmt.Print("New Product Price (Rp): ")
			scanner.Scan()
			priceStr := strings.TrimSpace(scanner.Text())
			// Gunakan ParseInt untuk BIGINT / int64
			price, err := strconv.ParseInt(priceStr, 10, 64)
			if err != nil {
				fmt.Println("Error: The price must be a number (without periods or commas)!")
				continue
			}

			fmt.Print("New Product Stock: ")
			scanner.Scan()
			stockStr := strings.TrimSpace(scanner.Text())
			stock, err := strconv.Atoi(stockStr)
			if err != nil {
				fmt.Println("Error: Stock must be a whole number!") //Stok harus berupa angka bulat
				continue
			}

			// Lempar semua data ke repository
			err = repository.UpdateProduct(db, id, catID, prodName, price, stock)
			if err != nil {
				fmt.Println("Failed to update product:", err)
			} else {
				fmt.Printf("Product data with ID %d has been successfully updated.\n", id)
			}
		case "8":
			fmt.Println("\nDelete Product")
			fmt.Print("Enter Product ID to be Deleted: ")
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
			// panggil func DeleteProduct di repo
			err = repository.DeleteProduct(db, id)
			if err != nil {
				fmt.Println("failed to delete Product:", err)
			} else {
				fmt.Printf("Product with ID %d has been deleted permanently from DB!\n", id)
			}
		case "9":
			fmt.Println("\nList all order")
			orders, err := repository.ListAllOrders(db)
			if err != nil {
				fmt.Println("Error loading order:", err)
				continue
			}

			if len(orders) == 0 {
				fmt.Println("No transactions yet.")
			} else {
				fmt.Printf("%-5s | %-20s | %-15s | %-15s\n", "ID", "Customer Name", "Status", "Total Amount")
				fmt.Println(strings.Repeat("-", 65))
				for _, o := range orders {
					fmt.Printf("%-5d | %-20s | %-15s | Rp.%-12d\n", o.ID, o.CustomerName, o.Status, o.TotalAmount)
				}
			}
		case "10":
			displayRevenue(db)
		case "11":
			fmt.Println("\nList product with the highest sales")
			topProducts, err := repository.GetHighestSalesProducts(db)
			if err != nil {
				fmt.Println("Error loading data:", err)
				continue
			}

			if len(topProducts) == 0 {
				fmt.Println("No products have been sold yet.")
			} else {
				fmt.Printf("%-30s | %-15s\n", "Product's Name", "Total Sold")
				fmt.Println(strings.Repeat("-", 48))
				for _, t := range topProducts {
					fmt.Printf("%-30s | %-15d pcs\n", t.Name, t.Sales)
				}
			}
		case "12":
			fmt.Println("\nReport user with most frequent order")
			// manggil func GetMostFrequentUsers dari repo
			reports, err := repository.GetMostFrequentUsers(db)
			if err != nil {
				fmt.Println("Failed to load:", err)
				continue
			}

			//cek tabel order jika masih kosong
			if len(reports) == 0 {
				fmt.Println("No data transaction on the DB")
			} else {
				// header tabel
				fmt.Printf("%-5s | %-20s | %-30s | %-15s\n", "ID", "Nama Customer", "Email", "Total Orders")
				fmt.Println(strings.Repeat("=", 90))
				// isi tabel
				for _, r := range reports {
					fmt.Printf("%-5d | %-20s | %-30s | %-15d\n", r.UserID, r.Name, r.Email, r.TotalOrders)
				}
			}

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
			var order_item []model.Order_items
			order_item, err := makeOrders(db, scanner)
			if err != nil {
				fmt.Println("failed to retrieve order, please try again!")
				continue
			}
			displayPayMethod(db)
			payMethod := choosePaymentMethod(db, scanner)
			displayOverallOrder(db, order_item)
			// var insertSuccess bool
			if orderConfirmation(scanner) {
				totalPrice := sumOrderPrice(order_item)
				err := repository.CreatingOrdersDB(db, customerID, totalPrice, payMethod, order_item)
				if err != nil {
					fmt.Println("Error inserting order database:", err)
					continue
				}
				fmt.Println("Please wait while Process your order! ")
				time.Sleep(3 * time.Second)
				fmt.Println("Thank you for waiting, here's your order!")
				continue
			} else if !orderConfirmation(scanner) {
				fmt.Println("order cancelled")
			}
		case "2":
			displayAllProducts(db)
		case "3":
			displayUserHistory(db, customerID)
		case "0":
			fmt.Println("/nLoggingOut...")
			return
		default:
			fmt.Println("Wrong Input!")
		}
	}
}

// push arjun func customer
func displayAllProducts(db *sql.DB) {
	products, err := repository.FetchingProducts(db)
	if err != nil {
		fmt.Println("Error Fetching Products:", err)
		return
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tStocks\tPrice")
	fmt.Fprintln(w, "--\t-----\t-----\t-------")
	for _, p := range products {
		fmt.Fprintf(w, "%d\t%s\t%d\tRp.%d\n", p.ID, p.Name, p.Stock, p.Price)
	}
	w.Flush()
}

func displayAddOns(db *sql.DB, product_id int) {
	add_ons, err := repository.FetchProductAddOn(db, product_id)
	if err != nil {
		fmt.Println("Error Fetching Products:", err)
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tPrice")
	fmt.Fprintln(w, "--\t-----\t-----")
	for _, a := range add_ons {
		fmt.Fprintf(w, "%d\t%s\tRp.%d\n", a.ID, a.Name, a.Price)
	}
	w.Flush()
}

func countOrderSubtotal(add_onPrice int, productPrice int, quantity int) int {
	pricePerProduct := add_onPrice + productPrice
	subtotal := pricePerProduct * quantity
	return subtotal
}

func makeOrders(db *sql.DB, scanner *bufio.Scanner) ([]model.Order_items, error) {
	var order_item []model.Order_items
	for {
		displayAllProducts(db)
		var (
			product_id  int
			add_on_id   int
			quantity    int
			unit_price  int
			subtotal    int
			add_onPrice int
			note        string
		)

		fmt.Println("Enter products ID you want to order ->")
		scanner.Scan()
		inputProduct := strings.TrimSpace(scanner.Text())
		inputProductNum, err := strconv.Atoi(inputProduct)
		if err != nil {
			fmt.Println("Conversion failed! The input is not a valid number:", err)
			continue
		}
		if repository.CheckValidProductId(db, inputProductNum) {
			product_id = inputProductNum
		} else {
			fmt.Println("your product id is invalid, please input valid product id!")
			continue
		}
		displayAddOns(db, product_id)
		fmt.Println("Enter add on ID you want to add or enter 0 to skip add on ->")
		scanner.Scan()
		inputAddOn := strings.TrimSpace(scanner.Text())
		inputAddOnNum, err := strconv.Atoi(inputAddOn)
		if inputAddOnNum == 0 {
			add_onPrice = 0
		} else if repository.ChecksValidAddOn(db, product_id, inputAddOnNum) {
			add_on_id = inputAddOnNum
			add_onPrice, err = repository.GetAddOnPrice(db, add_on_id)
			if err != nil {
				fmt.Println("fetching Add on price failed! please try again", err)
				continue
			}
		} else {
			fmt.Println("your add on id is invalid! please try again")
			continue
		}

		fmt.Println("Enter your order item quantity ->")
		scanner.Scan()
		inputQuantity := strings.TrimSpace(scanner.Text())
		inputQuantityNum, err := strconv.Atoi(inputQuantity)
		quantity = inputQuantityNum
		unit_price, err = repository.GetProductPrice(db, product_id)
		if err != nil {
			fmt.Println("fetching product price failed! please try again", err)
			continue
		}
		subtotal = countOrderSubtotal(add_onPrice, unit_price, quantity)
		fmt.Println("Leave a note for your order:")
		scanner.Scan()
		note = strings.TrimSpace(scanner.Text())

		order_item = append(order_item, model.Order_items{Product_id: product_id, Add_on_id: add_on_id, Quantity: quantity, Unit_price: unit_price, Subtotal: subtotal, Note: note})
		var again string
		for {
			fmt.Println("do you want to add another product to your order? Y/N")
			scanner.Scan()
			inputAgain := strings.ToLower(strings.TrimSpace(scanner.Text()))
			if inputAgain == "y" || inputAgain == "n" {
				again = inputAgain
				break
			} else {
				fmt.Println("invalid input")
			}
		}
		if again == "n" {
			break
		}
	}
	return order_item, nil
}

func displayPayMethod(db *sql.DB) {
	pay_method, err := repository.FetchPaymentMethod(db)
	if err != nil {
		fmt.Println("Error Fetching Products:", err)
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tName")
	fmt.Fprintln(w, "--\t-----")
	for _, p := range pay_method {
		fmt.Fprintf(w, "%d\t%s\n", p.ID, p.Name)
	}
	w.Flush()
}

func choosePaymentMethod(db *sql.DB, scanner *bufio.Scanner) int {
	var payMethodId int
	fmt.Println("Choose your payment method ->")
	for {
		scanner.Scan()
		inputMethod := strings.TrimSpace(scanner.Text())
		payMethodNum, err := strconv.Atoi(inputMethod)
		if err != nil {
			fmt.Println("Conversion failed! The input is not a valid number:", err)
			continue
		}
		if repository.CheckPayMethodValid(db, payMethodNum) {
			payMethodId = payMethodNum
			break
		}
	}
	return payMethodId
}

func displayOverallOrder(db *sql.DB, order_items []model.Order_items) {
	fmt.Println("here's your overall order")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "Name\tQuantity\tAdd on\tPrice")
	fmt.Fprintln(w, "-----\t-----\t-------\t------")

	for _, item := range order_items {
		productName, err := repository.GetProductName(db, item.Product_id)
		if err != nil {
			fmt.Println("error fetching product name", err)
		}
		var addOnName string
		if item.Add_on_id == 0 {
			addOnName = "-"
		} else {
			addOnName, err = repository.GetAddOnName(db, item.Add_on_id)
			if err != nil {
				fmt.Println("error fetching Add On name", err)
				addOnName = ""
			}
		}
		fmt.Fprintf(w, "%s\t%d\t%s\tRp.%d\n", productName, item.Quantity, addOnName, item.Subtotal)
	}
	w.Flush()
}

func orderConfirmation(scanner *bufio.Scanner) bool {
	var confirmInput bool
	fmt.Println("Type 'confirm' to confirm your order or 'cancel' to cancel your order ")

	for {
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		formattedInput := strings.ToLower(input)
		if formattedInput == "confirm" {
			confirmInput = true
			break
		} else if formattedInput == "cancel" {
			confirmInput = false
			break
		} else {
			fmt.Println("invalid input, please input correct option")
			continue
		}

	}
	return confirmInput
}

func sumOrderPrice(order_items []model.Order_items) int {
	var totalPrice int

	for _, item := range order_items {
		totalPrice += item.Subtotal
	}
	return totalPrice
}

func displayUserHistory(db *sql.DB, customerID int) {
	userHistory, err := repository.GetUserHistory(db, customerID)
	if err != nil {
		fmt.Println("Error Fetching Products:", err)
		return
	}
	fmt.Println("here is your order history")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "Product Name\tQuantity\tPrice")
	fmt.Fprintln(w, "--------------\t---------\t------")
	for _, u := range userHistory {
		fmt.Fprintf(w, "%s\t%d\tRp.%d\n", u.ProductName, u.Quantity, u.Subtotal)
	}
	w.Flush()
}

func displayRevenue(db *sql.DB) {
	revenue, err := repository.GenerateTotalRevenue(db)
	if err != nil {
		fmt.Println("error generating revenue:", err)
		return
	}
	fmt.Printf(" **TOTAL REALIZED REVENUE** \n")
	fmt.Printf(" Rp.%d\n", revenue.Revenue)
}
