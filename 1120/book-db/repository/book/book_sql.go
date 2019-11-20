package bookRepository

import (
	"book-db/model"
	"database/sql"
)

type BookRepository struct{}

func (b BookRepository) GetBooks(db *sql.DB, book model.Book, books []model.Book) ([]model.Book, error) {
	//尋找data
	rows, err := db.Query("select * from books")
	if err != nil {
		return []model.Book{}, err
	}
	//最後要關閉
	defer rows.Close()
	//處理每一行
	//Next method 迭代查詢資料，回傳bool
	//func (rs *Rows) Next() bool
	for rows.Next() {
		//Scan method方法用來讀取每一列的值
		//func (rs *Rows) Scan(dest ...interface{}) error
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return []model.Book{}, err
		}
		//add slice
		books = append(books, book)
	}
	return books, nil
}

func (b BookRepository) GetBook(db *sql.DB, book model.Book, id int) (model.Book, error) {
	rows := db.QueryRow("select * from books where id=?", id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	return book, err
}

func (b BookRepository) AddBook(db *sql.DB, book model.Book) (int, error) {
	_, err := db.Exec("insert into books (id, title, author, year) values(?, ?, ?, ?);",
		book.ID, book.Title, book.Author, book.Year)
	if err != nil {
		return 0, err
	}
	return book.ID, nil
}

func (b BookRepository) UpdateBook(db *sql.DB, book model.Book) (int64, error) {
	result, err := db.Exec("update books set title=?, author=?, year=? where id=?",
		book.Title, book.Author, book.Year, book.ID)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (b BookRepository) RemoveBook(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("delete from books where id=?", id)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}
