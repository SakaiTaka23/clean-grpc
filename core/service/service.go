package service

import (
	"hex-grpc/core/model"
	"hex-grpc/core/ports"
)

type ArticleService struct {
	repo ports.SecondaryPorts
}

func NewArticleService() ports.PrimaryPorts {
	return &ArticleService{}
}

func (a *ArticleService) Create(article model.Article) string {
	return a.repo.Create(article)
}

func (a *ArticleService) Delete(ID string) {
	a.repo.Delete(ID)
}

func (a *ArticleService) FindOne(ID string) model.Article {
	return a.repo.FindOne(ID)
}

func (a *ArticleService) GetAll() []model.Article {
	return a.repo.GetAll()
}

func (a *ArticleService) Update(article model.Article) {
	a.repo.Update(article)
}
