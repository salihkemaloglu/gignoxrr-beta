package repositories

import (
	"gopkg.in/mgo.v2/bson"
)

// User ...
type User struct {
	ID               bson.ObjectId `bson:"_id" json:"id" `
	Name             string        `bson:"name" json:"name"`
	Surname          string        `bson:"surname" json:"surname"`
	Email            string        `bson:"email" json:"email"`
	Username         string        `bson:"username" json:"username"`
	Password         string        `bson:"password" json:"password"`
	Description      string        `bson:"description" json:"description"`
	CreatedDate      string        `bson:"createddate" json:"createddate"`
	UpdatedDate      string        `bson:"updateddate" json:"updateddate"`
	ImagePath        string        `bson:"imagepath" json:"imagepath"`
	TotalSpace       int32         `bson:"totalspace" json:"totalspace"`
	LanguageCode     string        `bson:"languagecode" json:"languagecode"`
	IsAccountConfirm bool          `bson:"isuserverificated" json:"isuserverificated"`
}

//Login ...
// Crud operaions for User
func (r User) Login() (*User, error) {
	err := db.C("User").Find(bson.M{"$or": []bson.M{{"username": r.Username, "password": r.Password}, {"email": r.Username, "password": r.Password}}}).One(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}

//GetUser ...
func (r User) GetUser() (*User, error) {
	err := db.C("User").FindId(r.ID).One(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}

//GetUserByEmail ...
func (r User) GetUserByEmail() (*User, error) {
	err := db.C("User").Find(bson.M{"email": r.Email}).One(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}

//GetUserByUsername ...
func (r User) GetUserByUsername() (*User, error) {
	err := db.C("User").Find(bson.M{"username": r.Username}).One(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}

//Insert ...
func (r User) Insert() error {
	r.ID = bson.NewObjectId()
	err := db.C("User").Insert(&r)
	if err != nil {
		return err
	}
	return nil
}

//Update ...
func (r User) Update() error {
	err := db.C("User").Update(bson.M{"_id": r.ID}, &r)
	if err != nil {
		return err
	}
	return nil
}

//UpdateUserPassword ...
func (r User) UpdateUserPassword() error {
	err := db.C("User").Update(bson.M{"_id": r.ID}, bson.M{"$set": bson.M{"password": r.Password}})
	if err != nil {
		return err
	}
	return nil
}

//Delete ...
func (r User) Delete() error {
	err := db.C("User").Remove(&r)
	if err != nil {
		return err
	}
	return nil
}

//CheckUser user login
func (r User) CheckUser() error {
	err := db.C("User").Find(bson.M{"$or": []bson.M{{"username": r.Username}, {"email": r.Email}}}).One(&r)
	if err != nil {
		return err
	}
	return nil
}
