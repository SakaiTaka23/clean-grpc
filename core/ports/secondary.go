package ports

import "hex-grpc/core/model"

type SecondaryPorts interface {
	Create(model.Article) string
	Delete(ID string)
	FindOne(ID string) model.Article
	GetAll() []model.Article
	Update(model.Article)
}