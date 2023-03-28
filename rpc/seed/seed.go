package main

import (
	"context"
	"flag"
	"fmt"

	"bluebird/rpc/seed/internal/config"
	"bluebird/rpc/seed/internal/mqs/delayMq"
	"bluebird/rpc/seed/internal/server"
	"bluebird/rpc/seed/internal/svc"
	"bluebird/rpc/seed/seed"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/seed.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	defer svcCtx.Cleanup()

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		seed.RegisterSeedServiceServer(grpcServer, server.NewSeedServiceServer(svcCtx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	//defer s.Stop()
	serviceGroup.Add(s)

	at := delayMq.NewAsynqTask(context.Background(), svcCtx)
	serviceGroup.Add(at)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	//s.Start()
	serviceGroup.Start()
}
