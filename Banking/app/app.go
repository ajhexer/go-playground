package app

import (
	"Banking/domain"
	"Banking/service"
	"github.com/gin-gonic/gin"
)


func Start(){
	serv := gin.Default()
	ch:=CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	serv.GET("/customers", ch.getAllCustomers)
	serv.Run(":8181")
}


