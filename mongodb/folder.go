package mongodb

import (
	"gopkg.in/mgo.v2/bson"
)

type Folder struct {
	ID           	 bson.ObjectId `bson:"_id" json:"id" `
	Name 	   	 	 string        `bson:"name" json:"name"`
	UserId 	    	 string        `bson:"userid" json:"userid"`
	CreatedDate 	 string        `bson:"createddate" json:"createddate"`
	UpdatedDate 	 string        `bson:"updateddate" json:"updateddate"`
}