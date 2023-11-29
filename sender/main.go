/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	helloworld2 "github.com/Allan-Nava/GO-gRPC/proto/generated/helloworld.proto"
	start2 "github.com/Allan-Nava/GO-gRPC/proto/generated/start.proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
	port = flag.Int("port", 50052, "The server port")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	starterServer := &server{
		Client: helloworld2.NewGreeterClient(conn),
	}
	start2.RegisterStarterServer(s, starterServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	start2.UnimplementedStarterServer
	Client helloworld2.GreeterClient
}

func (s *server) Start(ctx context.Context, in *start2.StartRequest) (*start2.StartResponse, error) {
	log.Printf("Received: %v", in.GetName())
	go Send(s.Client)
	return &start2.StartResponse{Message: "Hello " + in.GetName()}, nil
}

func Send(c helloworld2.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	for i := 0; i < 100; i++ {
		r, err := c.SayHello(ctx, &helloworld2.HelloRequest{Name: *name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}
}
