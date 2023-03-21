package svc

import (
	"bluebird/api/internal/config"
	"bluebird/rpc/seed/seedservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	SeedService seedservice.SeedService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		SeedService: seedservice.NewSeedService(zrpc.MustNewClient(c.Seed)),
	}
}
