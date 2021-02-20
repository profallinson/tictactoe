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

		It("should return *player", func() {
			ttt := CreatePlayer()
			AssertEqual(reflect.TypeOf(ttt).String(), "*tictactoe.player")
		})

		It("should play opening move", func() {
			ttt := CreatePlayer()
			m := ttt.Move([]byte("         "))
			AssertEqual(string(m), "        X")
		})

		It("should start a new game", func() {
			ttt := CreatePlayer()
			m := ttt.Move([]byte("XOXXOO OX"))
			AssertEqual(string(m), "XOXXOO OX")
		})

	})

	Describe("IsLegalMove()", func() {

		It("should return 'true' as a opening move", func() {
			a := []byte("         ")
			b := []byte("X        ")
			AssertEqual(IsLegalMove(a, b), true)
		})

		It("should return 'true' as a second move", func() {
			a := []byte("X        ")
			b := []byte("XO       ")
			AssertEqual(IsLegalMove(a, b), true)
		})

		It("should return 'false' the boards are different sizes", func() {
			a := []byte("X       ")
			b := []byte("O        ")
			AssertEqual(IsLegalMove(a, b), false)
			a = []byte("X        ")
			b = []byte("O       ")
			AssertEqual(IsLegalMove(a, b), false)
		})

		It("should return 'false' as you can't play on a played space", func() {
			a := []byte("X        ")
			b := []byte("O        ")
			AssertEqual(IsLegalMove(a, b), false)
		})

		It("should return 'false' as there was no move", func() {
			a := []byte("X        ")
			b := []byte("X        ")
			AssertEqual(IsLegalMove(a, b), false)
		})

		It("should return 'false' as there was more than one move", func() {
			a := []byte("X        ")
			b := []byte("X  O   O ")
			AssertEqual(IsLegalMove(a, b), false)
		})

		It("should return 'false' as there was more than one move", func() {
			a := []byte("X  O   X ")
			b := []byte("X OOX  X ")
			AssertEqual(IsLegalMove(a, b), false)
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

	Describe("Playing with oneself", func() {

		It("should play itself for a 10000 moves", func() {
			ttt := CreatePlayer()
			b := []byte("         ")
			m := 10000
			g := 0
			t := 0
			x := 0
			o := 0
			for m > 0 {
				b = ttt.Move(b)
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
			fmt.Println("XWon", x)
			fmt.Println("OWon", o)
			fmt.Println("Tie ", t)
			AssertEqual(t+x+o, g)
		})

	})

	Report(t)
}
