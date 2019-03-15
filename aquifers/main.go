package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	nats "github.com/nats-io/go-nats"
	pb "github.com/stephenhillier/groundwater/aquifers/proto/aquifers"
)

// Repository is the set of methods available to a collection of aquifer data
type Repository interface {
	Get(int32) (*pb.Aquifer, error)
}

// AquiferRepository is a stand-in for a datastore
type AquiferRepository struct {
	aquifers []*pb.Aquifer
}

// Get finds an aquifer in the repository matching the provided id
func (repo *AquiferRepository) Get(id int32) (*pb.Aquifer, error) {
	var aq *pb.Aquifer

	for _, aquifer := range repo.aquifers {
		if aquifer.Id == id {
			aq = aquifer
			break
		}
	}

	return aq, nil
}

type service struct {
	repo Repository
}

func (s *service) GetAquifer(ctx context.Context, req *pb.SingleAquiferRequest) (*pb.Aquifer, error) {
	aq, err := s.repo.Get(req.Id)
	if err != nil {
		log.Println(err)
	}
	return aq, nil
}

func main() {
	port := 7778

	sampleAquifers := []*pb.Aquifer{
		&pb.Aquifer{Id: 1, Name: "Aquifer A23", Volume: 543.21},
		&pb.Aquifer{Id: 2, Name: "Aquifer A60", Volume: 262.85},
	}

	repo := &AquiferRepository{sampleAquifers}

	nc, err := nats.Connect("nats:4222")
	if err != nil {
		log.Println("Error connecting to NATS:", err)
	}

	ncPb, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Println("Error created encoded connection:", err)
	}
	defer ncPb.Close()
	defer nc.Close()

	go listenForEvents(ncPb)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()

	pb.RegisterAquiferServiceServer(srv, &service{repo})

	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

func listenForEvents(c *nats.EncodedConn) {
	ch := make(chan *nats.Msg, 64)
	c.Subscribe("well-create", func(m *nats.Msg) {
		ch <- m
	})
	log.Println("Listening for events on well-create")
	for {
		msg := <-ch
		log.Println(string(msg.Data))
	}

}
