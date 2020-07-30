package tictactoe

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
)

func TestPlayer(t *testing.T) {

	BeforeEach(func() {

	})

	AfterEach(func() {

	})

	Describe("Player", func() {

		It("should return *Player", func() {
			ttt := CreatePlayer()
			AssertEqual(reflect.TypeOf(ttt).String(), "*tictactoe.Player")
		})

		It("should play opening move", func() {
			ttt := CreatePlayer()
			m := ttt.Move([]byte("         "))
			AssertEqual(string(m), "        X")
		})

		It("should start a new game", func() {
			ttt := CreatePlayer()
			m := ttt.Move([]byte("XOXXOO OX"))
			AssertEqual(string(m), "         ")
		})

		It("should start a new game", func() {
			ttt := CreatePlayer()
			m := ttt.Move([]byte("XOXXOO OX"))
			AssertEqual(string(m), "         ")
		})

		It("should return 'NotEnded' as it's a new game", func() {
			AssertEqual(Condition([]byte("         ")), NotEnded)
		})

		It("should return 'Illegal' as there are too many Xs", func() {
			AssertEqual(Condition([]byte("XXX      ")), Illegal)
		})

		It("should return 'Illegal' as there are too many Os", func() {
			AssertEqual(Condition([]byte("X X O  OO")), Illegal)
		})

		It("should return 'Illegal' because of invalid input", func() {
			AssertEqual(Condition([]byte("F        ")), Illegal)
		})

		It("should return 'Illegal' because of no input", func() {
			AssertEqual(Condition([]byte("")), Illegal)
		})

		It("should return 'Illegal' because of too much input", func() {
			AssertEqual(Condition([]byte("          ")), Illegal)
		})

		It("should return 'Illegal' because no move was made", func() {
			ttt := CreatePlayer()
			_, v := ttt.Play([]byte("         "), []byte("         "))
			AssertEqual(v, Illegal)
		})

		It("should return 'Illegal' because O played first", func() {
			ttt := CreatePlayer()
			_, v := ttt.Play([]byte("        O"), []byte("         "))
			AssertEqual(v, Illegal)
		})

		It("should return 'Illegal' because the board was reset", func() {
			ttt := CreatePlayer()
			_, v := ttt.Play([]byte("        X"), []byte("         "))
			AssertEqual(v, Illegal)
		})

		It("should return 'Illegal' because the board didn't change", func() {
			ttt := CreatePlayer()
			_, v := ttt.Play([]byte("        X"), []byte("        X"))
			AssertEqual(v, Illegal)
		})

		It("should return 'Illegal' because a play was removed", func() {
			ttt := CreatePlayer()
			_, v := ttt.Play([]byte("    O   X"), []byte("        X"))
			AssertEqual(v, Illegal)
		})

		It("should return 'Illegal' because two plays were removed", func() {
			ttt := CreatePlayer()
			_, v := ttt.Play([]byte(" X  O   X"), []byte("        X"))
			AssertEqual(v, Illegal)
		})

		It("should return 'Illegal' because the board was altered with no play", func() {
			ttt := CreatePlayer()
			_, v := ttt.Play([]byte("    O   X"), []byte("      O X"))
			AssertEqual(v, Illegal)
		})

		It("should return 'Illegal' because the board was altered with play", func() {
			ttt := CreatePlayer()
			_, v := ttt.Play([]byte("    O   X"), []byte("   X  O X"))
			AssertEqual(v, Illegal)
		})
	})

	Report(t)
}
