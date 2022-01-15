# Gordle

[![CodeQL](https://github.com/TheDonDope/gordle/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/TheDonDope/gordle/actions/workflows/codeql-analysis.yml)

A TUI implementation of the popular word quiz [Wordle](https://www.powerlanguage.co.uk/wordle/)!

```shell
$ go run ./cmd/cli
Welcome to 🟩🟨⬛ Gordle ⬛🟨🟩
You have 6 trys to guess the word of the day.
NOTE: The current implementation will pick a new word on every run!
🟩 means, the letter is in the word and in the correct spot.
🟨 means, that the letter is in the word but in the wrong spot.
⬛ means, that the letter is in not in the word in any spot.
Enter 5 characters ALTER
ALTER 🟨🟩⬛⬛⬛ (Try 1/6)
Enter 5 characters ELONG
ELONG ⬛🟩⬛⬛⬛ (Try 2/6)
Enter 5 characters ILLIC
ILLIC ⬛🟩🟩⬛⬛ (Try 3/6)
Enter 5 characters ULLAG
ULLAG 🟩🟩🟩🟩⬛ (Try 4/6)
Enter 5 characters ULLAD
ULLAD 🟩🟩🟩🟩⬛ (Try 5/6)
Enter 5 characters ULLAN
ULLAN 🟩🟩🟩🟩⬛ (Try 6/6)

Your Gordle results (2022-01-16):
ALTER 🟨🟩⬛⬛⬛ (1/6)
ELONG ⬛🟩⬛⬛⬛ (2/6)
ILLIC ⬛🟩🟩⬛⬛ (3/6)
ULLAG 🟩🟩🟩🟩⬛ (4/6)
ULLAD 🟩🟩🟩🟩⬛ (5/6)
ULLAN 🟩🟩🟩🟩⬛ (6/6)

The solution was: ULLAM
```

## Building

- Build the cli command:

```shell
$ go build ./cmd/cli
<Empty output on build success>
```

## Running

- Either run:

```shell
$ go run ./cmd/cli
[...]
```

- Or run this after having build the command:

```shell
$ ./cli
[...]
```

## Running Tests

- Run the testsuite with coverage enabled:

```shell
$ go test ./... -coverprofile=coverage.out
?       github.com/TheDonDope/gordle/cmd/cli    [no test files]
ok      github.com/TheDonDope/gordle/pkg/validation     0.001s  coverage: 100.0% of statements
```

- Open the results in the browser:

```shell
$ go tool cover -html=coverage.out
<Opens Browser>
```
