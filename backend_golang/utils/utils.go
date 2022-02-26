package utils

import (
	"backend_golang/constants"
	"backend_golang/router/sqldb"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//SetHeaders sets the cors headers
func SetHeaders(res *http.ResponseWriter) {
	(*res).Header().Set("Access-Control-Allow-Origin", "*")
	(*res).Header().Set("Access-Control-Allow-Methods", "*")
	(*res).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
}

//GenerateRandomString generates a random string of length 10
func GenerateRandomString() string {
	length := 10
	characters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charactersLength := len(characters)
	randomString := ""

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		randomString += string(characters[rand.Intn(charactersLength-1)])
	}

	return randomString
}

//GenerateToken generates a JWT Token
func GenerateToken(email string) string {
	temp, _ := os.LookupEnv("SECRETCODE")
	key := []byte(temp)

	randstring := GenerateRandomString()

	expirationTime := time.Now().Add(60 * time.Minute)

	claims := struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		jwt.StandardClaims
	}{
		Username: randstring,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Fatal(err.Error())
	}

	return tokenString
}

//VerifyToken verifies the token and return boolean
func VerifyToken(token string) string {
	temp, _ := os.LookupEnv("SECRETCODE")
	key := []byte(temp)

	claims := &struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		jwt.StandardClaims
	}{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return ""
	}
	if !tkn.Valid {
		return ""
	}

	return (*claims).Email
}

//IsValidUser returns whether user is valid or not
func IsValidUser(token string) string {

	email := VerifyToken(token)
	stmt := fmt.Sprintf("select * from %s.user where email = '%s' and token='%s", constants.MYSQL_DATABASE, email, token)
	rows, _ := sqldb.DB.Query(stmt)
	userExists := rows.Next()
	if userExists {
		return email
	}

	return ""
}
