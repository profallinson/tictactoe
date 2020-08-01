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

	Describe("Player.Move()", func() {

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

		It("should return 'Tie'", func() {
			AssertEqual(Condition([]byte("XOXXOOOXX")), Tie)
		})

	})

	Report(t)
}
