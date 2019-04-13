package mongodb

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
	ConnectionUrl    string `json:"connectionUrl"`
	DatabaseName     string `json:"databaseName"`	
}

// Connect Establish a connection to database
func Connect(connectionUrl string,databaseName string) {
	info := &mgo.DialInfo{
		Addrs:    []string{connectionUrl},
		Timeout:  5 * time.Second,
		Database: databaseName,
		Username: "",
		Password: "",
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		fmt.Println(err.Error())
	}
	db = session.DB(databaseName)
}

//LoadConfiguration Parse the configuration file 'config.toml', and establish a connection to DB
func LoadConfiguration(connectionType bool) string {
	
	config:=Config{}

	if connectionType {
		configFile, err := os.Open("app-root/config-files/config-prod.json")
		defer configFile.Close()
		if err != nil {
			return fmt.Sprintf("Config file err: %v", err.Error())
		}
		jsonParser := json.NewDecoder(configFile)
		jsonParser.Decode(&config)
	}else {
		configFile, err := os.Open("app-root/config-files/config-local.json")
		defer configFile.Close()
		if err != nil {
			return fmt.Sprintf("Config file err: %v", err.Error())
		}
		jsonParser := json.NewDecoder(configFile)
		jsonParser.Decode(&config)
	}
   
	Connect(config.ConnectionUrl,config.DatabaseName)
	return "ok"
}
