package injector

import (
	"hex-grpc/controllers"
	"hex-grpc/gateways/infra/db"
	article_repo "hex-grpc/gateways/repository/article"
	"hex-grpc/usecase/inputport"
	article_uc "hex-grpc/usecase/interactor/article"
	"hex-grpc/usecase/repository"

	"gorm.io/gorm"
)

func injectDB() *gorm.DB {
	return db.Connect()
}

func injectArticleRepository() repository.Article {
	return article_repo.New(injectDB())
}

func injectArticlePort() inputport.Article {
	return article_uc.New(injectArticleRepository())
}

func InjectArticleController() *controllers.ArticleController {
	return controllers.NewArticleController(injectArticlePort())
}
