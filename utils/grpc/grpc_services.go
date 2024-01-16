package utils_grpc

import (
	"errors"
	"fmt"
	"golang-api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	ConnectionBooks = "books_grpc"
)

func (s *GrpcServices) GetConnectionFromPool(serviceKey string, host *string) (*grpc.ClientConn, error) {
	// check if connection already exists
	if con, ok := s.connPool[serviceKey]; ok {
		return con, nil
	}

	// check if host is not nil or empty
	if host == nil || *host == "" {
		return nil, errors.New("host is empty")
	}

	// dial new connection
	address := fmt.Sprintf("%s:///%s", "dns", *host)
	grpcCred := grpc.WithTransportCredentials(insecure.NewCredentials())

	conn, err := grpc.Dial(address, grpcCred)
	if err != nil {
		return nil, err
	}

	// save connection to pool
	s.connPool[serviceKey] = conn

	return conn, nil
}

func (s *GrpcServices) GetBookService() (proto.BookServiceClient, error) {
	conn, err := s.GetConnectionFromPool(ConnectionBooks, s.hosts.BookHost)
	if err != nil {
		return nil, err
	}
	return proto.NewBookServiceClient(conn), nil
}
