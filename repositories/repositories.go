package repositories

import "net/http"

// NewAPITweetRepository returns new intance of APITweetRepository.
func NewAPITweetRepository(apiHost, apiKey, apiSecretKey string) TweetRepositoryContract {
	return &APITweetRepository{
		apiHost:      apiHost,
		apiKey:       apiKey,
		apiSecretKey: apiSecretKey,
		client:       &http.Client{},
	}
}
