package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接异常", err)
	}
	defer conn.Close()

	client := pb.NewUserInfoServiceClient(conn)
	req := new(pb.UserRequest)
	req.Name = "kiri"

	response, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Println("响应异常", err)
	}
	fmt.Println("响应结果", response)
}
