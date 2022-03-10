package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/t6tg/grpc-poc/server/novelpb"

	"google.golang.org/grpc"
)

var port int = 3000

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	novelpb.RegisterNovelServer(grpcServer, newServer())
	grpcServer.Serve(lis)

}

func newServer() novelpb.NovelServer {
	return novel{}
}

type novel struct {
	novelpb.NovelServer
}

func (novel) GetNovelById(ctx context.Context, req *novelpb.NovelIdRequest) (*novelpb.NovelIdResponse, error) {
	fmt.Println("novel req: ", req.GetId())
	return &novelpb.NovelIdResponse{
		Id:          1,
		Title:       "dek-d",
		Description: "novel.dek-d.com",
	}, nil
}
