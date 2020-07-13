package eval

import (
	"fmt"
	"testing"
)

func TestDefaultDeck(t *testing.T) {

	t.Run("Len", func(t *testing.T) {
		result := len(DefaultDeck)
		if result != 52 {
			t.Errorf("\nFAIL: \nWant: %d \nGot: %d \n", 52, result)
		}
	})

	t.Run("NoZeroCards", func(t *testing.T) {
		for i, card := range DefaultDeck {
			if card == 0 {
				t.Errorf("\nFAIL: \n Card value 0 found at index %d", i)
			}
		}
	})

	tables := []struct {
		description string
		ndx         int32
		want        Card
	}{
		{
			"TopCard_AceOfSpades",
			51,
			268442665,
		},
		{
			"BottomCard_DeuceOfClubs",
			0,
			98306,
		},
	}

	for _, table := range tables {
		testname := fmt.Sprintf("%s", table.description)
		t.Run(testname, func(t *testing.T) {
			result := DefaultDeck[table.ndx]
			if result != table.want {
				t.Errorf("\nFAIL:\nIn: %d \nWant: %d  \nGot: %d \n", table.ndx, table.want, result)
			}
		})
	}
}

func TestShuffledDeck(t *testing.T) {
	var testDeck Deck
	var testDeckTwo Deck

	testDeck = DefaultDeck
	testDeckTwo = DefaultDeck
	testDeck.Shuffle()
	testDeckTwo.Shuffle()

	t.Run("Len", func(t *testing.T) {
		result := len(testDeck)
		if result != 52 {
			t.Errorf("\nFAIL: \nWant: %d \nGot: %d \n", 52, result)
		}
	})

	t.Run("NoZeroCards", func(t *testing.T) {
		for i, card := range testDeck {
			if card == 0 {
				t.Errorf("\nFAIL: \n Card value 0 found at index %d", i)
			}
		}
	})

	t.Run("NonRepeatable", func(t *testing.T) {
		var identical int = 0
		for i := range testDeck {
			if testDeck[i] == testDeckTwo[i] {
				identical = identical + 1
			}
			if identical == 52 {
				t.Errorf("\nFAIL: \n Two shuffled decks are identical.\n%v\n%v\n", testDeck, testDeckTwo)
			}
		}
	})
}

func TestPop(t *testing.T) {
	var testDeckLen Deck
	var testDeckEmpty Deck
	var testDeckRetVal Deck

	testDeckLen = DefaultDeck
	testDeckRetVal = DefaultDeck

	t.Run("LenAfter", func(t *testing.T) {
		testDeckLen.Pop()
		result := len(testDeckLen)
		if result != 51 {
			t.Errorf("\nFAIL: \nWant: %d \nGot: %d \n", 51, result)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		result := testDeckEmpty.Pop()
		if result != 0 {
			t.Errorf("\nFAIL: \nWant: %d \nGot: %d \n", 0, result)
		}
	})

	t.Run("ReturnValue", func(t *testing.T) {
		result := testDeckRetVal.Pop()
		if result != 268442665 {
			t.Errorf("\nFAIL: \nWant: %d \nGot: %d \n", 268442665, result)
		}
	})

}
