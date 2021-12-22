package service

import (
	"Intersection/domain"
	"Intersection/dto"
)

type RectangleService interface{
	AddRectangle(main domain.Rectangle, rectangles []domain.Rectangle) error
	GetAllRectangles() ([]domain.Rectangle, error)
}


type DefaultRectangleService struct{
	repo domain.RectangleRepository
}

func (d DefaultRectangleService) GetAllRectangles() ([]dto.RectangleDto, error){
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







