package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	"github.com/fairwinds0099/itoagrpc/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type ItoaManagementServer struct {
}

func (server *ItoaManagementServer) CreateItoa(ctx context.Context, in *pb.CreateNewAnalytics) {

}

func (s *ItoaManagementServer) CreateNewAnalytics(ctx context.Context, in *pb.AnalyticsInput) {
	log.Printf("Received probability distribution function: %v", in.Getpdf)
	var cdf_id int32 = int32(rand.Intn(100))
	return &pb.Result{cdf: in.GetPdf(), id: cdf_id}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterItoaManagementServer(s, &ItoaManagementServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
