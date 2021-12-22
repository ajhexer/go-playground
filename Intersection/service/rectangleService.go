package service

import "Intersection/domain"

type RectangleService interface{
	AddRectangle(main domain.Rectangle, rectangles []domain.Rectangle)
}


type DefaultRectangleService struct{

}

func (d DefaultRectangleService) AddRectangle(main domain.Rectangle, rectangles []domain.Rectangle) {
	panic("implement me")
}







