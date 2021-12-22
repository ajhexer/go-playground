package app

import (
	"Intersection/DB"
	"Intersection/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func Start(){
	serv := gin.Default()
	repo := DB.NewDefaultRepo(NewDb())
	rectangleService := service.NewRectangleService(repo)
	rectangleHandler := RectangleHandlers{service: rectangleService}
	serv.POST("/", rectangleHandler.AddRectangle)
	serv.GET("/", rectangleHandler.GetRectangles)
	serv.Run(":8000")
}




func NewDb() *gorm.DB{
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err!=nil{
		log.Fatalln(err)
	}
	err = db.AutoMigrate(&DB.RectangleDbModel{})
	if err!=nil{
		log.Fatalln(err)
	}
	return db
}
