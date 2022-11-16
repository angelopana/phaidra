# syntax = docker/dockerfile:1

# Build Stage
FROM golang:1.19-alpine as build-env

ENV APP_NAME phaidra
ENV CMD_PATH app/main.go

WORKDIR /go/src/
COPY . .

RUN ls ./

RUN go mod download

# build binary
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME ./$CMD_PATH


# Deploy
FROM alpine:3.14

ENV APP_NAME phaidra

RUN apk add bash
COPY --from=build-env /$APP_NAME /$APP_NAME

EXPOSE 8080
EXPOSE 9095

CMD ["./phaidra"]