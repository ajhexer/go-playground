package DB

import (
	"Intersection/domain"
	"gorm.io/gorm"
)

type RectangleDbModel struct{
	gorm.Model
	X		int
	Y		int
	Width	int
	Height	int
}

func  RectangleToDbModel(r domain.Rectangle) *RectangleDbModel{
	rb := RectangleDbModel{
		X: r.X,
		Y: r.Y,
		Width: r.Width,
		Height: r.Height,
	}
	return &rb
}

//func (rd *RectangleDto) ToGetDto() RectangleGetDto{
//	r := RectangleGetDto{Rectangle: rd.ToRectangle(), Time: rd.CreatedAt}
//	return r
//}
//
//func (rd *RectangleDto) ToRectangle() domain.Rectangle{
//	r := domain.Rectangle{rd.X, rd.Y, rd.Width, rd.Height}
//	return r
//}

