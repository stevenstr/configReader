package main

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	configFile = "config.json"
	configType = "json"
)

type ConfigStructure struct {
	Linux        string `mapstructure:"linux"`
	Windows      string `mapstructure:"windows"`
	Active       bool   `mapstructure:"active"`
	Postgres     string `mapstructure:"postgres"`
	PostgresUser string `mapstructure:"postgresUser"`
	PostgresPort int    `mapstructure:"postgresPort"`
	Mongodb      string `mapstructure:"mongodb"`
	MongodbUser  string `mapstructure:"mongodbUser"`
	MongodbPort  int    `mapstructure:"mongodbPort"`
}

func main() {
	viper.SetConfigType(configType)
	viper.SetConfigFile(configFile)

	fmt.Println("config file - ", viper.ConfigFileUsed(), " - was used.")
}
