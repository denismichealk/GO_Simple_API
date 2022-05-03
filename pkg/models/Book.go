package models

import "gopkg.in/mgo.v2/bson"

type Book struct {
	Object_id bson.ObjectId `bson:"_id"`
	Id        int           `bson:"id"`
	Title     string        `bson:"title"`
	Author    string        `bson:"author"`
	Desc      string        `bson:"desc"`
}
