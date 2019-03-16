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

// Repository is the set of methods used to work with a collection of aquifers
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

// service represents the aquifers service; the service is created with a Repository that manages
// an aquifer collection. It will receive rpc methods that can be called by an aquifer client.
// the client can be created by importing "github.com/stephenhillier/groundwater/aquifers/proto/aquifers"
type service struct {
	repo Repository
}

// GetAquifer returns an aquifer by ID
func (s *service) GetAquifer(ctx context.Context, req *pb.SingleAquiferRequest) (*pb.Aquifer, error) {
	aq, err := s.repo.Get(req.Id)
	if err != nil {
		log.Println(err)
	}
	return aq, nil
}

func main() {
	port := 7778

	// sample collection of aquifers that will be used to populate a Repository
	sampleAquifers := []*pb.Aquifer{
		&pb.Aquifer{Id: 1, Name: "Aquifer A23", Volume: 543.21},
		&pb.Aquifer{Id: 2, Name: "Aquifer A60", Volume: 262.85},
	}

	repo := &AquiferRepository{sampleAquifers}

	// messages are published and subscribed to using NATS: open a NATS connection and then use
	// an EncodedConn wrapper to serialize the messages.
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

	// after opening the NATS connection, start listening for new
	// messages.  The subjects to subscribe to are defined within listenForEvents
	go listenForEvents(ncPb)

	// this service is accessible via gRPC. Start listening on a port
	// and then register the aquifer service; see proto definitions in
	// ./proto/aquifers/aquifers.proto
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

// listenForEvents creates channels and then subscribes
// to message subjects, placing messages on the channel when received.
// In this example, we are interested in messages about wells that may
// have been created in or near an aquifer.
func listenForEvents(c *nats.EncodedConn) {
	ch := make(chan *nats.Msg, 64)
	c.Subscribe("well-created", func(m *nats.Msg) {
		ch <- m
	})
	log.Println("Listening for events on well-create")
	for {
		msg := <-ch
		log.Println("received well created notification:", string(msg.Data))
	}
}
