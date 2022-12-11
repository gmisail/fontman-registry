FROM golang:1.18

WORKDIR /go/src/fontman

COPY . .

# build binary
RUN go build -o ./bin/fontman-registry ./cmd/fontman

ARG DATABASE_URL

# run database migrations
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN cd migrations && goose postgres $DATABASE_URL up

ENTRYPOINT ["./bin/fontman-registry"]
