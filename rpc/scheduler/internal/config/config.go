package config

import "github.com/zeromicro/go-zero/zrpc"

type Task struct {
	// 任务名称
	TaskName string `yaml:"task_name"`
	// cron类型
	Cronspec string `yaml:"cronspec"`
}

type Config struct {
	zrpc.RpcServerConf
	// 任务列表
	Tasks []Task `yaml:"Tasks"`
}
