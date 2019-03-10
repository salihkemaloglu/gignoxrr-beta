package mongodb

import (
	"gopkg.in/mgo.v2/bson"
)

// Mongodb User database structs
type User struct {
	ID          	 bson.ObjectId `bson:"_id" json:"id" `
	Name 	 		 string        `bson:"name" json:"name"`
	Surname 		 string        `bson:"surname" json:"surname"`
	Email 	    	 string        `bson:"email" json:"email"`
	Username 		 string        `bson:"username" json:"username"`
	Password 		 string        `bson:"password" json:"password"`
	Description 	 string        `bson:"description" json:"description"`
	CreatedDate  	 string        `bson:"createddate" json:"createddate"`
	UpdatedDate  	 string        `bson:"updateddate" json:"updateddate"`
	ImagePath 		 string        `bson:"imagepath" json:"imagepath"`
	TotalSpace  	 string        `bson:"totalspace" json:"totalspace"`
	LanguageType 	 string        `bson:"languagetype" json:"languagetype"`
}

// Crud operaions for User
func (r User) Login() (bool, error) {
	err := db.C("user").Find(bson.M{"username": r.Username, "password": r.Password}).One(&r)
	if err != nil {
		return false,err
	}
	return true,nil
}
func (r User) GetUser() (*User, error) {
	err := db.C("user").FindId(r.ID).One(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}
func (r User) Insert()  error {
	r.ID = bson.NewObjectId()
	err := db.C("User").Insert(&r)
	if err!=nil{
		return err
	}
	return nil
}

func (r User) Update() error {
	err := db.C("user").Update(bson.M{"_id": r.ID}, &r)
	if err!=nil {
		return err
	}
	return nil
}
func (r User) Delete() error {
	err := db.C("user").Remove(&r)
	if err!=nil {
		return err
	}
	return nil
}