# gordle

[![CodeQL](https://github.com/TheDonDope/gordle/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/TheDonDope/gordle/actions/workflows/codeql-analysis.yml)

A TUI implementation of the popular word quiz [Wordle](https://www.powerlanguage.co.uk/wordle/)!

```shell
$ go run ./cmd/cli
Welcome to 🟩🟨⬛ Gordle ⬛🟨🟩
You have 6 trys to guess the word of the day.
🟩 means, the letter is in the word and in the correct spot.
🟨 means, that the letter is in the word but in the wrong spot.
⬛ means, that the letter is in not in the word in any spot.
Enter 5 characters 12345
12345 ⬛⬛⬛⬛⬛ (Try 1/6)
Enter 5 characters ALTER
ALTER 🟩🟩🟩🟩🟩 (Try 2/6)

Congratulations, you won! 🥳🥳

Your Gordle results:
12345 ⬛⬛⬛⬛⬛ (1/6)
ALTER 🟩🟩🟩🟩🟩 (2/6)

The solution was: ALTER
```

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
$ kubectl apply -f server-deployment.yaml,server-service.yaml
deployment.apps/server created
service/server created
```