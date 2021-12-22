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


