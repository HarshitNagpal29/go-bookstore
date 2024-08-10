package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/HarshitNagpal29/go-bookstore/pkg/models"
	"github.com/HarshitNagpal29/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "application-json")
	w.WriteHeader(http.StatusOK)
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Write(res)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "application-json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	bookId := params["bookid"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "application-json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	bookId := params["bookid"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails := models.DeleteBook(ID)
	res, _ := json.Marshal(bookDetails)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "application-json")
	w.WriteHeader(http.StatusOK)
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	params := mux.Vars(r)
	bookId := params["bookid"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Write(res)

}
