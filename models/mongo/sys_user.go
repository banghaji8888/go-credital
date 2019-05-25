package mongo

import (
	"gopkg.in/mgo.v2/bson"
)

// SysUser ->>
type SysUser struct {
	ID       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
	Group    string        `bson:"group"`
}
