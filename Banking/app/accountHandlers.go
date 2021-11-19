package app

import (
	"Banking/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ch *CustomerHandlers) getAllCustomers(c *gin.Context)  {
	customers, err := ch.service.GetAllCustomer()
	if err!=nil{
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(200, gin.H{
		"customers":customers,
	})
	logger.Log.Info("Test")
}

func (ch *CustomerHandlers) getCustomer(c *gin.Context){
	t :=c.Param("id")
	customer, err :=ch.service.GetCustomer(t)
	if err!=nil{
		c.JSON(err.Code, nil)
		return
	}
	c.JSON(http.StatusOK, customer)
}



