package repositories

import (
	"gopkg.in/mgo.v2/bson"
)
type Folder struct {
	Id           	 bson.ObjectId `bson:"_id" json:"id" `
	Name 	   	 	 string        `bson:"name" json:"name"`
	UserId 	    	 string        `bson:"userid" json:"userid"`
	CreatedDate 	 string        `bson:"createddate" json:"createddate"`
	UpdatedDate 	 string        `bson:"updateddate" json:"updateddate"`
}

// Crud operaions for Folder
func (r Folder) GetFolder() (*Folder, error) {
	err := db.C("Folder").FindId(r.Id).One(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}
func (r Folder) GetAllFolders() ([]Folder, error) {
	var folders []Folder
	err := db.C("Folder").Find(bson.M{}).All(&folders)
	if err != nil {
		return nil, err
	}
	return folders, err
}

func (r Folder) Insert()  error {
	r.Id = bson.NewObjectId()
	err := db.C("Folder").Insert(&r)
	if err!=nil{
		return err
	}
	return nil
}

func (r Folder) Update() error {
	err := db.C("Folder").Update(bson.M{"_id": r.Id}, &r)
	if err!=nil {
		return err
	}
	return nil
}
func (r Folder) Delete() error {
	err := db.C("Folder").Remove(&r)
	if err!=nil {
		return err
	}
	return nil
}