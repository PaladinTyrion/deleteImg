package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Trick struct {
	Id              bson.ObjectId `bson:"_id"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Slug            string
	Screenshot      string
	User            bson.ObjectId
	Title           string
	Origin_url      string
	Description     string
	Is_active       bool
	Favorite_counts int32
	Click_counts    int32
	View_counts     int32
	Tags            []string
	v               int32 `bson:"__v"`
}

func CreateNewTrick() *Trick {
	return new(Trick)
}
