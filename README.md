# twtx

GraphQL API for fetching tweets (For learning purpose).

![Screen Recording 2019-10-05 at 00 42 33 2019-10-05 00_51_49](https://user-images.githubusercontent.com/3234011/66228978-ff465780-e70a-11e9-99ef-5b6209391e94.gif)


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
