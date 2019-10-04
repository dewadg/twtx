package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dewadg/twtx/gql"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

type gqlRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

// GQLHandler acts as HTTP handler for GQL requests.
func GQLHandler(schema graphql.Schema) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Body == nil {
			http.Error(writer, "No query given", 400)
			return
		}

		var payload gqlRequest
		err := json.NewDecoder(request.Body).Decode(&payload)
		if err != nil {
			http.Error(writer, "Error parsing query", 400)
			return
		}

		ctx := context.Background()
		result := gql.ExecuteQuery(ctx, payload.Query, payload.Variables, schema)
		render.JSON(writer, request, result)
	}
}
