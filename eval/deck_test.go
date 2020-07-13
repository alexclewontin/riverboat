//* Copyright (c) 2020, AUTHOR, Alex Lewontin
//* All rights reserved.
//* 
//* Redistribution and use in source and binary forms, with or without
//* modification, are permitted provided that the following conditions are met:
//* 
//* - Redistributions of source code must retain the above copyright notice, this
//* list of conditions and the following disclaimer.
//* - Redistributions in binary form must reproduce the above copyright notice,
//* this list of conditions and the following disclaimer in the documentation
//* and/or other materials provided with the distribution.
//* 
//* THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
//* ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
//* WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
//* DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
//* FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
//* DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
//* SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
//* CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
//* OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
//* OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

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
