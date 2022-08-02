package psql

import (
	"fmt"

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
