package config

import "github.com/zeromicro/go-zero/zrpc"

type TwitterConfig struct {
	ApiKey       string `json:"APIKey"`
	ApiSecretKey string `json:"APISecretKey"`
}

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	Twitter    TwitterConfig
}
