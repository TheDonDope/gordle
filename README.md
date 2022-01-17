# Gordle

[![CodeQL](https://github.com/TheDonDope/gordle/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/TheDonDope/gordle/actions/workflows/codeql-analysis.yml)

A golang TUI implementation of the popular word quiz [Wordle](https://www.powerlanguage.co.uk/wordle/)!

## System requirements

A system dictionary must be installed. On debian based systems run `$ apt-get install wbritish` or `$ apt-get install wamerican`.

```shell
$ go run ./cmd/cli
Welcome to 🟩🟨⬛ Gordle ⬛🟨🟩
You have 6 trys to guess the word of the day.
NOTE: The current implementation will pick a new word on every run!
🟩 means, the letter is in the word and in the correct spot.
🟨 means, that the letter is in the word but in the wrong spot.
⬛ means, that the letter is in not in the word in any spot.
Enter 5 characters alter
alter ⬛⬛🟩🟩⬛ (Try 1/6)
Enter 5 characters tests
tests 🟨🟨🟨🟨🟩 (Try 2/6)
Enter 5 characters setts
setts 🟨🟨🟩🟨🟩 (Try 3/6)
Enter 5 characters estes
estes 🟨🟨🟩🟩🟩 (Try 4/6)
Enter 5 characters setes
setes 🟨🟨🟩🟩🟩 (Try 5/6)
Enter 5 characters mates
mates ⬛⬛🟩🟩🟩 (Try 6/6)

Your Gordle results (2022-01-16):
alter ⬛⬛🟩🟩⬛ (1/6)
tests 🟨🟨🟨🟨🟩 (2/6)
setts 🟨🟨🟩🟨🟩 (3/6)
estes 🟨🟨🟩🟩🟩 (4/6)
setes 🟨🟨🟩🟩🟩 (5/6)
mates ⬛⬛🟩🟩🟩 (6/6)

The solution was: dotes
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
ok      github.com/TheDonDope/gordle/pkg/guessing       0.007s  coverage: 92.5% of statements
```

- Open the results in the browser:

```shell
$ go tool cover -html=coverage.out
<Opens Browser>
```
