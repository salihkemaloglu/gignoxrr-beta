package mongodb

import (
	"gopkg.in/mgo.v2/bson"
)

type Follow struct {
	Id          	 bson.ObjectId `bson:"_id" json:"id" `
	FollowerId 		 string        `bson:"followerid" json:"followerid"`
	FollowedId 		 string        `bson:"followedid" json:"followedid"`
	FollowedDate 	 string        `bson:"followedate" json:"followedate"`
	UnfollowedDate 	 string        `bson:"unfollowedate" json:"unfollowedate"`
}