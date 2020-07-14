//* Copyright (c) 2020, Alex Lewontin
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

func TestParseCardBytes(t *testing.T) {
	type args struct {
		c []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Card
		wantErr error
	}{
		{
			"Two of Clubs (2C)",
			args{c: []byte("2C")},
			98306,
			nil,
		},
		{
			"Two of Clubs (2c)",
			args{c: []byte("2c")},
			98306,
			nil,
		},
		{
			"Jack of Hearts (JH)",
			args{c: []byte("JH")},
			33564957,
			nil,
		},
		{
			"Jack of Hearts (Jh)",
			args{c: []byte("Jh")},
			33564957,
			nil,
		},
		{
			"Jack of Hearts (jH)",
			args{c: []byte("jH")},
			33564957,
			nil,
		},
		{
			"Jack of Hearts (jh)",
			args{c: []byte("jh")},
			33564957,
			nil,
		},
		{
			"Ten of Spades (10S)",
			args{c: []byte("10S")},
			16783383,
			nil,
		},
		{
			"Ten of Spades (TS)",
			args{c: []byte("TS")},
			16783383,
			nil,
		},
		{
			"Ten of Nonsense (Tx)",
			args{c: []byte("Tx")},
			0,
			ErrBadCard,
		},
		{
			"Ten of Nonsense (10x)",
			args{c: []byte("10x")},
			0,
			ErrBadCard,
		},
		{
			"Nonsense of Spades (xS)",
			args{c: []byte("xS")},
			0,
			ErrBadCard,
		},
		{
			"Nonsense (BC)",
			args{c: []byte("BC")},
			0,
			ErrBadCard,
		},
		{
			"Nonsense (BCD)",
			args{c: []byte("BCD")},
			0,
			ErrBadCard,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCardBytes(tt.args.c)
			if got != tt.want {
				t.Errorf("ParseCardBytes() = %v, want %v", got, tt.want)
				return
			}

			if err != tt.wantErr {
				t.Errorf("ParseCardBytes() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

		})
	}
}

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
