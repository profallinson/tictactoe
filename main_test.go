package tictactoe

import (
	"fmt"
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
			m, _ := ttt.Move([]byte("         "))
			AssertEqual(string(m), "        X")
		})

		It("should start a new game", func() {
			ttt := CreatePlayer()
			m, _ := ttt.Move([]byte("XOXXOO OX"))
			AssertEqual(string(m), "XOXXOO OX")
		})

	})

	Describe("IsLegalBoard()", func() {

		It("should return 'Illegal' as there are too many Xs", func() {
			AssertEqual(IsLegalBoard([]byte("XXX      ")), false)
		})

		It("should return 'Illegal' as there are too many Os", func() {
			AssertEqual(IsLegalBoard([]byte("X X O  OO")), false)
		})

		It("should return 'Illegal' because of invalid input", func() {
			AssertEqual(IsLegalBoard([]byte("F        ")), false)
		})

	})

	Describe("Condition()", func() {

		It("should return 'NotEnded' as it's a new game", func() {
			AssertEqual(Condition([]byte("         ")), NotEnded)
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

	Describe("Playing oneself", func() {

		It("should play its self and 'Tie' 10 times", func() {
			ttt := CreatePlayer()
			m := []byte("         ")
			x := 10000
			t := 10
			for x > 0 && t > 0 {
				m, _ = ttt.Move(m)
				if c := Condition(m); c != NotEnded {
					fmt.Println(ConditionToString(c))
					m = []byte("         ")
					if c == Tie {
						t--
					}
				}
				fmt.Println(string(m))
				x--
			}
			AssertEqual(t, 0)
		})

	})

	Report(t)
}
