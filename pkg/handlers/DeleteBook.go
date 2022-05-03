package handlers

import (
	"GO_Rest_API/pkg/MongoDA"
	"GO_Rest_API/pkg/models"
	"github.com/gofiber/fiber"
	"gopkg.in/mgo.v2/bson"
)

func DeleteBook(c *fiber.Ctx) {
	session := MongoDA.GetMongoSession()

	var res = session.DB(dbName).C(collectionName)
	defer session.Close()

	id := c.Query("id")

	ob := bson.ObjectIdHex(id)
	err := res.RemoveId(ob)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	//jsonResponse, _ := json.Marshal("DELETED")
	//c.Set("Content-type", "application/json; charset=utf-8")
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
