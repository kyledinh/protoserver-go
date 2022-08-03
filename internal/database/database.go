package database

import (
	"database/sql"
	"fmt"

	"github.com/kyledinh/protoserver-go/internal/hashing"
	"github.com/kyledinh/protoserver-go/pkg/model"
	"github.com/spf13/viper"
)

func PsqlConnString() string {

	psqlconn := fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable",
		viper.GetString("postgresUser"),
		viper.GetString("postgresPassword"),
		"localhost",
		5432,
		viper.GetString("postgresDB"))

	return psqlconn
}

func InsertNewUser(user model.User) error {
	hashedPassword, err := hashing.HashPassword(user.Password)
	if err != nil {
		return err
	}
	db, err := sql.Open("postgres", PsqlConnString())
	if err != nil {
		return err
	}

	// close database
	defer db.Close()

	// Seed database
	// TODO: scrub inputs for XSS
	query := fmt.Sprintf("INSERT INTO users (email, password, firstname, lastname) VALUES ('%s', '%s', '%s', '%s')",
		user.Email,
		hashedPassword,
		user.Firstname,
		user.Lastname)

	_, err = db.Exec(query)
	return err
}

func FetchUserByEmail(email model.User) (model.User, error) {
	var (
		user model.User
		err  error
	)
	return user, err
}

func FetchUsers(limit int) ([]model.User, error) {
	// default limit 0 to mean: fetch all users
	var err error
	users := make([]model.User, 0)
	return users, err
}