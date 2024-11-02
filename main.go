package main

import (
	"fmt"
	"os"

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
	PostgresPort string `mapstructure:"postgresPort"`
	Mongodb      string `mapstructure:"mongodb"`
	MongodbUser  string `mapstructure:"mongodbUser"`
	MongodbPort  string `mapstructure:"mongodbPort"`
}

func main() {
	viper.SetConfigFile(configFile)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("try to unmarshal")
	var cfgData ConfigStructure
	viper.Unmarshal(&cfgData)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println("go structure was recieved:")
	fmt.Printf("%#v\n", cfgData)
}
