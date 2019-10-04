package repositories

import (
	"testing"
)

func TestAPITweetRepositoryGet(t *testing.T) {
	repository := NewAPITweetRepository(
		"https://api.twitter.com",
		"gptC5y5XnVMMkZhqPraySpc0E",
		"0BZo81BCGx3g7z5mi8kjL8Ua3sr6tL7kD9AkDv0QglmbKIEuNh",
	)

	_, err := repository.Get("dewadg", 3)
	if err != nil {
		t.Error(err)
	}
}
