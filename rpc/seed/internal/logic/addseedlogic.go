package logic

import (
	"context"

	"bluebird/rpc/seed/internal/svc"
	"bluebird/rpc/seed/seed"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSeedLogic {
	return &AddSeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddSeedLogic) AddSeed(in *seed.SeedAddRequest) (*seed.SeedAddResponse, error) {
	// todo: add your logic here and delete this line

	return &seed.SeedAddResponse{}, nil
}
