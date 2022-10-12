package main

import (
	"hex-grpc/gateways/infra/server"
	"hex-grpc/injector"
	"log"
	"net"
)

const (
	port = ":3000"
)

func main() {
	lis, _ := net.Listen("tcp", port)

	articleHandler := injector.InjectArticleController()
	s := server.Create(articleHandler)
	if err := s.Serve(lis); err != nil {
		log.Panicln(err)
	}
}
