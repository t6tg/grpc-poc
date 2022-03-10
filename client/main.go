package main

import (
	"log"

	"google.golang.org/grpc"
)

const serverAddress = "127.0.0.1:3000"

func main() {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("connection error : %v", err)
	}
	if conn == nil {
		log.Fatalln("connection error nil")
	}

}
