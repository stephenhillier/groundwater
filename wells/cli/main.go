package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/stephenhillier/groundwater/wells/proto/wells"
	"google.golang.org/grpc"
)

func main() {
	host := "127.0.0.1"
	port := "7777"
	var conn *grpc.ClientConn
	var err error

	conn, err = grpc.Dial(fmt.Sprintf("%s:%v", host, port), grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	client := pb.NewWellServiceClient(conn)

	wells, err := client.GetWells(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Println(err)
	}
	log.Println(wells)
}
