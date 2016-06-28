package mongo

import (
	"github.com/bramvdbogaerde/dbmodelgo/generic"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDatabase struct {
	Session *mgo.Session
}

func (m *MongoDatabase) Open(uri string) {
	m.Session, _ = mgo.Dial(uri)
}

func (m *MongoDatabase) Filter(query interface{}) generic.Result {
	q := query.(generic.Query)

	return MongoResult{
		Table:   q.Table,
		Query:   q.Query.(bson.M),
		Session: m.Session,
	}
}
