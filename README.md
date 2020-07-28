# Tic-Tac-Toe Rule Engine

## Testing

	cd $GOPATH/src/github.com/profallinson/tictactoe
	go test

## Coverage

	cd $GOPATH/src/github.com/profallinson/tictactoe
	go test -covermode=count -coverprofile=count.out; go tool cover -html=count.out
