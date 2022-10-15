package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	ID    string
	Title string
	Price string
}

func (a *Article) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := uuid.NewUUID()
	a.ID = id.String()
	return
}
