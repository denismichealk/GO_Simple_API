package handlers

import (
	"GO_Rest_API/pkg/MongoDA"
	"GO_Rest_API/pkg/models"
	"github.com/gofiber/fiber"
	"gopkg.in/mgo.v2/bson"
)

func GetAllBooks(c *fiber.Ctx) {

	session := MongoDA.GetMongoSession()

	//results := models.Book{}
	results := make([]models.Book, 0, 10)
	var res = session.DB(dbName).C(collectionName)
	defer session.Close()

	err := res.Find(bson.D{}).All(&results)

	if err != nil {
		c.SendStatus(404)
		return
	}

	//json, _ := json.Marshal(results)
	//c.Send(json)
	response := models.Response{}
	response.StatusCode = "0"
	response.StatusDesc = "SUCCESS"
	//response, _ := json.Marshal("Inserted")
	err = c.JSON(results)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

}
