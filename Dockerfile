# syntax = docker/dockerfile:1


# Build our microservice
FROM golang:latest-alpine
WORKDIR /app
COPY . .
RUN go mod download