FROM golang:1.18

WORKDIR /go/src/fontman

COPY . .

# build binary
RUN go build -o ./bin/fontman-registry ./cmd/fontman

# run database migrations
# RUN go install github.com/pressly/goose/v3/cmd/goose@latest
# RUN cd migrations && goose sqlite3 ../data/registry.db up

ENTRYPOINT ["./bin/fontman-registry"]
