// Code generated by goctl. DO NOT EDIT.
// Source: seed.proto

package seedservice

import (
	"context"

	"bluebird/rpc/seed/seed"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SeedAddRequest  = seed.SeedAddRequest
	SeedAddResponse = seed.SeedAddResponse

	SeedService interface {
		AddSeed(ctx context.Context, in *SeedAddRequest, opts ...grpc.CallOption) (*SeedAddResponse, error)
	}

	defaultSeedService struct {
		cli zrpc.Client
	}
)

func NewSeedService(cli zrpc.Client) SeedService {
	return &defaultSeedService{
		cli: cli,
	}
}

func (m *defaultSeedService) AddSeed(ctx context.Context, in *SeedAddRequest, opts ...grpc.CallOption) (*SeedAddResponse, error) {
	client := seed.NewSeedServiceClient(m.cli.Conn())
	return client.AddSeed(ctx, in, opts...)
}
