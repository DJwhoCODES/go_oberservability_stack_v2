package server

import (
	pb "instagram/user-service/internal/gen/user/proto"
	"net"

	"google.golang.org/grpc"
)

func StartGRPC(port string, handler pb.UserServiceServer) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, handler)

	return s.Serve(lis)
}
