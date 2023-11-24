// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: proto/start.proto

package start

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Starter_Start_FullMethodName = "/start.Starter/Start"
)

// StarterClient is the client API for Starter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StarterClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
}

type starterClient struct {
	cc grpc.ClientConnInterface
}

func NewStarterClient(cc grpc.ClientConnInterface) StarterClient {
	return &starterClient{cc}
}

func (c *starterClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, Starter_Start_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StarterServer is the server API for Starter service.
// All implementations must embed UnimplementedStarterServer
// for forward compatibility
type StarterServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	mustEmbedUnimplementedStarterServer()
}

// UnimplementedStarterServer must be embedded to have forward compatible implementations.
type UnimplementedStarterServer struct {
}

func (UnimplementedStarterServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedStarterServer) mustEmbedUnimplementedStarterServer() {}

// UnsafeStarterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StarterServer will
// result in compilation errors.
type UnsafeStarterServer interface {
	mustEmbedUnimplementedStarterServer()
}

func RegisterStarterServer(s grpc.ServiceRegistrar, srv StarterServer) {
	s.RegisterService(&Starter_ServiceDesc, srv)
}

func _Starter_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarterServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Starter_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarterServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Starter_ServiceDesc is the grpc.ServiceDesc for Starter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Starter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "start.Starter",
	HandlerType: (*StarterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Starter_Start_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/start.proto",
}
