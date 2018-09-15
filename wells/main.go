package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/stephenhillier/groundwater/wells/proto/wells"
	"google.golang.org/grpc"
)

// Repository is the set of methods available to use with a well data repo
type Repository interface {
	GetAll() []*pb.Well
}

// WellRepository is a stand-in for a datastore
type WellRepository struct {
	wells []*pb.Well
}

// GetAll returns all wells in the repo
func (repo *WellRepository) GetAll() []*pb.Well {
	return repo.wells
}

type service struct {
	repo Repository
}

func (s *service) GetWells(ctx context.Context, req *pb.GetRequest) (*pb.ListResponse, error) {
	wells := s.repo.GetAll()
	return &pb.ListResponse{Wells: wells}, nil
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

	srv := grpc.NewServer()

	pb.RegisterWellServiceServer(srv, &service{repo})

	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}

}
