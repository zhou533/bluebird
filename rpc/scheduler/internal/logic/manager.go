package logic

import (
	"bluebird/rpc/scheduler/internal/svc"
	"context"
	"time"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

type TaskManager struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskManager(ctx context.Context, svcCtx *svc.ServiceContext) *TaskManager {
	return &TaskManager{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (m *TaskManager) GetConfigs() ([]*asynq.PeriodicTaskConfig, error) {
	var configs []*asynq.PeriodicTaskConfig
	for _, task := range m.svcCtx.Config.Tasks {
		configs = append(configs, &asynq.PeriodicTaskConfig{Cronspec: task.Cronspec, Task: asynq.NewTask(task.TaskName, nil)})
	}
	return configs, nil
}

func (m *TaskManager) TManager() (*asynq.PeriodicTaskManager, error) {
	mgr, err := asynq.NewPeriodicTaskManager(
		asynq.PeriodicTaskManagerOpts{
			RedisConnOpt:               asynq.RedisClientOpt{Addr: m.svcCtx.Config.Redis.Host},
			PeriodicTaskConfigProvider: m,                // this provider object is the interface to your config source
			SyncInterval:               10 * time.Second, // this field specifies how often sync should happen
		})
	if err != nil {
		logx.WithContext(m.ctx).Errorf("failed to create periodic task manager: %v", err)
		return nil, err
	}
	return mgr, nil
}
