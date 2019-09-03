package server

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/robojones/cloud-lib-go/identity"
	"github.com/robojones/cloud-identity/db"
	"github.com/robojones/cloud-identity/env"
	"github.com/robojones/cloud-identity/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func NewServer(e *env.Env, db db.CockroachDB, s *service.Service) *Server {
	return &Server{
		env:     e,
		db:      db,
		service: s,
	}
}

type Server struct {
	env     *env.Env
	db      db.CockroachDB
	service *service.Service
}

func (s *Server) List/home/jonathan/cloud/cloud-service-templateenAndServe() {
	defer s.clean()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.env.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	identity.RegisterIdentityServer(grpcServer, s.service)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(errors.Wrap(err, "error during gRPC serve"))
	}
}

func (s *Server) clean() {
	if err := s.db.Close(); err != nil {
		log.Fatal(errors.Wrap(err, "error disconnecting from database"))
	}
}
