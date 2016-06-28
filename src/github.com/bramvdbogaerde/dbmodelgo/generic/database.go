package generic

type Database interface {
	Filter(query interface{}) *Result
	Open(uri string)
}

type Result interface {
	All(receiver interface{})
	One(receiver interface{})
}

type Query struct {
	Table string
	Query interface{}
}

type H map[string]interface{}
