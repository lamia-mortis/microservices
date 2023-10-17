package api

import (
	"gateway/pb"
	"gateway/util"
)

type Server struct {
	pb.UnimplementedBrokerServer
	config util.Config
}

func NewServer(config util.Config) (*Server, error) {

	server := &Server{
		config: config,
	}

	return server, nil
}
