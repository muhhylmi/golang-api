package utils_grpc

import (
	"golang-api/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type (
	GrpcServices struct {
		hosts    *GrpcServiceHosts
		connPool map[string]*grpc.ClientConn
	}

	GrpcServicesInterface interface {
		GetBookService() (proto.BookServiceClient, error)
	}

	GrpcServer struct {
		serviceName string
		server      *grpc.Server
		log         *logrus.Logger
	}

	GrpcServerInterface interface {
		// Returns the gRPC server
		//
		// Returns:
		//      - *grpc.Server: the gRPC server
		GetServer() *grpc.Server
	}
)

type (
	GrpcServerOpts struct {
		Logger      *logrus.Logger
		ServiceName string
	}

	GrpcServiceHosts struct {
		BookHost *string
	}
)
