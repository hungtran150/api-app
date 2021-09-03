package server

import (
	"log"

	"google.golang.org/grpc/reflection"
)

// Server is the framework instance.
type Server struct {
	grpcServer *grpcServer
	config     *Config
}

// New creates a server intstance.
func New(opts ...Option) (*Server, error) {
	cfg := createConfig(opts)

	log.Println("Create grpc server")
	grpcServer := newGrpcServer(cfg.Grpc, cfg.ServiceServers)
	reflection.Register(grpcServer.server)

	return &Server{
		grpcServer: grpcServer,
		config:     cfg,
	}, nil
}

// Serve starts gRPC and Gateway servers.
func (s *Server) Serve() {
	go func() {
		log.Fatal(s.grpcServer.Serve())
	}()
}