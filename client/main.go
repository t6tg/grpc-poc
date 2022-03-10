package main

import (
	"context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/t6tg/grpc-poc/client/novelpb"
	"google.golang.org/grpc"
)

const serverAddr = "127.0.0.1:3000"

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
	}
	defer conn.Close()
	if conn == nil {
		log.Fatal("connnection is nil")
	}
	client := novelpb.NewNovelClient(conn)
	r := gin.Default()
	r.GET("/novel/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Fatalln("erro param")
			return
		}
		res, err := client.GetNovelById(context.Background(), &novelpb.NovelIdRequest{
			Id: int32(id),
		})
		if err != nil {
			log.Fatalln("erro param")
			return
		}
		c.JSON(200, res)
	})
	if err := r.Run(); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
