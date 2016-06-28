package rethinkdb

import (
	"github.com/bramvdbogaerde/dbmodelgo/generic"
	r "gopkg.in/dancannon/gorethink.v2"
)

type RethinkDB struct {
	DB      string
	Session *r.Session
}

type RethinkResult struct {
	Table   string
	Query   interface{}
	Session *r.Session
}

func Open(db string, uri string) {
	rt := RethinkDB{}
	rt.DB = db
	rt.Open(uri)
}

func (a *RethinkDB) Open(uri string) {
	session, _ := r.Connect(r.ConnectOpts{
		Address:  uri,
		Database: a.DB,
	})
	a.Session = session
}

func (a *RethinkDB) Filter(query interface{}) RethinkResult {
	q := query.(generic.Query)

	return RethinkResult{
		Table:   q.Table,
		Query:   q.Query,
		Session: a.Session,
	}
}

func (res RethinkResult) All(receiver interface{}) {
	result, _ := r.Table(res.Table).Filter(res.Query).Run(res.Session)
	result.All(receiver)
}
