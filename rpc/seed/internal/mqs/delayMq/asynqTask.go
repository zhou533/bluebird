package delayMq

import (
	"bluebird/rpc/seed/internal/svc"
	"context"
	"fmt"

	"github.com/hibiken/asynq"
)

type AsynqTask struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAsynqTask(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqTask {
	return &AsynqTask{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (a *AsynqTask) Start() {
	fmt.Println("AsynqTask Start")

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: a.svcCtx.Config.Redis.Host},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"default": 10,
			},
		},
	)

	mux := asynq.NewServeMux()

	mux.Handle("update:seed:details", NewUpdateUserDetailsHandler(a.svcCtx))

	if err := srv.Run(mux); err != nil {
		panic(err)
	}
}

func (a *AsynqTask) Stop() {
	fmt.Println("AsynqTask Stop")
}
