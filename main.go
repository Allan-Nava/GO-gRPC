package main


import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "helloworld/helloworld" // Update with your actual module path
)


func main(){
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	fmt.Println("Server is running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}