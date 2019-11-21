package main

import (
	"conn"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
)

var db *sql.DB

type User struct {
	ID       int
	Email    string
	Password string
}

type JWT struct {
	Token string
}

type Error struct {
	Message string
}

//設定資料庫資訊
var user = conn.MySqlUser{
	Host:     "127.0.0.1", //主機
	MaxIdle:  10,          //閒置的連接數
	MaxOpen:  10,          //最大連接數
	User:     "root",      //用戶名
	Password: "asdf4440",  //密碼
	Database: "user",      //資料庫名稱
	Port:     3306,        //端口
}

func init() {
	//建立初始化連線
	db = user.Init()
}

func main() {
	//create router
	//func NewRouter() *Router
	router := mux.NewRouter() //*Router

	//func (r *Router) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route
	//func (r *Router) Methods(methods ...string) *Route
	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	//被保護端點
	router.HandleFunc("/protected", TokenVerifyMiddleWare(protectedEndpoint)).Methods("GET")

	//connect server
	//log.Fatal record error
	log.Fatal(http.ListenAndServe(":8080", router))
}

//send error
func respondWithError(w http.ResponseWriter, status int, error Error) {
	//send status and request
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(error)
}

//encode
func responseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

//func(http.ResponseWriter, *http.Request)
func signup(w http.ResponseWriter, r *http.Request) {
	var user User
	var error Error

	//decode
	//decode必須pointer
	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		//respond error
		error.Message = "Email is missing."
		respondWithError(w, http.StatusBadRequest, error)
		return
	}
	if user.Password == "" {
		//respond error
		error.Message = "Password is missing."
		//send status and request
		respondWithError(w, http.StatusBadRequest, error)
		return
	}
	//密碼加密
	//func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	//convert []byte to string
	user.Password = string(hash)

	//execute sql
	_, err = db.Exec("insert into users (id, email, password) values(?, ?, ?)", user.ID, user.Email, user.Password)
	if err != nil {
		error.Message = "Server error"
		respondWithError(w, http.StatusInternalServerError, error)
		return
	}

	//if no error
	//將密碼設為空白
	user.Password = ""
	//set header
	w.Header().Set("Content-Type", "application/json")
	//encode
	responseJSON(w, user)
}

//json-web-token
func GenerateToken(user User) (string, error) {
	var err error
	secret := "secret"

	//a jwt
	//header.payload.secret
	//func NewWithClaims(method SigningMethod, claims Claims) *Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})

	//生成簽名字串(secret)
	//func (t *Token) SignedString(key interface{}) (string, error)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString, nil
}

//func(http.ResponseWriter, *http.Request)
func login(w http.ResponseWriter, r *http.Request) {
	var user User
	var jwt JWT
	var error Error

	//decode
	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		//respond error
		error.Message = "Email is missing."
		respondWithError(w, http.StatusBadRequest, error)
		return
	}
	if user.Password == "" {
		//respond error
		error.Message = "Password is missing."
		//send status and request
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	password := user.Password

	//尋找是否有符合的email
	row := db.QueryRow("select * from users where email=?", user.Email)
	//印在user資料上
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		//假設錯誤為無此資料
		if err == sql.ErrNoRows {
			error.Message = "The user does not exist."
			respondWithError(w, http.StatusBadRequest, error)
			return
		} else {
			log.Fatal(err)
		}
	}

	hashedPassword := user.Password

	//比較密碼是否符合
	//func CompareHashAndPassword(hashedPassword, password []byte) error
	//亂碼的密碼與純文本密碼比較
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	//假設密碼不符
	if err != nil {
		error.Message = "Invaild Password"
		respondWithError(w, http.StatusUnauthorized, error)
		return
	}

	//create token
	token, err := GenerateToken(user)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	jwt.Token = token

	responseJSON(w, jwt)
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protectedEndpoint invoked.")
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	//匿名函式
	return func(w http.ResponseWriter, r *http.Request) {
		var errorObject Error
		authHeader := r.Header.Get("Authorization")
		//split string(slice)
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			//保留第二個值
			//token值
			authHeader := bearerToken[1]

			//jwt解析並驗證
			//func Parse(tokenString string, keyFunc Keyfunc) (*Token, error)
			token, error := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
				/*
									type Token struct {
					    Raw       string                 // The raw token.  Populated when you Parse a token
					    Method    SigningMethod          // The signing method used or to be used
					    Header    map[string]interface{} // The first segment of the token
					    Claims    Claims                 // The second segment of the token
					    Signature string                 // The third segment of the token.  Populated when you Parse a token
					    Valid     bool                   // Is the token valid?  Populated when you Parse/Verify a token
					}*/
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error.")
				}
				return []byte("secret"), nil
			})
			if error != nil {
				errorObject.Message = error.Error()
				respondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}

			//假設token有效，返回true
			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				errorObject.Message = error.Error()
				respondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}
		} else {
			errorObject.Message = "Invaild token."
			respondWithError(w, http.StatusUnauthorized, errorObject)
			return
		}
	}
}
