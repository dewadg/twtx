FROM golang AS build

WORKDIR /go/src/github.com/dewadg/twtx

ADD . .

RUN go get -v ./...
RUN CGO_ENABLED=0 go build

FROM alpine

WORKDIR /usr/local/bin

COPY --from=build /go/src/github.com/dewadg/twtx/twtx .
RUN chmod +x . twtx

CMD ["twtx", "serve"]
