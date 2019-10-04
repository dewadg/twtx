package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/dewadg/twtx/gql"
	"github.com/dewadg/twtx/handlers"
	"github.com/dewadg/twtx/repositories"

	"github.com/go-chi/chi"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	switch command() {
	case "serve":
		gqlSchema := initGQL()
		router := initRouter(gqlSchema)

		fmt.Println("App running on port 8000")
		log.Fatal(http.ListenAndServe(":8000", router))
		break
	}
}

func command() string {
	args := os.Args[1:]

	if len(args) > 0 {
		return args[0]
	}
	return ""
}

func initGQL() graphql.Schema {
	twitterAPIHost := os.Getenv("TWITTER_API_HOST")
	twitterAPIKey := os.Getenv("TWITTER_API_KEY")
	twitterAPISecretKey := os.Getenv("TWITTER_API_SECRET_KEY")

	apiTweetRepository := repositories.NewAPITweetRepository(
		twitterAPIHost,
		twitterAPIKey,
		twitterAPISecretKey,
	)

	gqlRootQuery := gql.NewRootQuery(apiTweetRepository)

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: gqlRootQuery,
	})

	if err != nil {
		log.Fatal(err)
	}
	return schema
}

func initRouter(schema graphql.Schema) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.StripSlashes,
		middleware.Recoverer,
	)

	router.Post("/query", handlers.GQLHandler(schema))

	return router
}
