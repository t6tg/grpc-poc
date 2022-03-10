package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/t6tg/grpc-poc/server/store/storepb"

	"google.golang.org/grpc"
)

var port int = 3001

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	storepb.RegisterStoreServer(grpcServer, newServer())
	fmt.Println("store service start at port 3001 ( grpc ) .... ")
	grpcServer.Serve(lis)
}

func newServer() storepb.StoreServer {
	return store{}
}

type store struct {
	storepb.StoreServer
}

func (store) GetStoreLatest(ctx context.Context, req *storepb.StoreLatestRequest) (*storepb.StoreLatestResponse, error) {
	var latest []*storepb.StoreLatests
	latest = append(latest, &storepb.StoreLatests{
		Id:          1,
		Title:       "novel dek-d",
		Description: "novel.dek-d.com",
		Coin:        30,
	})
	latest = append(latest, &storepb.StoreLatests{
		Id:          2,
		Title:       "novel dek-d",
		Description: "novel.dek-d.com",
		Coin:        20,
	})
	latest = append(latest, &storepb.StoreLatests{
		Id:          3,
		Title:       "novel dek-d",
		Description: "novel.dek-d.com",
		Coin:        10,
	})
	return &storepb.StoreLatestResponse{StoreLatests: latest}, nil
}
