# gordle

[![CodeQL](https://github.com/TheDonDope/gordle/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/TheDonDope/gordle/actions/workflows/codeql-analysis.yml)

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

## Deploying to Kubernetes

- Build and run the server:

```shell
$ kubectl.exe apply -f server-deployment.yaml,server-service.yaml
kubectl.exe apply -f server-deployment.yaml,server-service.yaml
```