# Gordle

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
Enter 5 characters ALTEN
ALTEN 🟩🟩🟩🟩⬛ (Try 2/6)
Enter 5 characters ATLAS
ATLAS 🟩🟨🟨🟨⬛ (Try 3/6)
Enter 5 characters RASEN
RASEN 🟨🟨⬛🟩⬛ (Try 4/6)
Enter 5 characters NORMS
NORMS ⬛⬛🟨⬛⬛ (Try 5/6)
Enter 5 characters TANGY
TANGY 🟨🟨⬛⬛⬛ (Try 6/6)

Your Gordle results:
12345 ⬛⬛⬛⬛⬛ (1/6)
ALTEN 🟩🟩🟩🟩⬛ (2/6)
ATLAS 🟩🟨🟨🟨⬛ (3/6)
RASEN 🟨🟨⬛🟩⬛ (4/6)
NORMS ⬛⬛🟨⬛⬛ (5/6)
TANGY 🟨🟨⬛⬛⬛ (6/6)

The solution was: ALTER
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
$ go test -coverpkg=all ./... -coverprofile=coverage.out
?       github.com/TheDonDope/gordle/cmd/cli    [no test files]
ok      github.com/TheDonDope/gordle/pkg/validation     0.164s  coverage: 13.0% of statements in all
```

- Open the results in the browser:

```shell
$ go tool cover -html=coverage.out
<Opens Browser>
```
