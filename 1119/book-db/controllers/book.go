package controllers

import (
	"book-db/model"
	bookRepository "book-db/repository/book"
	"book-db/utils"
	"database/sql"
	"net/http"
)

type Controller struct{}

var books []model.Book

//get all datas
//需要回傳http.HandlerFunc
//需要在HandleFunc使用
//type HandlerFunc func(ResponseWriter, *Request)
func (c Controller) Getbooks(db *sql.DB) http.HandlerFunc {
	//匿名函式
	return func(w http.ResponseWriter, r *http.Request) {
		var book model.Book
		var error model.Error
		bookRepo := bookRepository.BookRepository{}
		books, err := bookRepo.GetBooks(db, book, books)

		if err != nil {
			error.Message = "Serve error"
			//encode
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		//set header
		w.Header().Set("Content-Type", "application/json")
		//encode
		utils.SendSuccess(w, books)

	}
}
