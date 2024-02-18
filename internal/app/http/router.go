package http

import (
	"github.com/fasthttp/router"
	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/http/controller"
)

func GetRouter() *router.Router {
	r := router.New()
	r.POST("/clientes/{id}/transacoes", controller.ClientTrasactionController)
	r.GET("/clientes/{id}/extrato", controller.ClientExtractController)
	return r
}
