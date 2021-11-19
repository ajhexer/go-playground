package app

import (
	"Banking/service"
	"Banking/transfer"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountHandlers struct{
	service service.AccountService
}

func (a AccountHandlers) NewAccount(c *gin.Context) {
	var request transfer.AccountRequest
	if err:=c.ShouldBindJSON(&request); err!=nil{
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	account, err:=a.service.NewAccount(request)
	if err!=nil{
		c.JSON(err.Code, nil)
		return
	}
	c.JSON(http.StatusOK, account)
}

func (a AccountHandlers) NewTransaction(c *gin.Context){
	var request transfer.TransactionRequest
	if err:=c.ShouldBindJSON(&request); err!=nil{
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	transaction, err := a.service.NewTransaction(request)
	if err!=nil{
		c.JSON(err.Code, nil)
		return
	}
	c.JSON(http.StatusOK, transaction)

}








