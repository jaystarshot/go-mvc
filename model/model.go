package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Game struct {
	ID          bson.ObjectId `bson:"_id"`
	Name        string        `bson:"name"`
	Platform    string        `bson:"platform"`
	Genre       string        `bson:"genre"`
	ReleaseDate string        `bson:"dateTime"`
	NoOfPlayers int32         `bson:"noOfPlayers"`
	Publisher   string        `bson:"publisher"`
	BoxArt      string        `bson:"BoxArt"`
}
