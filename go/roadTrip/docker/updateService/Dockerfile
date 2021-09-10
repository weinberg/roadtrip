##
## Build
##
FROM golang:1.16-alpine as build

WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o /main /app/cmd/updateService/main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /

COPY --from=build /main /main

EXPOSE 9066

USER nonroot:nonroot

ENTRYPOINT ["/main"]
