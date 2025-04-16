package criatoken

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://postgres:123@postgres:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	return db, err
}

type Usuario struct {
	Login string `json:"login"`
	Senha string `json:"senha"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
	Admin bool   `json:"admin"`
}

func verificaCadastro(login string, senha string) (status bool) {
	existe := false
	var u Usuario
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer con.Close()
	rows, err := con.Query(`SELECT * FROM usuario`)
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		rows.Scan(&u.Login, &u.Senha, &u.Email, &u.Cpf, &u.Admin)
	}
	if u.Login == login && u.Senha == senha {
		existe = true
	}
	defer rows.Close()
	return existe
}

func consultaRole(login string, senha string) (role string, cpf string) {
	role = "default"
	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer con.Close()
	rows, err := con.Query(`SELECT * FROM usuario WHERE login = $1 AND password = $2`, login, senha)
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var u Usuario
		rows.Scan(&u.Login, &u.Senha, &u.Email, &u.Cpf, &u.Admin)
		cpf = u.Cpf
		if u.Admin {
			role = "admin"
		}
	}
	defer rows.Close()
	return role, cpf
}

var secretKey = []byte("hackatons")

func CriaToken(c *gin.Context) {
	role, cpf := consultaRole(c.Query("login"), c.Query("senha"))
	if verificaCadastro(c.Query("login"), c.Query("senha")) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256,
			jwt.MapClaims{
				"username": c.Query("login"),
				"role":     role,
				"exp":      time.Now().Add(time.Hour * 24).Unix(),
				"cpf":      cpf,
			})
		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao assinar token"})
		} else {
			c.IndentedJSON(http.StatusCreated, gin.H{"token": tokenString})
		}
	} else {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"erro": "Usuário não existe na base"})
	}
}
