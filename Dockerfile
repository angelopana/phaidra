# syntax = docker/dockerfile:1


# Build our microservice
FROM golang:1.19-alpine
WORKDIR /go/src
COPY . .
EXPOSE 9095
RUN go mod download 
CMD ["go build go/src/app/main.go", "go run /go/src/app/main.go"]