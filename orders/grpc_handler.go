package main

import (
	"context"

	pb "github.com/haileyism/commons/api"
	"google.golang.org/grpc"
)

type hrpcHandler struct{
	pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server){
	handler := &grpcHandler{}
	pb.registerOrderServiceServer(grpcServer,handler)

}

func (h *grpcHandler) CreateOrder(context.Context, *pb.CreateOrderRequest)(*pb.Order,error){
	log.Println("New order received!")
	o:=&pb.Order{
		ID: "42",

	}
	return o,nil
}