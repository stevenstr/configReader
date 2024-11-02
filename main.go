package main

import "github.com/spf13/viper"

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

// func ParseConfigs() (ConfigStructure, error) {
// 	stream, err := os.Open(configFile)
// 	if err != nil {
// 		return ConfigStructure{}, err
// 	}
// 	defer stream.Close()

// 	reader := bufio.NewReader(stream)

// 	resString := ""
// 	for {
// 		line, err := reader.ReadString(sep)
// 		if err == io.EOF {
// 			resString += line
// 			break
// 		}
// 		if err != nil {
// 			break
// 		}
// 		resString += line
// 	}

// 	fmt.Println("Try to unmarshal configs...")
// 	configs := ConfigStructure{}
// 	err = json.Unmarshal([]byte(resString), &configs)
// 	if err != nil {
// 		return ConfigStructure{}, err
// 	}

// 	return configs, err
// }

func main() {
	// myConfigs, err := ParseConfigs()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Println("We recieved this structure:")
	// fmt.Println(myConfigs)

	viper.SetConfigFile(configFile)
	viper.SetConfigType(configType)
	viper.ReadInConfig()

}
