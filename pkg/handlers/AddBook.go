package handlers

import (
	"GO_Rest_API/pkg/MongoDA"
	"GO_Rest_API/pkg/models"
	"encoding/json"
	"github.com/gofiber/fiber"
	"gopkg.in/mgo.v2/bson"
)

func AddBook(c *fiber.Ctx) {
	session := MongoDA.GetMongoSession()

	var bk models.Book
	json.Unmarshal([]byte(c.Body()), &bk)

	var res = session.DB(dbName).C(collectionName)
	defer session.Close()

	bk.Object_id = bson.NewObjectId()
	err := res.Insert(bk)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	response := models.Response{}

	response.StatusCode = "0"
	response.StatusDesc = "SUCCESS"
	//response, _ := json.Marshal("Inserted")
	err = c.JSON(response)

	if err != nil {
		c.Status(500).Send(err)
		return
	}
}
