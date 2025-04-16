package validatoken

import (
	"database/sql"
	"fmt"

	"github.com/golang-jwt/jwt"
)

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://postgres:123@postgres:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Conectado ao banco!")
	}

	err = db.Ping()

	return db, err
}

var secretKey = []byte("hackaton")

func ValidaToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false
	}

	if !token.Valid {
		return false
	} else {
		return true
	}
}

func ValidaTokenAdmin(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false
	}

	if !token.Valid {
		return false
	} else {
		claims := token.Claims.(jwt.MapClaims)
		fmt.Print(claims["role"])
		if claims["role"] == "admin" {
			return true
		} else {
			return false
		}
	}
}
