package mongodb

import (
	"gopkg.in/mgo.v2/bson"
)

// Mongodb User database structs
type UserTemporaryInformation struct {
	Id          	 bson.ObjectId `bson:"_id" json:"id" `
	Email 		 	 string        `bson:"email" json:"email"`
	RegisterVerificationCode 	     string           `bson:"registerverificationcode" json:"registerverificationcode"`
	ForgotPasswordVerificationCode 	 string           `bson:"forgotpasswordverificationcode" json:"forgotpasswordverificationcode"`
	RegisterVerificationCodeCreateDate 	     string       `bson:"registerverificationcodecreatedate" json:"registerverificationcodecreatedate"`
	ForgotPasswordVerificationCodeCreateDate 	 string       `bson:"forgotpasswordverificationcodecreatedate" json:"forgotpasswordverificationcodecreatedate"`
	IsCodeUsed 	 	 bool          `bson:"iscodeused" json:"iscodeused"`
	IsCodeExpired 	 bool          `bson:"iscodeexpired" json:"iscodeexpired"`
}


func (r UserTemporaryInformation) CheckRegisterVerificationCode() (*UserTemporaryInformation, error) {
	err := db.C("UserTemporaryInformation").Find(bson.M{"email":r.Email,"registerverificationcode": r.RegisterVerificationCode}).One(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}
func (r UserTemporaryInformation) CheckForgotPasswordVerificationCode() (*UserTemporaryInformation, error) {
	err := db.C("UserTemporaryInformation").Find(bson.M{"email":r.Email,"forgotpasswordverificationcode": r.ForgotPasswordVerificationCode}).One(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}
func (r UserTemporaryInformation) Insert()  error {
	r.Id = bson.NewObjectId()
	err := db.C("UserTemporaryInformation").Insert(&r)
	if err!=nil{
		return err
	}
	return nil
}
func (r UserTemporaryInformation) Update() error {
	err := db.C("UserTemporaryInformation").Update(bson.M{"_id": r.Id}, &r)
	if err!=nil {
		return err
	}
	return nil
}
func (r UserTemporaryInformation) Delete() error {
	err := db.C("UserTemporaryInformation").Remove(&r)
	if err!=nil {
		return err
	}
	return nil
}
func (r UserTemporaryInformation) GetAllUserTemporaryInformation() ([]UserTemporaryInformation, error) {
	var informations []UserTemporaryInformation
	err := db.C("UserTemporaryInformation").Find(bson.M{}).All(&informations)
	if err != nil {
		return nil, err
	}
	return informations, err
}


