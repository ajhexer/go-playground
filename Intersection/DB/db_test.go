package DB

import (
	"Intersection/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestSave(t *testing.T){
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&RectangleDbModel{})
	repo:= NewDefaultRepo(db)
	r := domain.Rectangle{
		X: 5,
		Y: 6,
		Width: 2,
		Height: 7,
	}
	repo.SaveRectangle(r)
	rectangles, _ := repo.GetRectangles()
	rb := rectangles[0]
	if rb.X!=r.X{
		t.Error()
	}else if rb.Y!=r.Y{
		t.Error()
	}

}

