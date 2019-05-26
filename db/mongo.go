package db

import (
	"go-credital/config"
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session
var err error

//InitMongo - init mongo connection
func InitMongo() {
	conf := config.GetConfig()

	session, err = mgo.DialWithTimeout(conf.MongoHost, 1*time.Hour)

	if err != nil {
		log.Fatal(err)
	} else {
		session.SetMode(mgo.Monotonic, true)
		session.SetSyncTimeout(1 * time.Hour)
		session.SetSocketTimeout(1. * time.Hour)
	}
}

// GetConnection - get current connection
func GetConnection() *mgo.Session {
	return session
}
