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
// Crud operaions for Folder
func (r Folder) GetFolder() (*Folder, error) {
	err := db.C("folder").FindId(r.ID).One(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}
func (r Folder) GetAllFolders() ([]Folder, error) {
	var folders []Folder
	err := db.C("folder").Find(bson.M{}).All(&folders)
	if err != nil {
		return nil, err
	}
	return folders, err
}

func (r Folder) Insert()  error {
	r.ID = bson.NewObjectId()
	err := db.C("folder").Insert(&r)
	if err!=nil{
		return err
	}
	return nil
}

func (r Folder) Update() error {
	err := db.C("folder").Update(bson.M{"_id": r.ID}, &r)
	if err!=nil {
		return err
	}
	return nil
}
func (r Folder) Delete() error {
	err := db.C("folder").Remove(&r)
	if err!=nil {
		return err
	}
	return nil
}