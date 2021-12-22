package domain


type RectangleRepository interface {
	SaveRectangle(rectangle Rectangle) error
	GetRectangles() ([]Rectangle, error)
}

type DefaultRectangleRepository struct{

}

func (d DefaultRectangleRepository) SaveRectangle(rectangle Rectangle) error{
	panic("implement me")
}

func (d DefaultRectangleRepository) GetRectangles() ([]Rectangle, error){
	panic("implement me")
}


