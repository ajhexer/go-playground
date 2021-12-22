package service

import (
	"Intersection/DB"
	"Intersection/domain"
)

type RectangleService interface{
	AddRectangle(main domain.Rectangle, rectangles []domain.Rectangle) error
	GetAllRectangles() ([]DB.RectangleDbModel, error)
}


type DefaultRectangleService struct{
	repo DB.RectangleRepository
}

func (d DefaultRectangleService) GetAllRectangles() ([]DB.RectangleDbModel, error){
	return d.repo.GetRectangles()
}

func (d DefaultRectangleService) AddRectangle(main domain.Rectangle, rectangles []domain.Rectangle) error{
	for _, r:= range rectangles{
		if main.CheckIntersection(r){
			d.repo.SaveRectangle(r)
		}
	}
	return nil
}
func NewRectangleService(repo DB.RectangleRepository) RectangleService{
	return DefaultRectangleService{repo: repo}
}







