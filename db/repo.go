package db

import (
	"go-credital/config"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoRepo - interface repo
type MongoRepo interface {
	Save(interface{}) error
	Update(interface{}) error
	Delete(interface{}) error
	FindOne(interface{}, interface{}) (interface{}, error)
	FindMany(interface{}) ([]bson.M, error)
	FindMenu(interface{}) ([]bson.M, error)
}

type mongoRepo struct {
	session    *mgo.Session
	collection string
	dbname     string
}

//InitMongoRepo - init repo
func InitMongoRepo(collection string) *mongoRepo {
	conf := config.GetConfig()
	return &mongoRepo{
		session:    GetConnection(),
		collection: collection,
		dbname:     conf.MongoDBName,
	}
}

func (r *mongoRepo) Save(flight interface{}) error {
	sessionCopy := r.session.Copy()
	defer sessionCopy.Close()

	err := sessionCopy.DB(r.dbname).C(r.collection).Insert(flight)
	return err
}

func (r *mongoRepo) Update(where interface{}, data interface{}) error {
	sessionCopy := r.session.Copy()
	defer sessionCopy.Close()

	err := sessionCopy.DB(r.dbname).C(r.collection).Update(where, data)
	return err
}

func (r *mongoRepo) Delete(data interface{}) error {
	sessionCopy := r.session.Copy()
	defer sessionCopy.Close()

	_, err := sessionCopy.DB(r.dbname).C(r.collection).RemoveAll(data)

	return err
}

func (r *mongoRepo) FindOne(data interface{}, out interface{}) (interface{}, error) {
	sessionCopy := r.session.Copy()
	defer sessionCopy.Close()

	err := sessionCopy.DB(r.dbname).C(r.collection).Find(data).One(&out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *mongoRepo) FindMany(data interface{}) ([]bson.M, error) {
	sessionCopy := r.session.Copy()
	defer sessionCopy.Close()

	var out []bson.M

	err := sessionCopy.DB(r.dbname).C(r.collection).Find(data).All(&out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *mongoRepo) FindMenu(data interface{}) ([]bson.M, error) {
	sessionCopy := r.session.Copy()
	defer sessionCopy.Close()

	var out []bson.M

	err := sessionCopy.DB(r.dbname).C(r.collection).Find(data).Sort("parent", "urutan").All(&out)

	if err != nil {
		return nil, err
	}

	return out, nil
}
