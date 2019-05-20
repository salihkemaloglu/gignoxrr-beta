package repositories

import (
	"gopkg.in/mgo.v2/bson"
)
type UserTemporaryInformation struct {
	Id          	 bson.ObjectId `bson:"_id" json:"id" `
	Email 		 	 string        `bson:"email" json:"email"`
	RegisterVerificationToken	     string           `bson:"registerverificationtoken" json:"registerverificationtoken"`
	ForgotPasswordVerificationToken 	 string           `bson:"forgotpasswordverificationtoken" json:"forgotpasswordverificationtoken"`
	RegisterVerificationTokenCreateDate 	     string       `bson:"registerverificationtokencreatedate" json:"registerverificationtokencreatedate"`
	ForgotPasswordVerificationTokenCreateDate 	 string       `bson:"forgotpasswordverificationtokencreatedate" json:"forgotpasswordverificationtokencreatedate"`
	EmailType 	 	 string        `bson:"emailtype" json:"emailtype"`
	IsTokenUsed 	 	 bool          `bson:"istokenused" json:"istokenused"`
	IsTokenExpired 	 bool          `bson:"istokenexpired" json:"istokenexpired"`
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
func (r UserTemporaryInformation) UpdateByEmail() error {
	err := db.C("UserTemporaryInformation").Update(bson.M{"email": r.Email,"forgotpasswordverificationtoken": r.ForgotPasswordVerificationToken}, bson.M{"$set": bson.M{"istokenused": r.IsTokenUsed}})
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
func (r UserTemporaryInformation) CheckRegisterVerificationToken() (*UserTemporaryInformation, error) {
	err := db.C("UserTemporaryInformation").Find(bson.M{"emailtype":r.EmailType,"registerverificationtoken": r.RegisterVerificationToken}).One(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}
func (r UserTemporaryInformation) CheckForgotPasswordVerificationToken() (*UserTemporaryInformation, error) {
	err := db.C("UserTemporaryInformation").Find(bson.M{"emailtype":r.EmailType,"forgotpasswordverificationtoken": r.ForgotPasswordVerificationToken}).One(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}
func (r UserTemporaryInformation) CheckVerificationTokenResentEmail() (*UserTemporaryInformation, error) {
	err := db.C("UserTemporaryInformation").Find(bson.M{"email":r.Email,"emailtype":r.EmailType,"istokenused":false,"istokenexpired": false}).One(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}
func (r UserTemporaryInformation) GetAllUserTemporaryInformation() ([]UserTemporaryInformation, error) {
	var informations []UserTemporaryInformation
	err := db.C("UserTemporaryInformation").Find(bson.M{}).All(&informations)
	if err != nil {
		return nil, err
	}
	return informations, err
}


