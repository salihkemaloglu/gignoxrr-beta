package repositories

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

// Config db connection struct
type Config struct {
	ConnectionURL    string `json:"connectionUrl"`
	DatabaseName     string `json:"databaseName"`
	DatabaseUsername string `json:"databaseUsername"`
	DatabasePassword string `json:"databasePassword"`
}

// Connect Establish a connection to database
func Connect(con Config) {
	info := &mgo.DialInfo{
		Addrs:    []string{con.ConnectionURL},
		Timeout:  5 * time.Second,
		Database: con.DatabaseName,
		Username: con.DatabaseUsername,
		Password: con.DatabasePassword,
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		fmt.Println(err.Error())
	}
	db = session.DB(con.DatabaseName)
}

//LoadConfiguration Parse the configuration file 'config.toml', and establish a connection to DB
func LoadConfiguration(connectionType string) string {

	config := Config{}

	if connectionType == "dev" {
		configFile, err := os.Open("app_root/config_files/dev.json")
		defer configFile.Close()
		if err != nil {
			return fmt.Sprintf("Config file err: %v", err.Error())
		}
		jsonParser := json.NewDecoder(configFile)
		jsonParser.Decode(&config)
	} else if connectionType == "prod" {
		configFile, err := os.Open("app_root/config_files/prod.json")
		defer configFile.Close()
		if err != nil {
			return fmt.Sprintf("Config file err: %v", err.Error())
		}
		jsonParser := json.NewDecoder(configFile)
		jsonParser.Decode(&config)
	} else if connectionType == "local" {
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
