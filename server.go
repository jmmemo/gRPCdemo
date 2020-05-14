package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc/proto"
	"net"
)

type UserInfoService struct {
}

var u = UserInfoService{}

func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	name := req.Name

	if name == "kiri" {
		resp = &pb.UserResponse{
			Id:    1,
			Name:  name,
			Age:   22,
			Hobby: []string{"Sing", "Run"},
		}
	}
	return
}
func main() {
	addr := "127.0.0.1:8080"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("监听异常", err)
	}
	fmt.Println("监听端口", addr)
	server := grpc.NewServer()
	pb.RegisterUserInfoServiceServer(server, &u)
	server.Serve(listen)
}
