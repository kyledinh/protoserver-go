package migration

import (
	"database/sql"
	"fmt"

	"github.com/kyledinh/protoserver-go/internal/database"
	_ "github.com/lib/pq"
)

// TODO: Implement transaction and rollback support
// https://data-nerd.blog/2020/06/07/implementing-transactions-in-postgresql-using-go-lang-database-sql-package/

// Default to localhost
func Initialize() error {
	psqlconn := database.PsqlConnString()

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return err
	}

	// close database
	defer db.Close()

	// create table
	queries := []string{
		`DROP TABLE IF EXISTS users`,
		`CREATE TABLE users(
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			firstname VARCHAR(255) NOT NULL,
			lastname VARCHAR(255) NOT NULL
		)`,
		`DROP TABLE IF EXISTS permissions`,
		`CREATE TABLE permissions(
			email VARCHAR(255) UNIQUE NOT NULL,
			permission VARCHAR(255) NOT NULL
		)`,
	}
	for _, q := range queries {
		_, err = db.Exec(q)
		if err != nil {
			return err
		}
	}

	return err
}

func Ping() error {

	psqlconn := database.PsqlConnString()

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return err
	}

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Connected!")
	return err
}

func Seed() error {

	psqlconn := database.PsqlConnString()

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return err
	}

	// close database
	defer db.Close()

	// Seed database
	queries := []string{
		`INSERT INTO users (email, password, firstname, lastname) VALUES ('kyle@email.com','un-salted', 'Kyle', 'Dinh')`,
	}
	for _, q := range queries {
		_, err = db.Exec(q)
		return err
	}

	fmt.Println("Connected!")
	return err
}
