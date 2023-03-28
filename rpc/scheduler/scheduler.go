package main

import (
	"bluebird/rpc/scheduler/internal/config"
	"bluebird/rpc/scheduler/internal/logic"
	"bluebird/rpc/scheduler/internal/svc"
	"context"
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/scheduler.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		panic(err)
	}
	svcCtx := svc.NewServiceContext(c)
	taskMgr := logic.NewTaskManager(context.Background(), svcCtx)
	mgr, err := taskMgr.TManager()
	if err != nil {
		panic(err)
	}

	if err := mgr.Run(); err != nil {
		logx.WithContext(context.Background()).Errorf("failed to run periodic task manager: %v", err)
	}
}
