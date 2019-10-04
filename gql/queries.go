package gql

import (
	"errors"

	"github.com/dewadg/twtx/repositories"
	"github.com/graphql-go/graphql"
)

func tweetListQuery(tweetRepository repositories.TweetRepositoryContract) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(TweetType),
		Args: graphql.FieldConfigArgument{
			"user": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"count": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			user, ok := params.Args["user"].(string)
			if !ok {
				return nil, errors.New("Invalid value for `user`")
			}

			count, ok := params.Args["count"].(int)
			if !ok {
				return nil, errors.New("Invalid value for `count`")
			}

			return tweetRepository.Get(user, count)
		},
	}
}

func latestTweetQuery(tweetRepository repositories.TweetRepositoryContract) *graphql.Field {
	return &graphql.Field{
		Type: TweetType,
		Args: graphql.FieldConfigArgument{
			"user": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			count := 1
			user, ok := params.Args["user"].(string)
			if !ok {
				return nil, errors.New("Invalid value for `user`")
			}

			tweets, err := tweetRepository.Get(user, count)
			if err != nil {
				return nil, err
			}
			if len(tweets) == 0 {
				return nil, nil
			}
			return tweets[0], nil
		},
	}
}
