package DB

import (
	"Intersection/domain"
	"gorm.io/gorm"
)



type RectangleRepository interface {
	SaveRectangle(rectangle domain.Rectangle)
	GetRectangles() ([]RectangleDbModel, error)
}

type DefaultRectangleRepository struct{
	db *gorm.DB
}

func (d DefaultRectangleRepository) SaveRectangle(rectangle domain.Rectangle) {
	d.db.Create(RectangleToDbModel(rectangle))
}

func (d DefaultRectangleRepository) GetRectangles() ([]RectangleDbModel, error){
	var rectangles []RectangleDbModel
	d.db.Find(&rectangles)
	return rectangles, nil
}

func NewDefaultRepo(db *gorm.DB) RectangleRepository{
	return DefaultRectangleRepository{db: db}
}
