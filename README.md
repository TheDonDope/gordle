# Gordle

[![CodeQL](https://github.com/TheDonDope/gordle/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/TheDonDope/gordle/actions/workflows/codeql-analysis.yml) [![codecov](https://codecov.io/gh/TheDonDope/gordle/branch/develop/graph/badge.svg?token=DM0KH9IJLG)](https://codecov.io/gh/TheDonDope/gordle)

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
⬛🟨🟨🟨🟨 (Try 1/6): alter
Enter 5 characters rutel
🟨⬛🟨🟨🟨 (Try 2/6): rutel
Enter 5 characters rolet
🟨⬛🟨🟨🟨 (Try 3/6): rolet
Enter 5 characters toler
🟨⬛🟨🟨🟨 (Try 4/6): toler
Enter 5 characters mulls
⬛⬛🟨🟨🟨 (Try 5/6): mulls
Enter 5 characters mills
⬛🟩🟨🟨🟨 (Try 6/6): mills

Your Gordle results (2022-01-18):
⬛🟨🟨🟨🟨 (1/6): alter
🟨⬛🟨🟨🟨 (2/6): rutel
🟨⬛🟨🟨🟨 (3/6): rolet
🟨⬛🟨🟨🟨 (4/6): toler
⬛⬛🟨🟨🟨 (5/6): mulls
⬛🟩🟨🟨🟨 (6/6): mills

The solution was: lister
```

## Building

- Build the cli command (alternatively `$ task build` if you are using [Task](https://taskfile.dev/#/)):

```shell
$ go build ./cmd/cli
<Empty output on build success>
```

## Running

- Either run (alternatively `$ task run` if you are using [Task](https://taskfile.dev/#/)):

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

- Run the testsuite with coverage enabled (alternatively `$ task test` if you are using [Task](https://taskfile.dev/#/)):

```shell
$ go test ./... -coverprofile=coverage.out
?       github.com/TheDonDope/gordle/cmd/cli    [no test files]
ok      github.com/TheDonDope/gordle/pkg/guessing       0.125s  coverage: 63.8% of statements
ok      github.com/TheDonDope/gordle/pkg/storage        0.001s  coverage: 33.3% of statements
```

- Open the results in the browser:

```shell
$ go tool cover -html=coverage.out
<Opens Browser>
```
