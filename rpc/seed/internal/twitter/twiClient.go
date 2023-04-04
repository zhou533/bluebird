package twitter

import (
	"bluebird/rpc/seed/internal/config"
	"net/http"

	"github.com/g8rswimmer/go-twitter/v2"
)

type TwitterClient struct {
	Client twitter.Client
}

func NewTwitterClient(cfg config.Config) (*TwitterClient, error) {
	auth, err := NewTwitterAuthorizer(cfg)
	if err != nil {
		return nil, err
	}

	tc := twitter.Client{
		Authorizer: auth,
		Client:     http.DefaultClient,
		Host:       "https://api.twitter.com",
	}
	return &TwitterClient{
		Client: tc,
	}, nil
}

func (tc *TwitterClient) LookupUseydr(usernames []string) ([]TwitterUser, error) {
	// resp, err := tc.Client.UserNameLookup(context.Background(), usernames)
	return nil, nil
}
