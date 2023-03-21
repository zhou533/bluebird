package logic

import (
	"context"

	"bluebird/api/internal/svc"
	"bluebird/api/internal/types"
	"bluebird/rpc/seed/seed"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSeedLogic {
	return &AddSeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSeedLogic) AddSeed(req *types.SeedAddRequest) (r *types.SeedAddResponse, e error) {

	resp, err := l.svcCtx.SeedService.AddSeed(l.ctx, &seed.SeedAddRequest{
		ScreenName: req.ScreenName,
	})
	if err != nil {
		return nil, err
	}

	return &types.SeedAddResponse{
		Code: resp.Code,
		Msg:  resp.Msg,
	}, nil
}
