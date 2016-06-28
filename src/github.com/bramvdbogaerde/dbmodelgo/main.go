package main

import (
	"fmt"
	"github.com/bramvdbogaerde/dbmodelgo/generic"
	"github.com/bramvdbogaerde/dbmodelgo/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Friend struct {
	Name string
}

func main() {
	mDb := mongo.Open("localhost")

	var results []Friend

	mDb.Filter(generic.Query{
		Table: "friends",
		Query: bson.M{},
	}).All(&results)

	fmt.Println(results)
}
