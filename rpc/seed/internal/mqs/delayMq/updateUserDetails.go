package delayMq

import (
	"bluebird/rpc/seed/internal/svc"
	"context"

	"github.com/hibiken/asynq"
)

type UpdateUserDetailsHandler struct {
	svcCtx *svc.ServiceContext
}

func NewUpdateUserDetailsHandler(svcCtx *svc.ServiceContext) *UpdateUserDetailsHandler {
	return &UpdateUserDetailsHandler{
		svcCtx: svcCtx,
	}
}

func (h *UpdateUserDetailsHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {
	// TODO: implement the business logic of the task
	return nil
}
