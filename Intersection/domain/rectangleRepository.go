package domain

import (
	"Intersection/dto"
	"gorm.io/gorm"
)

type RectangleRepository interface {
	SaveRectangle(rectangle Rectangle)
	GetRectangles() ([]dto.RectangleDto, error)
}

type DefaultRectangleRepository struct{
	db *gorm.DB
}

func (d DefaultRectangleRepository) SaveRectangle(rectangle Rectangle) {
	d.db.Create(rectangle.ToDto())
}

func (d DefaultRectangleRepository) GetRectangles() ([]dto.RectangleDto, error){
	var rectangles []dto.RectangleDto
	d.db.Find(&rectangles)
	return rectangles, nil
}


