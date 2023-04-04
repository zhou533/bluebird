package twitter

import (
	"bluebird/rpc/seed/internal/config"

	"github.com/g8rswimmer/go-twitter/v2"
)

type TwitterClient struct {
	Config config.Config
	Client twitter.Client
}

func NewTwitterClient(cfg config.Config) *TwitterClient {
	tc := twitter.Client{}
	return &TwitterClient{
		Config: cfg,
		Client: tc,
	}
}
