package dto

import (
	"Intersection/domain"
	"gorm.io/gorm"
)

type RectangleDto struct{
	gorm.Model
	X		int
	Y		int
	Width	int
	Height	int
}

func (rd *RectangleDto) ToRectangle() domain.Rectangle{
	r := domain.Rectangle{rd.X, rd.Y, rd.Width, rd.Height}
	return r
}


