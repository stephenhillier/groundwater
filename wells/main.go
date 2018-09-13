package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	pb "github.com/stephenhillier/groundwater/wells/proto/wells"
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

func (s *service) GetWells(ctx context.Context, req *pb.GetRequest, res *pb.ListResponse) error {
	wells := s.repo.GetAll()
	res.Wells = wells
	return nil
}

func main() {

	// example wells
	sampleWells := []*pb.Well{
		&pb.Well{Id: 1, Aquifer: 1, Owner: "Steve", Depth: 99.1},
		&pb.Well{Id: 2, Aquifer: 1, Owner: "Steve", Depth: 41.4},
		&pb.Well{Id: 3, Aquifer: 2, Owner: "Steve", Depth: 67.8},
	}

	repo := &WellRepository{sampleWells}

	srv := k8s.NewService(
		micro.Name("wells"),
	)
	srv.Init()

	pb.RegisterWellServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
