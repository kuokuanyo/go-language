package controllers

import (
	"book-db/model"
	bookRepository "book-db/repository/book"
	"book-db/utils"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		books = []model.Book{}
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

//get one data
//需要回傳http.HandlerFunc
//需要在HandleFunc使用
//type HandlerFunc func(ResponseWriter, *Request)
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	//匿名函式
	return func(w http.ResponseWriter, r *http.Request) {
		var book model.Book
		var error model.Error

		//return map
		//func Vars(r *http.Request) map[string]string
		params := mux.Vars(r)

		books = []model.Book{}
		bookRepo := bookRepository.BookRepository{}

		//convert string to int
		id, _ := strconv.Atoi(params["id"])

		book, err := bookRepo.GetBook(db, book, id)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Not Found"
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				error.Message = "Server error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}
		//set header
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book model.Book
		var error model.Error
		var bookID int

		//decode(r.Body)
		//必須使用pointer
		json.NewDecoder(r.Body).Decode(&book)

		if book.ID == 0 || book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "Enter missing fields."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		bookID, err := bookRepo.AddBook(db, book)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, bookID)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book model.Book
		var error model.Error

		//decode(req.Body)
		//pointer
		json.NewDecoder(r.Body).Decode(&book)

		if book.ID == 0 || book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "All fields are required."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		rows, err := bookRepo.UpdateBook(db, book)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rows)
	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error model.Error
		//return map
		//func Vars(r *http.Request) map[string]string
		params := mux.Vars(r)
		bookRepo := bookRepository.BookRepository{}
		//convert string to int
		id, _ := strconv.Atoi(params["id"])

		rows, err := bookRepo.RemoveBook(db, id)
		if err != nil {
			error.Message = "Server error."
			utils.SendError(w, http.StatusInternalServerError, error)
		}

		if rows == 0 {
			error.Message = "Not Found"
			utils.SendError(w, http.StatusNotFound, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rows)
	}
}
