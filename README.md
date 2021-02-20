[![Build Status](https://www.travis-ci.org/profallinson/tictactoe.svg?branch=master)](https://www.travis-ci.org/profallinson/tictactoe)

# TicTacToe Rule Engine

## Testing

	cd $GOPATH/src/github.com/profallinson/tictactoe
	go test

## Coverage

	cd $GOPATH/src/github.com/profallinson/tictactoe
	go test -covermode=count -coverprofile=count.out; go tool cover -html=count.out -o=coverage.html
