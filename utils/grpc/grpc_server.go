package utils_grpc

import "google.golang.org/grpc"

func (s *GrpcServer) GetServer() *grpc.Server {
	return s.server
}
