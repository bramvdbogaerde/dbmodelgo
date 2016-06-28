package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoResult struct {
	Query   bson.M
	Session *mgo.Session
	Table   string
}

func (r MongoResult) All(receiver interface{}) {
	r.Session.DB("test").C(r.Table).Find(r.Query).All(receiver)
}

func (r MongoResult) One(receiver interface{}) {
	r.Session.DB("test").C(r.Table).Find(r.Query).One(receiver)
}
