package handlers

import (
	"GO_Rest_API/pkg/models"
	"github.com/gofiber/fiber"
)

func GetHealth(c *fiber.Ctx) {

	var free_book models.Book
	free_book.Id = 1
	free_book.Desc = "This is a story about a mustard seed's journey of growing into something big and magnificent."
	free_book.Author = "Deborah A .Smith"
	free_book.Title = "Faith Of a Mustard Seed"

	err := c.JSON(free_book)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

}
