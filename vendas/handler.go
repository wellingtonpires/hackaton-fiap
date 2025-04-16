package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	Routes(router)
	router.Run(":8000")
}

func Routes(route *gin.Engine) {
	v := route.Group("/external")
	v.GET("/consulta-por-preco", func(c *gin.Context) {

		url := "http://localhost:8080/veiculo/consulta-por-preco"

		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Erro ao fazer a requisição: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao obter dados externos"})
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Erro ao ler o corpo da resposta: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao ler dados externos"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": string(body),
		})
	})

	v.POST("/checkout", func(c *gin.Context) {

		var requestData map[string]interface{}

		if err := c.ShouldBindJSON(&requestData); err != nil {
			log.Printf("Erro ao processar JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de JSON inválido"})
			return
		}

		jsonData, err := json.Marshal(requestData)
		if err != nil {
			log.Printf("Erro ao converter JSON: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno ao processar os dados"})
			return
		}

		url := "http://localhost:8080/veiculo/checkout"

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("Erro ao chamar o serviço externo: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao conectar ao serviço externo"})
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Erro ao ler resposta do serviço externo: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar resposta do serviço externo"})
			return
		}

		c.JSON(resp.StatusCode, gin.H{
			"message": "Resposta do serviço externo",
			"data":    string(body),
		})

	})
}
