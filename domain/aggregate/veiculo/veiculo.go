package veiculo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wellingtonpires/hackaton/auth/validatoken"
	"github.com/wellingtonpires/hackaton/domain/entity/veiculo"
	"github.com/wellingtonpires/hackaton/infrastructure/persistence"
)

func Cadastro(c *gin.Context) {
	if !validatoken.ValidaTokenAdmin(c.GetHeader("authorization")) {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"resultado": "Erro na autenticação"})
	} else {
		var ve veiculo.Veiculo
		err := c.ShouldBindJSON(&ve)
		if err != nil {
			fmt.Println(err.Error())
		}
		persistence.CadastraVeiculo(ve)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"resultado": "Erro: " + err.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, gin.H{"resultado": "Veiculo cadastrado na base"})
		}
	}
}

func Atualizacao(c *gin.Context) {
	if !validatoken.ValidaTokenAdmin(c.GetHeader("authorization")) {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"resultado": "Erro na autenticação"})
	} else {
		var ve veiculo.Veiculo
		err := c.ShouldBindJSON(&ve)
		if err != nil {
			fmt.Println(err.Error())
		}
		persistence.AtualizaVeiculo(ve)
		c.IndentedJSON(http.StatusOK, gin.H{"resultado": "Veículo atualizado na base"})
	}
}

func Exclusao(c *gin.Context) {
	if !validatoken.ValidaTokenAdmin(c.GetHeader("authorization")) {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"resultado": "Erro na autenticação"})
	} else {
		var ve veiculo.Veiculo
		err := c.ShouldBindJSON(&ve)
		if err != nil {
			fmt.Println(err.Error())
		}
		persistence.ExcluiVeiculo(ve)
		c.IndentedJSON(http.StatusOK, gin.H{"resultado": "Veículo deletado"})
	}
}

func ConsultaPorPreco(c *gin.Context) {
	//	if !validatoken.ValidaToken(c.GetHeader("authorization")) {
	//		c.IndentedJSON(http.StatusUnauthorized, gin.H{"resultado": "Erro na autenticação"})
	//	} else {
	c.IndentedJSON(http.StatusOK, persistence.ConsultaPorPreco())
	//	}
}

func ConsultaVendidos(c *gin.Context) {
	if !validatoken.ValidaTokenAdmin(c.GetHeader("authorization")) {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"resultado": "Erro na autenticação"})
	} else {
		c.IndentedJSON(http.StatusOK, persistence.ConsultaVendidos())
	}
}

func Checkout(c *gin.Context) {
	if !validatoken.ValidaTokenAdmin(c.GetHeader("authorization")) {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"resultado": "Erro na autenticação"})
	} else {
		var ve veiculo.Veiculo
		err := c.ShouldBindJSON(&ve)
		if err != nil {
			fmt.Println(err.Error())
		}
		persistence.Checkout(ve, c.GetHeader("authorization"))
		c.IndentedJSON(http.StatusOK, gin.H{"resultado": "Veículo adquirido com sucesso"})
	}
}
