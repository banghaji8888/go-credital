package models

import "gopkg.in/mgo.v2/bson"

// SysMenu =>>
type SysMenu struct {
	ID     bson.ObjectId `bson:"_id"`
	Parent string        `bson:"parent"`
	Name   string        `bson:"name"`
	URL    string        `bson:"url"`
	Urutan int8          `bson:"urutan"`
	Status int8          `bson:"status"`
	icon   string        `bson:"icon"`
}
