package article

import (
	"hex-grpc/entities"
	"hex-grpc/usecase/repository"

	"gorm.io/gorm"
)

type article struct {
	conn *gorm.DB
}

func NewArticleRepository(conn *gorm.DB) repository.Article {
	return &article{
		conn: conn,
	}
}

func (a *article) Create(article entities.Article) string {
	a.conn.Create(&article)
	return article.ID
}

func (a *article) Delete(ID string) {
	a.conn.Delete(entities.Article{}, ID)
}

func (a *article) FindOne(ID string) entities.Article {
	var article entities.Article
	a.conn.Where("id = ?", ID).First(&article)
	return article
}

func (a *article) GetAll() []entities.Article {
	var articles []entities.Article
	a.conn.Model(entities.Article{}).Find(&articles)
	return articles
}

func (a *article) Update(article entities.Article) {
	a.conn.Model(&article).Updates(article)
}
