package gql

import (
	"github.com/graphql-go/graphql"
)

// UserType is type of tweep.
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"screenName": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// TweetType is type of tweet.
var TweetType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Tweet",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"truncated": &graphql.Field{
			Type: graphql.Boolean,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
		"user": &graphql.Field{
			Type: UserType,
		},
	},
})
