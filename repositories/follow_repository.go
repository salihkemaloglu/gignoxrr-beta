package repositories

import (
	"gopkg.in/mgo.v2/bson"
)

//Follow ...
type Follow struct {
	ID             bson.ObjectId `bson:"_id" json:"id" `
	FollowerID     string        `bson:"followerid" json:"followerid"`
	FollowedID     string        `bson:"followedid" json:"followedid"`
	FollowedDate   string        `bson:"followedate" json:"followedate"`
	UnfollowedDate string        `bson:"unfollowedate" json:"unfollowedate"`
}
