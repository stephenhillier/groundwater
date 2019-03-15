package main

import (
	"context"
	"fmt"
	"log"
	"net"

	nats "github.com/nats-io/go-nats"
	pb "github.com/stephenhillier/groundwater/wells/proto/wells"
	"google.golang.org/grpc"
)

// Repository is the set of methods available to use with a well data repo
type Repository interface {
	GetAll() []*pb.Well
	Search(*pb.WellSearchRequest) ([]*pb.Well, error)
}

// WellRepository is a stand-in for a datastore
type WellRepository struct {
	wells []*pb.Well
}

// GetAll returns all wells in the repo
func (repo *WellRepository) GetAll() []*pb.Well {
	return repo.wells
}

// Search returns wells meeting all of the provided search criteria
func (repo *WellRepository) Search(req *pb.WellSearchRequest) ([]*pb.Well, error) {
	search := make([]*pb.Well, 0)

	for _, well := range repo.wells {
		if (req.Owner == "" || well.Owner == req.Owner) &&
			(req.Aquifer == 0 || well.Aquifer == req.Aquifer) {
			search = append(search, well)
		}
	}

	return search, nil
}


func (repo *WellRepository) CreateWell(pb.)

// service represents a Wells service with a repository of wells.
// The repository can be anything (a database, collection of fake objects, etc)
// as long as it implements methods defined in the Repository interface.
type service struct {
	repo     Repository
	messages *nats.EncodedConn
}

func (s *service) GetWells(ctx context.Context, req *pb.GetRequest) (*pb.ListResponse, error) {
	wells := s.repo.GetAll()
	return &pb.ListResponse{Wells: wells}, nil
}

func (s *service) FindWells(ctx context.Context, req *pb.WellSearchRequest) (*pb.ListResponse, error) {
	wells, err := s.repo.Search(req)
	return &pb.ListResponse{Wells: wells}, err
}

func (s *service) Health(ctx context.Context, req *pb.GetRequest) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{Ok: true}, nil
}

func main() {
	port := 7777

	// example wells
	sampleWells := []*pb.Well{
		&pb.Well{Id: 1, Aquifer: 1, Owner: "Steve", Depth: 99.1},
		&pb.Well{Id: 2, Aquifer: 1, Owner: "Steve", Depth: 41.4},
		&pb.Well{Id: 3, Aquifer: 2, Owner: "Steve", Depth: 67.8},
	}

	repo := &WellRepository{sampleWells}

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

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

	log.Print("Events client ready")

	srv := grpc.NewServer()
	pb.RegisterWellServiceServer(srv, &service{repo})
	log.Println("Listening on port", port)
	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}

}
