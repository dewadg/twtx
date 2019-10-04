package gql

import (
	"context"

	"github.com/dewadg/twtx/repositories"
	"github.com/graphql-go/graphql"
)

// ExecuteQuery executes query.
func ExecuteQuery(ctx context.Context, query string, variables map[string]interface{}, schema graphql.Schema) *graphql.Result {
	return graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  query,
		VariableValues: variables,
		Context:        ctx,
	})
}

// NewRootQuery returns new instance of queries.
func NewRootQuery(tweetRepository repositories.TweetRepositoryContract) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Queries",
		Fields: graphql.Fields{
			"tweets":      tweetListQuery(tweetRepository),
			"latestTweet": latestTweetQuery(tweetRepository),
		},
	})
}
