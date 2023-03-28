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

	mux.HandleFunc("update:seed:details", func(ctx context.Context, t *asynq.Task) error {
		fmt.Println("add_seed")
		return nil
	})

	if err := srv.Run(mux); err != nil {
		panic(err)
	}
}

func (a *AsynqTask) Stop() {
	fmt.Println("AsynqTask Stop")
}
