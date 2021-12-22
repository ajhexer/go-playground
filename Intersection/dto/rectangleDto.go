package dto

import (
	"Intersection/DB"
	"Intersection/domain"
	"time"
)


type RectangleRequestDto struct{
	Main 		domain.Rectangle 	`json:"main"`
	Input 		[]domain.Rectangle 	`json:"input"`
}

type RectangleResponseDto struct{
	domain.Rectangle
	Time	time.Time `json:"time"`
}


func ToResponseDto(r DB.RectangleDbModel) RectangleResponseDto{
	res := RectangleResponseDto{
		Rectangle: rectangleDbToRectangle(r),
		Time: r.CreatedAt,
	}
	return res
}



func rectangleDbToRectangle(r DB.RectangleDbModel) domain.Rectangle{
	return domain.Rectangle{
		X: r.X,
		Y: r.Y,
		Width: r.Width,
		Height: r.Height,
	}
}


