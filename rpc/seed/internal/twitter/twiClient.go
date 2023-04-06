package twitter

import (
	"bluebird/rpc/seed/internal/config"
	"context"
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

func (tc *TwitterClient) LookupUser(usernames []string) ([]*TwitterUser, error) {
	opts := twitter.UserLookupOpts{
		Expansions: []twitter.Expansion{twitter.ExpansionPinnedTweetID},
		UserFields: []twitter.UserField{
			twitter.UserFieldID,
			twitter.UserFieldName,
			twitter.UserFieldUserName,
			twitter.UserFieldCreatedAt,
			twitter.UserFieldDescription,
			twitter.UserFieldLocation,
			twitter.UserFieldPinnedTweetID,
			twitter.UserFieldProfileImageURL,
			twitter.UserFieldProtected,
			twitter.UserFieldURL,
			twitter.UserFieldVerified,
		},
	}
	resp, err := tc.Client.UserNameLookup(context.Background(), usernames, opts)
	if err != nil {
		return nil, err
	}

	tUsers := make([]*TwitterUser, 0, len(resp.Raw.Users))
	for _, user := range resp.Raw.Users {
		// do something with user
		tu := &TwitterUser{
			ID:              user.ID,
			Name:            user.Name,
			UserName:        user.UserName,
			CreatedAt:       user.CreatedAt,
			Description:     user.Description,
			Location:        user.Location,
			PinnedTweetID:   user.PinnedTweetID,
			ProfileImageURL: user.ProfileImageURL,
			Protected:       user.Protected,
			URL:             user.URL,
			Verified:        user.Verified,
		}
		tUsers = append(tUsers, tu)
	}
	return tUsers, nil
}
