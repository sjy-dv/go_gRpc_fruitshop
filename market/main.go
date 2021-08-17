package main

import (
	"context"
	"market/pb"
	"net"

	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

type Server struct {}

func (s *Server) ChoiceNum(ctx context.Context, snum *pb.SelectFruitNum) (*pb.FruitDetail, error) {

	num := snum.Num
	if num == 1 {
		return &pb.FruitDetail{
			Num: 1,
			Name: "banana",
			Price: 3000,
			Origin: "asia",
		}, nil
	}

	if num == 2 {
		return &pb.FruitDetail{
			Num: 2,
			Name: "apple",
			Price: 1500,
			Origin: "asia",
		}, nil
	}

	if num == 3 {
		return &pb.FruitDetail{
			Num: 3,
			Name: "cherry",
			Price: 5500,
			Origin: "europe",
		}, nil
	}

	return &pb.FruitDetail{}, nil
}


func (s *Server) ChoiceName(ctx context.Context, sname *pb.SelectFruitName) (*pb.FruitDetail, error) {

	name := sname.Name

	if name == "banana" {
		return &pb.FruitDetail{
			Num: 1,
			Name: "banana",
			Price: 3000,
			Origin: "asia",
		}, nil
	}

	if name == "apple" {
		return &pb.FruitDetail{
			Num: 2,
			Name: "apple",
			Price: 1500,
			Origin: "asia",
		}, nil
	}

	if name == "cherry" {
		return &pb.FruitDetail{
			Num: 3,
			Name: "cherry",
			Price: 5500,
			Origin: "europe",
		}, nil
	}

	return &pb.FruitDetail{}, nil
}

func main() { 

	conn, _ = grpc.Dial("localhost:8081", grpc.WithInsecure())

	defer conn.Close()

	//r := mux.NewRouter()

	lis , _ := net.Listen("tcp","localhost:8082")

	grpcServer := grpc.NewServer()

	server := Server{}

	//pb.RegisterGRpc_1Server(grpcServer, &server)
	pb.RegisterGRpc_2Server(grpcServer, &server)
	
	//lis := http.ListenAndServe(":8082", r)

	grpcServer.Serve(lis)
}