package mongodb

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database
//User info
type File struct {
	ID          bson.ObjectId `bson:"_id" json:"id" `
	Name 	 string       `bson:"name" json:"name"`
}
//Config db connection info struct
type Config struct {
	ConnectionUrl string `json:"connectionUrl"`
	DatabaseName  string `json:"databaseName"`	
}

func (r File) Insert()  (string,error) {
	r.ID = bson.NewObjectId()
	err := db.C("Files").Insert(&r)
	if err!=nil{
		return "false",err
	}
	return "true",nil
}

//Connect Establish a connection to database
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
func LoadConfiguration() {
	config:=Config{}
    configFile, err := os.Open("config.json")
    defer configFile.Close()
    if err != nil {
        fmt.Println(err.Error())
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    Connect(config.ConnectionUrl,config.DatabaseName)
}
