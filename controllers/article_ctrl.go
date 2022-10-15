package controllers

import (
	"context"
	"hex-grpc/entities"
	articlev1 "hex-grpc/proto/article/v1"
	"hex-grpc/usecase/inputport"

	"github.com/golang/protobuf/ptypes/empty"
)

type ArticleController struct {
	articlev1.UnimplementedArticleServiceServer
	article inputport.Article
}

func NewArticleController(article inputport.Article) *ArticleController {
	return &ArticleController{
		article: article,
	}
}

func (g *ArticleController) DeleteArticle(_ context.Context, req *articlev1.DeleteArticleRequest) (*empty.Empty, error) {
	articleID := req.GetId()
	g.article.Delete(articleID)
	return &empty.Empty{}, nil
}

func (g *ArticleController) GetArticle(_ context.Context, req *articlev1.GetArticleRequest) (*articlev1.GetArticleResponse, error) {
	articleID := req.GetId()
	article := g.article.FindOne(articleID)
	return &articlev1.GetArticleResponse{
		Article: &articlev1.Article{
			Id:    article.ID,
			Title: article.Title,
			Price: article.Price,
		},
	}, nil
}

func (g *ArticleController) GetArticles(_ context.Context, _ *empty.Empty) (*articlev1.GetArticlesResponse, error) {
	articles := g.article.GetAll()
	return &articlev1.GetArticlesResponse{
		Articles: convertArticle(&articles),
	}, nil
}

func (g *ArticleController) CreateArticle(_ context.Context, req *articlev1.CreateArticleRequest) (*articlev1.CreateArticleResponse, error) {
	articleID := g.article.Create(entities.Article{
		Title: req.GetTitle(),
		Price: req.GetPrice(),
	})
	return &articlev1.CreateArticleResponse{
		Id: articleID,
	}, nil
}

func (g *ArticleController) UpdateArticle(_ context.Context, req *articlev1.UpdateArticleRequest) (*empty.Empty, error) {
	pbArticle := req.GetArticle()
	article := entities.Article{
		ID:    pbArticle.GetId(),
		Title: pbArticle.GetTitle(),
		Price: pbArticle.GetPrice(),
	}
	g.article.Update(article)
	return &empty.Empty{}, nil
}

func convertArticle(articles *[]entities.Article) []*articlev1.Article {
	var pbArticles []*articlev1.Article
	for _, article := range *articles {
		pbArticles = append(pbArticles, &articlev1.Article{
			Id:    article.ID,
			Title: article.Title,
			Price: article.Price,
		})
	}
	return pbArticles
}
