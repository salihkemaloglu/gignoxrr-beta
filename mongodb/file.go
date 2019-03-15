
package mongodb

import (
	"gopkg.in/mgo.v2/bson"
)

type File struct {
	Id          	 bson.ObjectId `bson:"_id" json:"id" `
	UserId 	    	 string        `bson:"userid" json:"userid"`
	FolderId 		 string        `bson:"folderid" json:"folderid"`
	Name 	 		 string        `bson:"name" json:"name"`
	Description 	 string        `bson:"description" json:"description"`
	CreatedDate 	 string        `bson:"createddate" json:"createddate"`
	UpdatedDate 	 string        `bson:"updateddate" json:"updateddate"`
	FileHash    	 string        `bson:"filehash" json:"filehash"`
	IsBuried    	 bool          `bson:"isburied" json:"isburied"`
	IsFolderFile	 bool          `bson:"isfolderfile" json:"isfolderfile"`
	IsStarred   	 bool          `bson:"isstarred" json:"isstarred"`
	IsTrash     	 bool          `bson:"istrash" json:"istrash"`
	IsDeleted   	 bool          `bson:"isdeleted" json:"isdeleted"`
}
// Crud operaions for File
func (r File) GetFile() (*File, error) {
	err := db.C("File").FindId(r.Id).One(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}
func (r File) GetAllFiles() ([]File, error) {
	var files []File
	err := db.C("File").Find(bson.M{}).All(&files)
	if err != nil {
		return nil, err
	}
	return files, err
}

func (r File) Insert()  error {
	r.Id = bson.NewObjectId()
	err := db.C("File").Insert(&r)
	if err!=nil{
		return err
	}
	return nil
}

func (r File) Update() error {
	err := db.C("File").Update(bson.M{"_id": r.Id}, &r)
	if err!=nil {
		return err
	}
	return nil
}
func (r File) Delete() error {
	err := db.C("File").Remove(&r)
	if err!=nil {
		return err
	}
	return nil
}