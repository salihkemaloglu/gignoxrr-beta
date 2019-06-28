package repositories

import (
	"gopkg.in/mgo.v2/bson"
)

//File ...
type File struct {
	ID           bson.ObjectId `bson:"_id" json:"id" `
	UserID       string        `bson:"userid" json:"userid"`
	FolderID     string        `bson:"folderid" json:"folderid"`
	Name         string        `bson:"name" json:"name"`
	Description  string        `bson:"description" json:"description"`
	CreatedDate  string        `bson:"createddate" json:"createddate"`
	UpdatedDate  string        `bson:"updateddate" json:"updateddate"`
	FileHash     string        `bson:"filehash" json:"filehash"`
	IsBuried     bool          `bson:"isburied" json:"isburied"`
	IsFolderFile bool          `bson:"isfolderfile" json:"isfolderfile"`
	IsStarred    bool          `bson:"isstarred" json:"isstarred"`
	IsTrash      bool          `bson:"istrash" json:"istrash"`
	IsDeleted    bool          `bson:"isdeleted" json:"isdeleted"`
}

// GetFile ...
// Crud operaions for File
func (r File) GetFile() (*File, error) {
	err := db.C("File").FindId(r.ID).One(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}

//GetAllFiles ...
func (r File) GetAllFiles() ([]File, error) {
	var files []File
	err := db.C("File").Find(bson.M{}).All(&files)
	if err != nil {
		return nil, err
	}
	return files, err
}

//Insert ...
func (r File) Insert() error {
	r.ID = bson.NewObjectId()
	err := db.C("File").Insert(&r)
	if err != nil {
		return err
	}
	return nil
}

//Update ...
func (r File) Update() error {
	err := db.C("File").Update(bson.M{"_id": r.ID}, &r)
	if err != nil {
		return err
	}
	return nil
}

//Delete ...
func (r File) Delete() error {
	err := db.C("File").Remove(&r)
	if err != nil {
		return err
	}
	return nil
}
