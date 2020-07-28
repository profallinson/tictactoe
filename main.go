package tictactoe

import (
	// "fmt"
	"math/rand"
	"strings"
)

type Condition uint8

const (
	NotEnd Condition = iota
	XWon
	OWon
	Tie
	X = "X" // 88
	O = "O" // 79
	F = " " // 32
)

const (
	NEG = "1"
	NIL = "2"
	POS = "3"
)

type TicTacToe struct {
	board []string
	moves []string
	view  string
}

func Create() *TicTacToe {
	this := &TicTacToe{}
	this.moves = []string{}
	this.view = X
	rand.Seed(42)
	return this
}

func (this *TicTacToe) GenerateGames(c int) {
	for i := 0; i < c; i++ {
		this.GenerateGame()
		switch this.view {
		case X:
			this.view = O
		case O:
			this.view = X
		}
	}
}

func (this *TicTacToe) GenerateGame() {
	this.board = []string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
	player := X
	this.storeMove(player)
	for this.condition(this.board) == NotEnd {
		this.GenerateMove(player)
		switch player {
		case X:
			player = O
		case O:
			player = X
		}
	}
}

func (this *TicTacToe) GenerateMove(p string) {
	if this.condition(this.board) != NotEnd {
		return
	}
	// Find place to move for player.
	cell := rand.Intn(9)
	for this.board[cell] != F {
		cell = rand.Intn(9)
	}
	this.board[cell] = p
	this.storeMove(p)
}

func FuckIt(b []byte) []string {
	s := []string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
	for i := range b {
		s[i] = string(b[i])
	}
	return s
}

// Given an array of nine bytes it will return an appropriate move.
func (this *TicTacToe) RandomPlayer(b []byte) []byte {
	// If the game has ended return an empty array.
	if this.condition(FuckIt(b)) != NotEnd {
		return []byte("         ")
	}

	// We don't want to change the source data so copy it.
	board := make([]byte, 9)
	copy(board, b)

	// See who's move it is (X always goes first).
	x := 0
	o := 0
	for _, c := range board {
		switch c {
		case 88:
			x++
		case 79:
			o++
		}
	}

	player := byte(88) // X
	if o < x {
		player = 79 // O
	}

	cell := rand.Intn(9)
	for board[cell] != 32 { // " "
		cell = rand.Intn(9)
	}
	board[cell] = player
	return board
}

func (this *TicTacToe) ReplaceBoard(b []string, p string) {
	this.board = b
	this.storeMove(p)
}

func (this *TicTacToe) storeMove(p string) {
	this.moves = append(this.moves, this.conditionToInt()+strings.Join([]string(this.board), ""))
}

func (this *TicTacToe) conditionToInt() string {
	if this.view == X {
		switch this.condition(this.board) {
		case XWon:
			return POS
		case OWon:
			return NEG
		case Tie:
			return POS
		case NotEnd:
			return POS
		}
	} else {
		switch this.condition(this.board) {
		case XWon:
			return NEG
		case OWon:
			return POS
		case Tie:
			return POS
		case NotEnd:
			return POS
		}
	}

	return "1"
}

func (this *TicTacToe) GetMoves() []string {
	return this.moves
}

func (this *TicTacToe) GetMovesAsString() string {
	return strings.Join(this.moves, "\n")
}

func (this *TicTacToe) condition(board []string) Condition {
	var (
		x = (board[0] == X && board[1] == X && board[2] == X) || // Check all rows.
			(board[3] == X && board[4] == X && board[5] == X) ||
			(board[6] == X && board[7] == X && board[8] == X) ||

			(board[0] == X && board[3] == X && board[6] == X) || // Check all columns.
			(board[1] == X && board[4] == X && board[7] == X) ||
			(board[2] == X && board[5] == X && board[8] == X) ||

			(board[0] == X && board[4] == X && board[8] == X) || // Check all diagonals.
			(board[2] == X && board[4] == X && board[6] == X)

		o = (board[0] == O && board[1] == O && board[2] == O) || // Check all rows.
			(board[3] == O && board[4] == O && board[5] == O) ||
			(board[6] == O && board[7] == O && board[8] == O) ||

			(board[0] == O && board[3] == O && board[6] == O) || // Check all columns.
			(board[1] == O && board[4] == O && board[7] == O) ||
			(board[2] == O && board[5] == O && board[8] == O) ||

			(board[0] == O && board[4] == O && board[8] == O) || // Check all diagonals.
			(board[2] == O && board[4] == O && board[6] == O)

		freeCellsLeft = board[0] == F || board[1] == F || board[2] == F ||
			board[3] == F || board[4] == F || board[5] == F ||
			board[6] == F || board[7] == F || board[8] == F
	)

	switch {
	case x && !o:
		return XWon
	case o && !x:
		return OWon
	case !freeCellsLeft:
		return Tie
	default:
		return NotEnd
	}
}
