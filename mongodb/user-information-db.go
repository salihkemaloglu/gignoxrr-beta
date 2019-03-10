package mongodb

import (
	// "encoding/json"
	// "fmt"
	// "os"
	// "time"
	// mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
type UserInformation struct {
	ID          	 bson.ObjectId `bson:"_id" json:"id" `
	UserId 	 		 string        `bson:"userid" json:"userid"`
	ImagePath 		 string        `bson:"imagepath" json:"imagepath"`
	Description 	 string        `bson:"description" json:"description"`
	TotalSpace  	 string        `bson:"totalspace" json:"totalspace"`
	CreatedDate 	 string        `bson:"createddate" json:"createddate"`
	UpdatedDate 	 string        `bson:"updateddate" json:"updateddate"`
	LanguageType 	 string        `bson:"languagetype" json:"languagetype"`
}