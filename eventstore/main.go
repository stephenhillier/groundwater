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
	natsURL  = "nats:4222"
	grpcPort = "9000"
)

type server struct {
	nats *nats.EncodedConn
}

func publishEvent(nc *nats.EncodedConn, event *pb.Event) {
	err := nc.Publish(event.EventType, event)
	if err != nil {
		log.Println("Error publishing message:", err)
	}
	log.Println("Published message:", event.EventType, event)
}

func (s *server) CreateEvent(ctx context.Context, event *pb.Event) (*pb.Response, error) {
	// todo: persist event in db here
	publishEvent(s.nats, event)
	return &pb.Response{Success: true}, nil
}

// func (s *server) sampleMsg() {
// 	for {
// 		time.Sleep(1 * time.Second)
// 		s.CreateEvent(context.Background(), &pb.Event{
// 			EventId:       "aaa1",
// 			EventType:     "aquifer-create",
// 			AggregateId:   "aaa",
// 			AggregateType: "aquifers",
// 			EventData:     "A1",
// 		})
// 	}
// }

func main() {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Println("Error connecting to NATS:", err)
	}

	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()
	defer nc.Close()

	natsService := &server{c}

	srv := grpc.NewServer()
	pb.RegisterEventServiceServer(srv, natsService)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("%s listening on port %v", clientID, grpcPort)

	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
