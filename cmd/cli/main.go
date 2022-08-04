package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kyledinh/protoserver-go/internal/database"
	"github.com/kyledinh/protoserver-go/internal/migration"
	"github.com/kyledinh/protoserver-go/pkg/config"
	"github.com/kyledinh/protoserver-go/pkg/model"
	"github.com/kyledinh/protoserver-go/pkg/proto/protoerr"
)

func errCheckLogFatal(err error, me *error) {
	if err != nil {
		log.Fatal(protoerr.NewWrappedError(err.Error(), me))
		// panic(err)
	}
}

func main() {

	// PARSE INPUT
	migrate := flag.String("migrate", "", "Usage: -migrate (ping | initialize | up | down)")
	dbUser := flag.String("dbuser", "", "Usage: -dbuser (email)")
	dbAllUsers := flag.Bool("dbusers", false, "Usage: -dbusers fetches all users")

	flag.Parse()
	args := flag.Args()
	_ = args

	var (
		outBytes []byte
	)

	config.LoadConfig()

	if *dbAllUsers {
		var buf bytes.Buffer
		users, _ := database.FetchAllUsers()
		for _, user := range users {
			buf.WriteString(fmt.Sprintf("Email: %s, Firstname: %s, Lastname: %s\n", user.Email, user.Firstname, user.Lastname))
		}
		os.Stdout.Write(buf.Bytes())
		os.Exit(0)
	}

	if *dbUser != "" { // the email
		var buf bytes.Buffer

		u := model.User{Email: *dbUser}
		user, err := database.FetchUserByEmail(u)
		if err == nil {
			buf.WriteString(fmt.Sprintf("=== Email: %s, Firstname: %s, Lastname: %s\n", user.Email, user.Firstname, user.Lastname))
		}
		os.Stdout.Write(buf.Bytes())
		os.Exit(0)
	}

	// MAIN SWITCH
	if *migrate == "initialize" {
		err := migration.Initialize()
		if err != nil {
			outBytes = []byte(err.Error())
		} else {
			outBytes = []byte(`Database Initialized!`)
		}
		os.Stdout.Write(outBytes)
		os.Exit(0)
	}

	if *migrate == "seed" {
		err := migration.Seed()
		if err != nil {
			outBytes = []byte(err.Error())
		} else {
			outBytes = []byte(`Database Seeded!`)
		}
		os.Stdout.Write(outBytes)
		os.Exit(0)
	}

	if *migrate == "ping" {
		err := migration.Ping()
		if err != nil {
			outBytes = []byte(err.Error())
		} else {
			outBytes = []byte(`Pinged! Database OK`)
		}
		os.Stdout.Write(outBytes)
		os.Exit(0)
	}
}
