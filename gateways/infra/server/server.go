package server

import (
	"hex-grpc/controllers"
	articlev1 "hex-grpc/proto/article/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Create(articleHandler *controllers.ArticleController) *grpc.Server {
	s := grpc.NewServer()
	articlev1.RegisterArticleServiceServer(s, articleHandler)
	reflection.Register(s)
	return s
}
