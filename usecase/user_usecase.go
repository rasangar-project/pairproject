package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"pairproject/internal/model"
	"pairproject/internal/repository"
)

func CreateUserUsecase(db *sql.DB, user model.User) error {
	// Validasi data kosong
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("nama, email, & password must fill")
	}
	// mastiin role valid
	if user.UserType != "admin" && user.UserType != "customer" {
		return errors.New("user type must admin or customer!")
	}
	// jika semua valid panggil repo
	err := repository.InsertUser(db, user)
	if err != nil {
		return err //balikin ke handler jika gagal insert
	}
	return nil
}

func CreateProductUsecase(db *sql.DB, p model.Product) error {
	if p.Name == "" {
		return errors.New("Product name should not Null")
	}
	if p.Price <= 0 {
		return errors.New("product price must more than 0")
	}
	if p.Stock < 0 {
		return errors.New("stock must >= 0")
	}
	if p.CategoryID <= 0 {
		return errors.New("Category is not valid")
	}
	// next jika lolos
	err := repository.InsertProduct(db, p)
	if err != nil {
		return fmt.Errorf("Failed to insert DB: %v", err)
	}
	return nil
}
