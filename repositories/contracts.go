package repositories

import "github.com/dewadg/twtx/models"

// TweetRepositoryContract represents contract
// for fetching tweets.
type TweetRepositoryContract interface {
	Get(user string, count int) ([]models.Tweet, error)
	Add(tweet models.Tweet) (models.Tweet, error)
}
