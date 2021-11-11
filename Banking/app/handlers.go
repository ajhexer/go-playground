package app

import (
	"Banking/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomerHandlers struct{
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(c *gin.Context)  {
	customers, err := ch.service.GetAllCustomer()
	if err!=nil{
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(200, gin.H{
		"customers":customers,
	})
}


