//+build wireinjects

package main

import (
	"github.com/google/wire"
	"github.com/robojones/cloud-identity/db"
	"github.com/robojones/cloud-identity/env"
	"github.com/robojones/cloud-identity/server"
	"github.com/robojones/cloud-identity/service"
)

func InitServer() *server.Server {
	wire.Build(
		env.NewEnv,
		db.NewCockroachDB,
		service.NewService,
		server.NewServer,
	)

	return nil
}
