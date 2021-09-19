# syntax = docker/dockerfile:1-experimental
##
## Build
##
FROM golang:1.16-alpine as build

WORKDIR /app
VOLUME /cache

# Copy and download dependency using go mod
COPY go/roadTrip/go.mod .
COPY go/roadTrip/go.sum .
RUN go mod download

# Copy the code into the container
COPY go/roadTrip .

# Build the application
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOCACHE=/cache

RUN --mount=type=cache,target=/cache go build -o /main /app/cmd/mapServer/main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /

COPY --from=build /main /main

EXPOSE 9067

USER nonroot:nonroot

ENTRYPOINT ["/main"]
