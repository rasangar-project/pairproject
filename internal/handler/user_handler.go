package handler

import (
	"GamingStore/internal/repository"
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)



func Run(db *sql.DB) {
	fmt.Println("\n----Game Store CLI----")
	fmt.Println("Type 'exit' to quit.")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n Enter comand >")
		scanner.Scan()
		input:= scanner.Text()
		command := strings.TrimSpace(strings.ToLower(input))
		if command == "exit" {
			fmt.Println("Closing Gamestore CLI. Goodbye")
			break
		}
		switch command {
		case "customers" :
			customerMenu(db, scanner)
		}
	}
	
}

func promptUser(scanner *bufio.Scanner, question string) string {
		fmt.Print(question)
		scanner.Scan()
		return strings.TrimSpace(scanner.Text())
	}

func customerMenu(db *sql.DB, scanner *bufio.Scanner) {

	for {
		fmt.Println("\n --- Customer Module ---")
		fmt.Println("Commands: create, list, delete, search, back\n Enter command:")

		scanner.Scan()
		command := strings.TrimSpace(strings.ToLower(scanner.Text()))
		switch command {
		case "create": 
			name := promptUser(scanner, "Enter customer name:")
			email := promptUser(scanner, "Enter customer email:")
			phone := promptUser(scanner, "Enter customer phone number:")
			err:= repository.CreateCustomer(db, name, email, phone)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case "list" :
			customers,err := repository.ListCustomer(db)
			if err != nil {
				fmt.Println("Error Fetching Products:",err)
				continue
			}
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
			fmt.Fprintln(w, "ID\tName\tEmail\tPhone\tCreated At")
			fmt.Fprintln(w, "--\t-----\t-----\t-------\t------")
			for _,c := range customers {
				fmt.Fprintf(w,"%d\t%s\t%s\t%s\t%s\n", c.ID, c.Name, c.Email, c.Phone, c.Created_at)
			}
			w.Flush()
		case "back":
			return
		default: 
			fmt.Println("Unknown command. Type 'help' for a list of commands")
	}
	}

	
}