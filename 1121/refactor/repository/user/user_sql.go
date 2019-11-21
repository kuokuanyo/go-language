package userRepository

import (
	"database/sql"
	"log"
	"refactor/models"
)

type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (u UserRepository) Signup(db *sql.DB, user models.User) models.User {
	//execute sql
	_, err := db.Exec("insert into users (id, email, password) values(?, ?, ?)", user.ID, user.Email, user.Password)
	logFatal(err)

	//if no error
	//將密碼設為空白
	user.Password = ""
	return user
}

func (u UserRepository) Login(db *sql.DB, user models.User) (models.User, error) {
	//尋找是否有符合的email
	row := db.QueryRow("select * from users where email=?", user.Email)
	//印在user資料上
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}
