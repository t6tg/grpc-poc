package main

import (
	"context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/t6tg/grpc-poc/client/novelpb"
	"github.com/t6tg/grpc-poc/client/storepb"
	"google.golang.org/grpc"
)

const serverAddr = "127.0.0.1:3000"
const serverAddrStore = "127.0.0.1:3001"

func main() {
	connNovel, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
	}
	defer connNovel.Close()
	if connNovel == nil {
		log.Fatal("connnection is nil")
	}
	novelClient := novelpb.NewNovelClient(connNovel)

	connStore, err := grpc.Dial(serverAddrStore, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
	}
	defer connNovel.Close()
	if connNovel == nil {
		log.Fatal("connnection is nil")
	}
	storeClient := storepb.NewStoreClient(connStore)

	r := gin.Default()
	r.GET("/novel/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Fatalln("error param")
			return
		}
		res, err := novelClient.GetNovelById(context.Background(), &novelpb.NovelIdRequest{
			Id: int32(id),
		})
		if err != nil {
			log.Fatalln("error connection")
			return
		}
		c.JSON(200, res)
	})

	r.GET("/store/latest", func(c *gin.Context) {
		res, err := storeClient.GetStoreLatest(context.Background(), &storepb.StoreLatestRequest{})
		if err != nil {
			log.Fatalln("error connect")
			return
		}
		c.JSON(200, res.StoreLatests)
	})

	if err := r.Run(); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
