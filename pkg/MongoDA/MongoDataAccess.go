package MongoDA

import (
	"gopkg.in/mgo.v2"
	"log"
)

//GetMongoDbConnection get connection of mongodb
func GetMongoSession() *mgo.Session {

	mgoSession, err := mgo.Dial("localhost")
	mgoSession.SetMode(mgo.Monotonic, true)

	if err != nil {
		log.Fatal("Failed to start the Mongo session.")
	}

	return mgoSession.Clone()
}
