package pagamento

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wellingtonpires/hackaton/domain/entity/veiculo"
	"github.com/wellingtonpires/hackaton/infrastructure/persistence"
)

func Pagamento(c *gin.Context) {
	var ve veiculo.Veiculo
	err := c.ShouldBindJSON(&ve)
	if err != nil {
		fmt.Println(err.Error())
	}
	persistence.Pagamento(ve)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"resultado": "Erro: " + err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"resultado": "Notificação de pagamento enviada"})
	}
}
