package app

import (
	"Banking/domain"
	"Banking/logger"
	"Banking/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
)


func Start(){
	serv := gin.Default()
	client := NewDb()
	ch:=CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb(client))}
	ah := AccountHandlers{service: service.NewAccountService(domain.NewAccountRepositoryDb(client))}
	serv.GET("/customers", ch.getAllCustomers)
	serv.GET("/customers/:id", ch.getCustomer)
	serv.POST("/customers/:id/accounts", ah.NewAccount)
	serv.POST("/customers/:id/accounts/:account_id/transactions", ah.NewTransaction)
	serv.Run(":8181")

}


func NewDb() *sqlx.DB{
	client, err := sqlx.Open("mysql", "user:password@tcp(localhost:5555)/banking")
	if err!=nil{
		logger.Log.Error("Unexpected sql errors")
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute*4)
	client.SetMaxIdleConns(15)
	client.SetMaxOpenConns(15)
	return client
}


