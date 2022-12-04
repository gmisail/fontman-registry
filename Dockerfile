FROM golang:1.18

WORKDIR /go/src/fontman

COPY . .

# build binary
RUN go build -o ./bin/fontman-registry ./cmd/fontman

# run database migrations
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN goose sqlite3 ./data/registry.db up 

EXPOSE 8080

ENTRYPOINT ["./bin/fontman-registry"]
