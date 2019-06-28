package repositories

import (
	"gopkg.in/mgo.v2/bson"
)

// Buried ...
type Buried struct {
	ID          bson.ObjectId `bson:"_id" json:"id" `
	UserID      string        `bson:"userid" json:"userid"`
	FileName    string        `bson:"filename" json:"filename"`
	FileHash    string        `bson:"filehash" json:"filehash"`
	PublicHash  string        `bson:"publichash" json:"publichash"`
	Description string        `bson:"description" json:"description"`
	BuriedDate  string        `bson:"burieddate" json:"burieddate"`
	DiggingDate string        `bson:"diggingdate" json:"diggingdate"`
}
