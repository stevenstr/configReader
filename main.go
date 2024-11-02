package main

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	configFile = "config.json"
	configType = "json"
	sep        = '\n'
)

type ConfigStructure struct {
	Linux        string `json:"linux"`
	Windows      string `json:"windows"`
	Active       bool   `json:"active"`
	Postgres     string `json:"postgres"`
	PostgresUser string `json:"postgresUser"`
	PostgresPort string `json:"postgresPort"`
	Mongodb      string `json:"mongodb"`
	MongodbUser  string `json:"mongodbUser"`
	MongodbPort  string `json:"mongodbPort"`
}

func main() {
	viper.SetConfigFile(configFile)
	viper.SetConfigType(configType)
	viper.ReadInConfig()

	var cfgData ConfigStructure
	viper.Unmarshal(&cfgData)
	fmt.Println(cfgData)
}
