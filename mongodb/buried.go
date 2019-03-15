package mongodb

import (
	"gopkg.in/mgo.v2/bson"
)

type Buried struct {
	Id          	 bson.ObjectId `bson:"_id" json:"id" `
	UserId 	    	 string        `bson:"userid" json:"userid"`
	FileName 		 string        `bson:"filename" json:"filename"`
	FileHash 		 string        `bson:"filehash" json:"filehash"`
	PublicHash 		 string        `bson:"publichash" json:"publichash"`
	Description 	 string        `bson:"description" json:"description"`
	BuriedDate  	 string        `bson:"burieddate" json:"burieddate"`
	DiggingDate 	 string        `bson:"diggingdate" json:"diggingdate"`
}