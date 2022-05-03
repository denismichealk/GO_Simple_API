package handlers

import (
	"GO_Rest_API/pkg/MongoDA"
	"GO_Rest_API/pkg/models"
	"encoding/json"
	"github.com/gofiber/fiber"
)

func UpdateBook(c *fiber.Ctx) {

	session := MongoDA.GetMongoSession()

	var res = session.DB(dbName).C(collectionName)
	defer session.Close()

	var bk models.Book
	json.Unmarshal([]byte(c.Body()), &bk)

	err := res.UpdateId(bk.Object_id, &bk)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	//response, _ := json.Marshal("UPDATED")
	//c.Send(response)
	response := models.Response{}
	response.StatusCode = "0"
	response.StatusDesc = "SUCCESS"
	//response, _ := json.Marshal("Inserted")
	c.Set("Content-type", "application/json; charset=utf-8")
	c.Send(response)
}
