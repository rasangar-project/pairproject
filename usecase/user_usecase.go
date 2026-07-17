package usecase

import (
	"database/sql"
	"errors"
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
