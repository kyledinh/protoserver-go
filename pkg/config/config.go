package config

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/kyledinh/protoserver-go/pkg/model"
	"github.com/kyledinh/protoserver-go/pkg/proto"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Ready bool
var RouteConfig model.RouteConfig

func init() {
	log.Println("... config init()")
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
}

func LoadConfig() {
	ctx := context.Background()

	log.Println("... config LoadConfig() ...")
	viper.SetDefault("testMode", false)
	viper.SetDefault("serverPort", 8000)
	viper.SetDefault("serviceName", "protoserver")

	viper.SetDefault("jwtSecret", "SUPER_SECRET_JWT_KEY")
	viper.SetDefault("postgresDB", "proto")
	viper.SetDefault("postgresUser", "postgres")
	viper.SetDefault("postgresPassword", "postgres")

	viper.SetDefault("authtokens", []string{"PROTOSERVER_DEV_TOKEN", "PROTOSERVER_TEST_TOKEN"})
	viper.SetDefault("upperLimit", 100)
	viper.SetDefault("timeout", 100)
	viper.SetDefault("clientTimeout", time.Duration(5*time.Second))

	viper.SetDefault("log.format", "json")
	viper.SetDefault("log.type", "stderr")

	// Load "PROTOSERVER_*VARNAME*" environment vars into viper; ie PROTOSERVER_VERSION=1.0.13, will be accessable with viper.GetString("version")
	viper.SetEnvPrefix("PROTOSERVER")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Load server configuration file
	viper.AddConfigPath("/etc/")

	wd, err := os.Getwd()
	if err != nil {
		log.Panic("unable to get the current working directory", zap.Error(err))
	}
	for len(wd) > 0 {
		viper.AddConfigPath(wd)
		wd = wd[:strings.LastIndexByte(wd, '/')]
	}

	viper.SetConfigName("protoserver") //  /etc/protoserver.json
	if err := viper.ReadInConfig(); err != nil {
		log.Panic("Fatal error loading config file", zap.Error(err))
	}

	// Load the routes configuration
	data, err := ioutil.ReadFile("/etc/protoserver-routes.json")
	if err != nil {
		log.Println("Error reading /etc/protoserver-routes.json: ", err)
	} else {
		log.Println("/etc/protoserver-routes.json")
		log.Println(string(data))
		errUnmarshal := json.Unmarshal(data, &RouteConfig.Routes)
		if errUnmarshal != nil {
			log.Println("!! Error unmarshaling protoserver-routes.json: ", err)
		}
	}

	// Setup structured logger
	proto.SetupLogger(ctx, viper.GetString("serviceName"))
	proto.DefaultLogger().Info("Setup has set up logger to json output........... ")
	log.Printf("... CurrentConfig %v", CurrentConfig())

	proto.DefaultLogger().Info("current config", zap.String("version", viper.GetString("version")), zap.String("log level", viper.GetString("log.level")))

	Ready = true
}

func IsReady() bool {
	return Ready
}

func CurrentConfig() string {
	cc := prettyPrint(RouteConfig)
	return cc
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
