package article

import (
	"hex-grpc/entities"
	"hex-grpc/usecase/inputport"
	"hex-grpc/usecase/repository"
)

type interactor struct {
	repo repository.Article
}

func NewArticleInteractor(repo repository.Article) inputport.Article {
	return &interactor{repo: repo}
}

func (a *interactor) Create(article entities.Article) string {
	return a.repo.Create(article)
}

func (a *interactor) Delete(ID string) {
	a.repo.Delete(ID)
}

func (a *interactor) FindOne(ID string) entities.Article {
	return a.repo.FindOne(ID)
}

func (a *interactor) GetAll() []entities.Article {
	return a.repo.GetAll()
}

func (a *interactor) Update(article entities.Article) {
	a.repo.Update(article)
}
