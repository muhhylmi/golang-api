package utils_grpc

import (
	"errors"

	"google.golang.org/grpc"
)

func NewGrpcServices(hosts *GrpcServiceHosts) (GrpcServicesInterface, error) {
	return &GrpcServices{
		hosts:    hosts,
		connPool: make(map[string]*grpc.ClientConn),
	}, nil
}

func NewGrpcServer(opts *GrpcServerOpts) (GrpcServerInterface, error) {
	if opts.ServiceName == "" {
		return nil, errors.New("service name cannot be empty")
	}

	server := &GrpcServer{
		serviceName: opts.ServiceName,
		log:         opts.Logger,
	}

	gs := grpc.NewServer()

	server.server = gs

	return server, nil
}
