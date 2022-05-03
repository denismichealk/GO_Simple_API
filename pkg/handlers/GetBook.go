package handlers

import (
	"GO_Rest_API/pkg/MongoDA"
	"GO_Rest_API/pkg/models"
	"encoding/json"
	"github.com/gofiber/fiber"
	"gopkg.in/mgo.v2/bson"
)

func GetBook(c *fiber.Ctx) {
	session := MongoDA.GetMongoSession()

	results := models.Book{}
	var res = session.DB(dbName).C(collectionName)
	defer session.Close()

	//id := c.Params("id")

	id := c.Query("id")

	ob := bson.ObjectIdHex(id)
	err := res.FindId(ob).One(&results)

	if err != nil {
		c.SendStatus(404)
		return
	}

	json, _ := json.Marshal(results)
	c.Send(json)

}
