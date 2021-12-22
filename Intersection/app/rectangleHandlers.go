package app

import (
	"Intersection/service"
	"github.com/gin-gonic/gin"
)

type RectangleHandlers struct {
	service service.RectangleService
}

func (a RectangleHandlers) AddRectangle(c *gin.Context)  {

}

func (a RectangleHandlers) GetRectangles(c *gin.Context){

}

