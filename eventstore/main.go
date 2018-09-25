package main

import (
	"context"
	"fmt"
	"log"
	"net"

	nats "github.com/nats-io/go-nats"
	pb "github.com/stephenhillier/groundwater/eventstore/proto/events"
	grpc "google.golang.org/grpc"
)

const (
	clientID = "event-store-api"
	natsURL  = "localhost:4222"
	grpcPort = "9000"
)

type server struct {
	nats *nats.Conn
}

func publishEvent(nc *nats.Conn, event *pb.Event) {
	eventMsg := []byte(event.EventData)
	nc.Publish(event.EventType, eventMsg)
	log.Println("Published message:", event.EventType)
}

func (s *server) CreateEvent(ctx context.Context, event *pb.Event) (*pb.Response, error) {
	// todo: persist event in db here
	publishEvent(s.nats, event)
	return &pb.Response{Success: true}, nil
}

func main() {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Println("Error connecting to NATS:", err)
	}

	srv := grpc.NewServer()
	pb.RegisterEventServiceServer(srv, &server{nc})

	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("%s listening on port %v", clientID, grpcPort)

	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
