package app

import (
	"Intersection/dto"
	"Intersection/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RectangleHandlers struct {
	service service.RectangleService
}

func (a RectangleHandlers) AddRectangle(c *gin.Context)  {
	var data dto.RectangleRequestDto
	if err:=c.ShouldBind(&data); err!=nil{
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	err := a.service.AddRectangle(data.Main, data.Input)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error in saving objects"})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (a RectangleHandlers) GetRectangles(c *gin.Context){
	var getRectangles []dto.RectangleResponseDto
	rectangles, err:=a.service.GetAllRectangles()
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error occurred while getting objects"})
		return
	}
	for _, r := range rectangles{
		getRectangles = append(getRectangles, dto.ToResponseDto(r))
	}
	c.JSON(http.StatusOK, gin.H{"Rectangles":getRectangles})
}

