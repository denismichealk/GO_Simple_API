package models

type Response struct {
	StatusCode string `json:"code"`
	StatusDesc string `json:"message"`
}
