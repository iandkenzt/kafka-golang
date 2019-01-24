FROM golang:1.9.2-alpine3.6 AS builder

# Install tools required to build the project
# We will need to run `docker build --no-cache .` to update those dependencies
RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep

# Gopkg.toml and Gopkg.lock lists project dependencies
# These layers will only be re-built when Gopkg files are updated
COPY Gopkg.lock Gopkg.toml /go/src/github.com/iandkenzt/kafka-golang/
WORKDIR /go/src/github.com/iandkenzt/kafka-golang

# Install library dependencies
RUN dep ensure -vendor-only -v

# Copy all project and build it
# This layer will be rebuilt when ever a file has changed in the project directory
COPY . /go/src/github.com/iandkenzt/kafka-golang/
RUN env CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/kafka-golang -v

########### Second Stage, Get CA-Certificate And TimeZone Info ##############
FROM alpine:latest as alpine

RUN apk --no-cache add tzdata ca-certificates

COPY --from=builder /go/src/github.com/iandkenzt/kafka-golang/build/kafka-golang /kafka-golang
COPY --from=builder /go/src/github.com/iandkenzt/kafka-golang/.env /.env

# tell how to run this container 
CMD ["./kafka-golang"]

EXPOSE 3000