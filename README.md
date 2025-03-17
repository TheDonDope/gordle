# Gordle

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/d52553311a5a4fc69d04032c77791cfd)](https://app.codacy.com/gh/TheDonDope/gordle/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![Codecov Badge](https://codecov.io/gh/TheDonDope/gordle/graph/badge.svg?token=DM0KH9IJLG)](https://codecov.io/gh/TheDonDope/gordle)

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

- Build the cli command (alternatively `$ make build`):

```shell
$ go build -v -o ./bin/gordle ./cmd/cli/main.go
<Empty output on build success>
```

## Running

- Either run (alternatively `$ make`):

```shell
$ go run ./cmd/cli
[...]
```

- Or run this after having build the command:

```shell
$ ./bin/gordle
[...]
```

## Running Tests

- Run the testsuite with coverage enabled (alternatively `$ make test`):

```shell
$ go test -race ./... -coverprofile cp.out
?       github.com/TheDonDope/gordle/cmd/cli    [no test files]
ok      github.com/TheDonDope/gordle/pkg/game   0.584s  coverage: 63.8% of statements
ok      github.com/TheDonDope/gordle/pkg/storage        0.019s  coverage: 57.1% of statements
```

- Open the results in the browser:

```shell
$ go tool cover -html cp.out -o cp.html
<Opens Browser>
```
