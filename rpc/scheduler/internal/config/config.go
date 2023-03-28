package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Task struct {
	// 任务名称
	TaskName string `json:"task_name"`
	// cron类型
	Cronspec string `json:"cronspec"`
}

type Config struct {
	service.ServiceConf
	Redis redis.RedisConf
	// 任务列表
	Tasks []Task
}
