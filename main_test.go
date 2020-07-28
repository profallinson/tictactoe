package tictactoe

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
)

func TestTicTacToe(t *testing.T) {

	BeforeEach(func() {

	})

	AfterEach(func() {

	})

	Describe("TicTacToe", func() {

		It("should return *TicTacToe", func() {
			g := Create()
			AssertEqual(reflect.TypeOf(g).String(), "*tictactoe.TicTacToe")
		})

		It("should return 10 moves", func() {
			g := Create()
			g.GenerateGame()
			AssertEqual(len(g.GetMoves()), 10)
		})

		It("should return string of moves", func() {
			g := Create()
			g.GenerateGame()
			AssertEqual(len(g.GetMovesAsString()), 109)
		})
	})

	Report(t)
}
