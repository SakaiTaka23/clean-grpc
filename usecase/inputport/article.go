package inputport

import "hex-grpc/entities"

type Article interface {
	Create(entities.Article) string
	Delete(ID string)
	FindOne(ID string) entities.Article
	GetAll() []entities.Article
	Update(entities.Article)
}
