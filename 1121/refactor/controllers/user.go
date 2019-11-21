package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"refactor/models"
	userRepository "refactor/repository/user"
	"refactor/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct{}

//func(http.ResponseWriter, *http.Request)
func (c Controller) Signup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var error models.Error

		//decode
		//decode必須pointer
		json.NewDecoder(r.Body).Decode(&user)

		if user.Email == "" {
			//respond error
			error.Message = "Email is missing."
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}
		if user.Password == "" {
			//respond error
			error.Message = "Password is missing."
			//send status and request
			utils.RespondWithError(w, http.StatusBadRequest, error)
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
		userRepo := userRepository.UserRepository{}
		user = userRepo.Signup(db, user)

		//if no error
		//將密碼設為空白
		user.Password = ""
		//set header
		w.Header().Set("Content-Type", "application/json")
		//encode
		utils.ResponseJSON(w, user)
	}
}

//func(http.ResponseWriter, *http.Request)
func (c Controller) Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var jwt models.JWT
		var error models.Error

		//decode
		json.NewDecoder(r.Body).Decode(&user)

		if user.Email == "" {
			//respond error
			error.Message = "Email is missing."
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}
		if user.Password == "" {
			//respond error
			error.Message = "Password is missing."
			//send status and request
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		password := user.Password

		//尋找是否有符合的email
		//execute sql
		userRepo := userRepository.UserRepository{}
		user, err := userRepo.Login(db, user)

		hashedPassword := user.Password

		//比較密碼是否符合
		//func CompareHashAndPassword(hashedPassword, password []byte) error
		//亂碼的密碼與純文本密碼比較
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		//假設密碼不符
		if err != nil {
			error.Message = "Invaild Password"
			utils.RespondWithError(w, http.StatusUnauthorized, error)
			return
		}

		//create token
		token, err := utils.GenerateToken(user)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		jwt.Token = token

		utils.ResponseJSON(w, jwt)
	}
}

func (c Controller) TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	//匿名函式
	return func(w http.ResponseWriter, r *http.Request) {
		var errorObject models.Error
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
				utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}

			//假設token有效，返回true
			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				errorObject.Message = error.Error()
				utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}
		} else {
			errorObject.Message = "Invaild token."
			utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
			return
		}
	}
}
