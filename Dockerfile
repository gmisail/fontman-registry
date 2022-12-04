FROM golang:1.18

WORKDIR /go/src/fontman

COPY . .
RUN go build -o ./bin/fontman-registry ./cmd/fontman

EXPOSE 8080

ENTRYPOINT ["./bin/fontman-registry"]
