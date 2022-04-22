package handlers

import (
	"GO_Rest_API/pkg/models"
	"encoding/json"
	"github.com/gofiber/fiber"
)

func GetHealth(c *fiber.Ctx) {

	var free_book models.Book
	free_book.Id = 1
	free_book.Desc = "This is a story about a mustard seed's journey of growing into something big and magnificent."
	free_book.Author = "Deborah A .Smith"
	free_book.Title = "Faith Of a Mustard Seed"
	//body := "Hello world"
	//t := &http.Response{
	//	Status:        "200 OK",
	//	StatusCode:    200,
	//	Proto:         "HTTP/1.1",
	//	ProtoMajor:    1,
	//	ProtoMinor:    1,
	//	Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
	//	ContentLength: int64(len(body)),
	//	Header:        make(http.Header, 0),
	//}

	jsonResp, _ := json.Marshal(free_book)
	c.Send(jsonResp)
	return

}
