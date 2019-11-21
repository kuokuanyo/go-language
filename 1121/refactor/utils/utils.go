package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"refactor/models"

	"github.com/dgrijalva/jwt-go"
)

//send error
func RespondWithError(w http.ResponseWriter, status int, error models.Error) {
	//send status and request
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(error)
}

//encode
func ResponseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

//json-web-token
func GenerateToken(user models.User) (string, error) {
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
