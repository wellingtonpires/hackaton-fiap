package cadastrausuario

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
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

type Usuario struct {
	Login string `json:"login"`
	Senha string `json:"senha"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
	Admin bool   `json:"admin"`
}

func CadastraUsuario(c *gin.Context) {
	var u Usuario

	err := c.ShouldBindJSON(&u)
	if err != nil {
		fmt.Println(err.Error())
	}

	con, err := OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Conectado DB")
	}
	defer con.Close()

	rows, err := con.Query(`INSERT INTO usuario VALUES ($1, $2, $3, $4, $5)`, u.Login, u.Senha, u.Email, u.Cpf, u.Admin)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()
}
