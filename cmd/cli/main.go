package main

import (
	"flag"
	"log"
	"os"

	"github.com/kyledinh/protoserver-go/internal/migration"
	"github.com/kyledinh/protoserver-go/pkg/config"
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
	migrate := flag.String("migrate", "", "Usage: migrate (ping | initialize | up | down)")

	flag.Parse()
	args := flag.Args()
	_ = args

	var (
		outBytes []byte
	)

	config.LoadConfig()

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
