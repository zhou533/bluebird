package twitter

import (
	"bluebird/rpc/seed/internal/config"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const OAuth2TokenEndpoint = "https://api.twitter.com/oauth2/token"

type TwitterAuthorizer struct {
	Token string
}

type OAuth2TokenResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

func NewTwitterAuthorizer(cfg config.Config) (*TwitterAuthorizer, error) {

	// TODO get token from redis

	// if token is empty, request it from twitter and save it into redis
	token, err := requestToken(cfg.Twitter.ApiKey, cfg.Twitter.ApiSecretKey)
	if err != nil {
		return nil, err
	}

	return &TwitterAuthorizer{
		Token: token,
	}, nil
}

func requestToken(apiKey, apiKeySecret string) (string, error) {
	uv := url.Values{}
	uv.Add("grant_type", "client_credentials")
	body := strings.NewReader(uv.Encode())

	req, err := http.NewRequest(http.MethodPost, OAuth2TokenEndpoint, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.SetBasicAuth(apiKey, apiKeySecret)

	c := http.DefaultClient
	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	raw := &OAuth2TokenResponse{}
	if err := decoder.Decode(raw); err != nil {
		return "", err
	}

	return raw.AccessToken, nil
}

func (ta *TwitterAuthorizer) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ta.Token))
}
