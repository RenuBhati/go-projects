package controllers

import (
	"encoding/json"
	"net/http"

	"example.com/hello/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(NewBook)
}
