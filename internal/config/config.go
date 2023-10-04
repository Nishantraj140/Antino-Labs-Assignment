package config

import (
	"Antino-Labs-Assignment/pkg/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var Config *AppConfig

type AppConfig struct {
	DBConfig *sql.DBConfig `json:"db_config"`
}

func ReadConfig(configFile string) (Config *AppConfig) {
	Config = &AppConfig{}
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	if err := json.Unmarshal(file, &Config); err != nil {
		log.Fatalf("unable to marshal config data,err:%v", err)

		return
	}
	fmt.Println("config loaded ", Config)
	fmt.Println(Config.DBConfig.Url)
	fmt.Println(Config.DBConfig.Password)
	fmt.Println(Config.DBConfig.DataBase)
	fmt.Println(Config.DBConfig.UserName)

	return Config
}
