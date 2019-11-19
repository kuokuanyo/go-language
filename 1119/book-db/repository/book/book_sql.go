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
