package handlers

import (
	"GO_Rest_API/pkg/MongoDA"
	"GO_Rest_API/pkg/models"
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

	//json, _ := json.Marshal(results)
	//c.Send(json)
	//response := models.Response{}
	//response.StatusCode = "0"
	//response.StatusDesc = "SUCCESS"
	////response, _ := json.Marshal("Inserted")
	//c.Set("Content-type", "application/json; charset=utf-8")
	//c.Send(response)
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
