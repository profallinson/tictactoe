package tictactoe

import (
	// "fmt"
	"bytes"
	"math/rand"
)

type Status uint8

const (
	NotEnded Status = iota
	Illegal
	XWon
	OWon
	Tie
	X byte = 88 // "X"
	O byte = 79 // "O"
	F byte = 32 // " "
)

type Player struct{}

func CreatePlayer() *Player {
	this := &Player{}
	rand.Seed(42)
	return this
}

// Given an array of nine bytes it will return an appropriate move.
func (this *Player) Move(b []byte) []byte {

	// If the game has ended return an empty array.
	if c := Condition(b); c == Illegal || c != NotEnded {
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
		case X:
			x++
		case O:
			o++
		}
	}

	player := X // X
	if o < x {
		player = O // O
	}

	cell := rand.Intn(9)
	for board[cell] != F { // " "
		cell = rand.Intn(9)
	}
	board[cell] = player
	return board
}

func count(b []byte) (x int, o int, f int) {
	for _, p := range b {
		switch p {
		case X:
			x++
		case O:
			o++
		case F:
			f++
		}
	}
	return x, o, f
}

func IsLegalMove(last []byte, curr []byte) bool {
	if len(last) != 9 || len(curr) != 9 {
		return false
	}
	if bytes.Equal(last, curr) {
		return false
	}
	for i := 0; i < 9; i++ {
		if last[i] != curr[i] && last[i] != F {
			return false
		}
	}
	return true
}

func IsLegalBoard(b []byte) bool {
	if len(b) != 9 {
		return false
	}
	x, o, f := count(b)
	if x+o+f != 9 {
		return false
	}
	if f == 9 {
		return true
	}
	if o > x || x > o+1 {
		return false
	}
	return true
}

func Condition(board []byte) Status {
	if !IsLegalBoard(board) {
		return Illegal
	}
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
		return NotEnded
	}
}
