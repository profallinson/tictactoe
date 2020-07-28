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

		It("should return string of moves", func() {

		})
	})

	Report(t)
}
