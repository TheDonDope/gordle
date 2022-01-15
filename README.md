# gordle

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/74929ad9fdd64d3e844a4709e1c87924)](https://app.codacy.com/gh/TheDonDope/gordle?utm_source=github.com&utm_medium=referral&utm_content=TheDonDope/gordle&utm_campaign=Badge_Grade)

A TUI implementation of the popular word quiz [Wordle](https://www.powerlanguage.co.uk/wordle/)!

## Building

- Build the cli command:

```shell
$ go build ./cmd/cli
<Empty output on build success>
```

- Build the server command:

```shell
$ go build ./cmd/server
<Empty output on build success>
```

## Running Tests

- Run the testsuite with coverage enabled:

```shell
$ go test -coverpkg=all ./... -coverprofile=coverage.out
?       github.com/TheDonDope/gordle/cmd/cli    [no test files]
?       github.com/TheDonDope/gordle/cmd/server [no test files]
?       github.com/TheDonDope/gordle/pkg/config [no test files]
?       github.com/TheDonDope/gordle/pkg/http   [no test files]
```

- Open the results in the browser:

```shell
$ go tool cover -html=coverage.out
<Opens Browser>
```

## Running with Docker

- Build the server:

```shell
$ docker-compose -p server build
Building server
Step 1/12 : FROM golang:1.17-buster AS builder
 ---> 4a581cd6feb1
Step 2/12 : LABEL maintainer="thedondope@hey.com"
 ---> Using cache
 ---> de7bdc3fcadd
Step 3/12 : WORKDIR /src
 ---> Using cache
 ---> 4ce0250120de
Step 4/12 : COPY go.mod go.sum ./
 ---> Using cache
 ---> b0ee6161c395
Step 5/12 : RUN go mod download -x
 ---> Using cache
 ---> d94c6d082b1c
Step 6/12 : COPY . ./
 ---> Using cache
 ---> 3f97a1d54c47
Step 7/12 : RUN go build -v -o /bin/server cmd/server/*.go
 ---> Using cache
 ---> 6d76a0d114ec

Step 8/12 : FROM debian:buster-slim
 ---> f49666103347
Step 9/12 : RUN set -x && apt-get update &&   DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates &&   rm -rf /var/lib/apt/lists/*
 ---> Using cache
 ---> f79505d005ac
Step 10/12 : WORKDIR /app
 ---> Using cache
 ---> a838cca6711f
Step 11/12 : COPY --from=builder /bin/server ./
 ---> Using cache
 ---> be056c36b528
Step 12/12 : CMD ["./server"]
 ---> Using cache
 ---> d64a6d294da8

Successfully built d64a6d294da8
Successfully tagged server_server:latest
```

- Run the server synchronous:

```shell
$ docker-compose -p server up
Creating gordle-server ... done
Attaching to gordle-server
gordle-server | [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
gordle-server |
gordle-server | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
gordle-server |  - using env:     export GIN_MODE=release
gordle-server |  - using code:    gin.SetMode(gin.ReleaseMode)
gordle-server |
gordle-server | [GIN-debug] GET    /healthy                  --> github.com/TheDonDope/gordle/pkg/http.Healthy (4 handlers)
gordle-server | [GIN-debug] Listening and serving HTTP on :4242
```

- Run the server as a daemonized service:

```shell
$ docker-compose -p server up -d
Starting gordle-server ... done
```