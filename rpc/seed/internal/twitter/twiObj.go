package twitter

type TwitterUser struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	UserName        string `json:"username"`
	CreatedAt       string `json:"created_at,omitempty"`
	Description     string `json:"description,omitempty"`
	Location        string `json:"location,omitempty"`
	PinnedTweetID   string `json:"pinned_tweet_id,omitempty"`
	ProfileImageURL string `json:"profile_image_url,omitempty"`
	Protected       bool   `json:"protected,omitempty"`
	URL             string `json:"url,omitempty"`
	Verified        bool   `json:"verified,omitempty"`
}

type TwitterTweet struct {
}
