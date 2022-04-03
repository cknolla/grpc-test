package server

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"grpc-test/proto"
)

type Server struct{}

func New() *grpc.Server {
	log.Info().Msg("creating new grpc server")
	creds := insecure.NewCredentials()

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	proto.RegisterGrpcServiceServer(grpcServer, &Server{})
	reflection.Register(grpcServer)
	return grpcServer
}

func (*Server) Hello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	log.Info().Msg("sending hello")
	return &proto.HelloResponse{
		Message: "hello",
	}, nil
}
