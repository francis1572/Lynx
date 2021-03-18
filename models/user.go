package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

//User structure
type User struct {
	Name        string `bson:"name" json:"name"`
	AccessToken string `bson:"accessToken" json:"accessToken"`
	ImageUrl    string `bson:"imageUrl" json:"imageUrl"`
	Email       string `bson:"email" json:"email"`
	FamilyName  string `bson:"familyName" json:"familyName"`
	GivenName   string `bson:"givenName" json:"givenName"`
	UserId      string `bson:"userId" json:"userId"`
}

func (u *User) ToQueryBson() bson.M {
	var queryObject bson.M
	if u.UserId != "" {
		queryObject = bson.M{
			"userId": u.UserId,
		}
	} else {
		queryObject = bson.M{
			"email": u.Email,
		}
	}
	return queryObject
}
