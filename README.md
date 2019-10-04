# twtx

GraphQL API for fetching tweets (For learning purpose).

### Environment

Can be seen at `.env.example`:

```
TWITTER_API_HOST=https://api.twitter.com
TWITTER_API_KEY=
TWITTER_API_SECRET_KEY=
```

### Running

Run this command to start server at port 8000:

```
twtx serve
```

GraphQL endpoint is at `/query`.

### Building

Run this command to build the Docker image:

```
docker build . -t dewadg/twtx
```
