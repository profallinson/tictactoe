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

		It("should play its self for a 1000 moves", func() {
			ttt := CreatePlayer()
			b := []byte("         ")
			m := 10000
			g := 0
			t := 0
			x := 0
			o := 0
			for m > 0 {
				b, _ = ttt.Move(b)
				// fmt.Println(string(b))

				switch Condition(b) {
				case XWon:
					b = []byte("         ")
					x++
					g++
				case OWon:
					b = []byte("         ")
					o++
					g++
				case Tie:
					b = []byte("         ")
					t++
					g++
				}

				m--
			}
			fmt.Println("Tie ", t)
			fmt.Println("XWon", x)
			fmt.Println("OWon", o)
			AssertEqual(t+x+o, g)
		})

	})

	Report(t)
}
