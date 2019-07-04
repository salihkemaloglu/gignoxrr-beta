package repositories

import (
	"gopkg.in/mgo.v2/bson"
)

//Folder ...
type Folder struct {
	ID          bson.ObjectId `bson:"_id" json:"id" `
	Name        string        `bson:"name" json:"name"`
	UserID      string        `bson:"userid" json:"userid"`
	CreatedDate string        `bson:"createddate" json:"createddate"`
	UpdatedDate string        `bson:"updateddate" json:"updateddate"`
}

//GetFolder ...
// Crud operaions for Folder
func (r Folder) GetFolder() (*Folder, error) {
	err := db.C("Folder").FindId(r.ID).One(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}

//GetAllFolders ...
func (r Folder) GetAllFolders() ([]Folder, error) {
	var folders []Folder
	err := db.C("Folder").Find(bson.M{}).All(&folders)
	if err != nil {
		return nil, err
	}
	return folders, err
}

//Insert ...
func (r Folder) Insert() error {
	r.ID = bson.NewObjectId()
	err := db.C("Folder").Insert(&r)
	if err != nil {
		return err
	}
	return nil
}

//Update ...
func (r Folder) Update() error {
	err := db.C("Folder").Update(bson.M{"_id": r.ID}, &r)
	if err != nil {
		return err
	}
	return nil
}

//Delete ...
func (r Folder) Delete() error {
	err := db.C("Folder").Remove(&r)
	if err != nil {
		return err
	}
	return nil
}
