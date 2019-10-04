package repositories

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/dewadg/twtx/models"
)

// APITweetRepository acquires tweets from Twitter API.
type APITweetRepository struct {
	apiHost      string
	apiKey       string
	apiSecretKey string
	accessToken  string
	client       *http.Client
}

type twitterAuthResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

func (repository *APITweetRepository) authenticate() error {
	credentialsString := fmt.Sprintf("%s:%s", repository.apiKey, repository.apiSecretKey)
	credentials := base64.StdEncoding.EncodeToString([]byte(credentialsString))
	url := fmt.Sprintf("%s/oauth2/token?grant_type=client_credentials", repository.apiHost)

	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", credentials))
	response, err := repository.client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var responseData twitterAuthResponse
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		return err
	}

	repository.accessToken = responseData.AccessToken
	return nil
}

// Get returns tweets.
func (repository *APITweetRepository) Get(user string, count int) ([]models.Tweet, error) {
	if repository.accessToken == "" {
		if err := repository.authenticate(); err != nil {
			return []models.Tweet{}, err
		}
	}

	requestQuery := url.Values{}
	requestQuery.Add("screen_name", user)
	requestQuery.Add("count", fmt.Sprintf("%d", count))
	requestQuery.Add("exclude_replies", "true")
	requestQuery.Add("include_rts", "false")

	url := fmt.Sprintf("%s/1.1/statuses/user_timeline.json?%s", repository.apiHost, requestQuery.Encode())
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []models.Tweet{}, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", repository.accessToken))
	response, err := repository.client.Do(request)
	if err != nil {
		return []models.Tweet{}, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []models.Tweet{}, err
	}

	payload := make([]models.Tweet, 0)
	err = json.Unmarshal(responseBody, &payload)
	if err != nil {
		return []models.Tweet{}, err
	}
	return payload, nil
}

// Add saves new tweet.
func (repository *APITweetRepository) Add(models.Tweet) (models.Tweet, error) {
	return models.Tweet{}, errors.New("Tweeting is currently disabled")
}
