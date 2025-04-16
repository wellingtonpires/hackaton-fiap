package veiculo_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wellingtonpires/hackaton/auth/cadastrausuario"
	"github.com/wellingtonpires/hackaton/auth/criatoken"
	"github.com/wellingtonpires/hackaton/domain/aggregate/veiculo"
	"github.com/wellingtonpires/hackaton/webhook/pagamento"
)

func Routes(route *gin.Engine) {
	v := route.Group("/veiculo")
	v.GET("/vendidos", veiculo.ConsultaVendidos)
	v.GET("/consulta-por-preco", veiculo.ConsultaPorPreco)
	v.POST("/checkout", veiculo.Checkout)
	v.POST("/cadastra-veiculo", veiculo.Cadastro)
	v.PATCH("/atualiza-veiculo", veiculo.Atualizacao)
	v.DELETE("/exclui-veiculo", veiculo.Exclusao)

	u := route.Group("/usuario")
	u.POST("/cria-token", criatoken.CriaToken)
	u.POST("/cadastra-usuario", cadastrausuario.CadastraUsuario)

	w := route.Group("/webhook")
	w.POST("/pagamento", pagamento.Pagamento)
}
