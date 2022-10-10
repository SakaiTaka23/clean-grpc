package grpc

import (
	"context"
	"hex-grpc/core/model"
	"hex-grpc/core/ports"
	"hex-grpc/core/service"
	articlev1 "hex-grpc/proto/article/v1"

	"github.com/golang/protobuf/ptypes/empty"
)

type GrpcAdapter struct {
	articlev1.UnimplementedArticleServiceServer
	service ports.PrimaryPorts
}

func NewGrpcAdapter(service service.ArticleService) GrpcAdapter {
	return GrpcAdapter{
		service: &service,
	}
}

func (g *GrpcAdapter) DeleteArticle(_ context.Context, req *articlev1.DeleteArticleRequest) (*empty.Empty, error) {
	articleID := req.GetId()
	g.service.Delete(articleID)
	return &empty.Empty{}, nil
}

func (g *GrpcAdapter) GetArticle(_ context.Context, req *articlev1.GetArticleRequest) (*articlev1.GetArticleResponse, error) {
	articleID := req.GetId()
	article := g.service.FindOne(articleID)
	return &articlev1.GetArticleResponse{
		Article: &articlev1.Article{
			Id:    article.ID,
			Title: article.Title,
			Price: article.Price,
		},
	}, nil
}

func (g *GrpcAdapter) GetArticles(_ context.Context, _ *empty.Empty) (*articlev1.GetArticlesResponse, error) {
	articles := g.service.GetAll()
	return &articlev1.GetArticlesResponse{
		Articles: convertArticle(&articles),
	}, nil
}

func (g *GrpcAdapter) CreateArticle(_ context.Context, req *articlev1.CreateArticleRequest) (*articlev1.CreateArticleResponse, error) {
	articleID := g.service.Create(model.Article{
		Title: req.GetTitle(),
		Price: req.GetPrice(),
	})
	return &articlev1.CreateArticleResponse{
		Id: articleID,
	}, nil
}

func (g *GrpcAdapter) UpdateArticle(_ context.Context, req *articlev1.UpdateArticleRequest) (*empty.Empty, error) {
	pbArticle := req.GetArticle()
	article := model.Article{
		ID:    pbArticle.GetId(),
		Title: pbArticle.GetTitle(),
		Price: pbArticle.GetPrice(),
	}
	g.service.Update(article)
	return &empty.Empty{}, nil
}

func convertArticle(articles *[]model.Article) []*articlev1.Article {
	pbArticles := make([]*articlev1.Article, len(*articles))
	for _, article := range *articles {
		pbArticles = append(pbArticles, &articlev1.Article{
			Id:    article.ID,
			Title: article.Title,
			Price: article.Price,
		})
	}
	return pbArticles
}
