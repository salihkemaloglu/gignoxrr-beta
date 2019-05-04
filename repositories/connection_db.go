package repositories

import (
	
	"fmt"
	"os"
	"time"
	"encoding/json"
	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

// Config db connection struct
type Config struct {
	ConnectionUrl    	 string `json:"connectionUrl"`
	DatabaseName     	 string `json:"databaseName"`	
	DatabaseUsername     string `json:"databaseUsername"`	
	DatabasePassword     string `json:"databasePassword"`	
}

// Connect Establish a connection to database
func Connect(con_ Config) {
	info := &mgo.DialInfo{
		Addrs:    []string{con_.ConnectionUrl},
		Timeout:  5 * time.Second,
		Database: con_.DatabaseName,
		Username: con_.DatabaseUsername,
		Password: con_.DatabasePassword,
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		fmt.Println(err.Error())
	}
	db = session.DB(con_.DatabaseName)
}

//LoadConfiguration Parse the configuration file 'config.toml', and establish a connection to DB
func LoadConfiguration(connectionType_ string) string {
	
	config:=Config{}

	if connectionType_ == "dev" {
		configFile, err := os.Open("app_root/config_files/dev.json")
		defer configFile.Close()
		if err != nil {
			return fmt.Sprintf("Config file err: %v", err.Error())
		}
		jsonParser := json.NewDecoder(configFile)
		jsonParser.Decode(&config)
	}else if connectionType_ == "prod" {
		configFile, err := os.Open("app_root/config_files/prod.json")
		defer configFile.Close()
		if err != nil {
			return fmt.Sprintf("Config file err: %v", err.Error())
		}
		jsonParser := json.NewDecoder(configFile)
		jsonParser.Decode(&config)
	}else if connectionType_ == "local" {
		configFile, err := os.Open("app_root/config_files/local.json")
		defer configFile.Close()
		if err != nil {
			return fmt.Sprintf("Config file err: %v", err.Error())
		}
		jsonParser := json.NewDecoder(configFile)
		jsonParser.Decode(&config)
	}
   
	Connect(config)
	return "ok"
}
