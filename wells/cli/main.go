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

	health, err := client.Health(context.Background(), &pb.GetRequest{})
	if health.Ok != true || err != nil {
		log.Println(err)
	}

	log.Println("Health check passed")

	wells, err := client.GetWells(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Println(err)
	}
	log.Println(wells)

	searchParams := &pb.WellSearchRequest{
		Aquifer: 1,
		Owner:   "Steve",
	}
	wellSearch, err := client.FindWells(context.Background(), searchParams)
	if err != nil {
		log.Println(err)
	}

	log.Println(wellSearch)
}
