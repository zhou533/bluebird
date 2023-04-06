package svc

import (
	"bluebird/rpc/seed/internal/config"
	"bluebird/rpc/seed/internal/twitter"
	"bluebird/rpc/seed/model"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config        config.Config
	Cleanup       func()
	TwitterClient *twitter.TwitterClient
	SeedModel     model.SeedModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	log := logx.WithContext(context.Background())
	repo, cleanup, err := model.NewRepository(c.DataSource, log)
	if err != nil {
		panic(err)
	}

	twiClient, err := twitter.NewTwitterClient(c)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		Cleanup: func() {
			cleanup()
		},
		TwitterClient: twiClient,
		SeedModel:     model.NewSeedModel(repo, log),
	}
}
