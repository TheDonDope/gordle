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
alter ⬛⬛⬛⬛⬛ (Try 1/6)
Enter 5 characters lists
lists ⬛⬛🟨⬛🟩 (Try 2/6)
Enter 5 characters softs
softs 🟨⬛🟩⬛🟩 (Try 3/6)
Enter 5 characters usfus
usfus 🟨🟨🟩🟨🟩 (Try 4/6)
Enter 5 characters sufss
sufss 🟨🟩🟩🟨🟩 (Try 5/6)
Enter 5 characters rufus
rufus ⬛🟩🟩🟨🟩 (Try 6/6)

Your Gordle results (2022-01-17):
alter ⬛⬛⬛⬛⬛ (1/6)
lists ⬛⬛🟨⬛🟩 (2/6)
softs 🟨⬛🟩⬛🟩 (3/6)
usfus 🟨🟨🟩🟨🟩 (4/6)
sufss 🟨🟩🟩🟨🟩 (5/6)
rufus ⬛🟩🟩🟨🟩 (6/6)

The solution was: muffs
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
ok      github.com/TheDonDope/gordle/pkg/guessing       0.129s  coverage: 55.3% of statements
ok      github.com/TheDonDope/gordle/pkg/storage        0.007s  coverage: 81.0% of statements
```

- Open the results in the browser:

```shell
$ go tool cover -html=coverage.out
<Opens Browser>
```
