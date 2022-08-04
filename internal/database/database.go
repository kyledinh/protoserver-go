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
	query := fmt.Sprintf("INSERT INTO users (email, password, firstname, lastname) VALUES ('%s', '%s', '%s', '%s')",
		user.Email,
		hashedPassword,
		user.Firstname,
		user.Lastname)

	_, err = db.Exec(query)
	return err
}

func ValidateLogin(email, password string) (bool, error) {
	var dbHash string
	db, err := sql.Open("postgres", PsqlConnString())
	if err != nil {
		return false, err
	}

	// close database
	defer db.Close()

	query := fmt.Sprintf("SELECT password FROM users WHERE email = '%s')", email)

	row := db.QueryRow(query)
	err = row.Scan(&dbHash)
	if err != nil {
		return false, err
	}

	userHash, err := hashing.HashPassword(password)
	if err != nil {
		return false, err
	}

	validLogin := dbHash == userHash
	return validLogin, err
}

func FetchUserByEmail(email model.User) (model.User, error) {
	var (
		user model.User
		err  error
	)

	db, err := sql.Open("postgres", PsqlConnString())
	if err != nil {
		return user, err
	}

	// close database
	defer db.Close()

	query := fmt.Sprintf("SELECT email, firstname, lastname FROM users WHERE email = '%s'", email.Email)

	row := db.QueryRow(query)
	err = row.Scan(&user.Email, &user.Firstname, &user.Lastname)
	if err != nil {
		return user, err
	}

	return user, err
}

func FetchAllUsers() ([]model.User, error) {
	var err error
	users := make([]model.User, 0)

	db, err := sql.Open("postgres", PsqlConnString())
	if err != nil {
		return nil, err
	}

	// close database
	defer db.Close()

	rows, err := db.Query("SELECT email, firstname, lastname FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var u model.User
		err = rows.Scan(&u.Email, &u.Firstname, &u.Lastname)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	// get any error encountered during iteration
	err = rows.Err()
	return users, err
}
