package service

import (
	"Intersection/domain"
)

type RectangleService interface{
	AddRectangle(main domain.Rectangle, rectangles []domain.Rectangle) error
	GetAllRectangles() ([]domain.Rectangle, error)
}


type DefaultRectangleService struct{
	repo domain.RectangleRepository
}

func (d DefaultRectangleService) GetAllRectangles() ([]domain.Rectangle, error){
	return d.repo.GetRectangles()
}

func (d DefaultRectangleService) AddRectangle(main domain.Rectangle, rectangles []domain.Rectangle) error{
	for _, r:= range rectangles{
		if main.CheckIntersection(r){
			err := d.repo.SaveRectangle(r)
			if err!=nil{
				return err
			}
		}
	}
	return nil
}







