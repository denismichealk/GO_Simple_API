package handlers

import (
	"GO_Rest_API/pkg/MongoDA"
	"GO_Rest_API/pkg/models"
	"context"
	"encoding/json"
	"github.com/gofiber/fiber"
)

func AddBook(c *fiber.Ctx) {
	collection, err := MongoDA.GetMongoDbCollection(dbName, collectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var bk models.Book
	json.Unmarshal([]byte(c.Body()), &bk)

	res, err := collection.InsertOne(context.Background(), bk)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	response, _ := json.Marshal(res)
	c.Send(response)
}
