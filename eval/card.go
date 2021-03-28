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
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"regexp"
	"unicode"
	//"log"
)

func init() {
	cardRE = regexp.MustCompile("^\\s*(10|[2-9]|[TJQKAtjqka])([CDHScdhs])\\s*$")
	for _, s := range suits {
		for j := 0; j < 13; j++ {
			DefaultDeck.Push(Card((1 << (16 + j)) | s | (int32(j) << 8) | primeRanks[j]))
		}
	}
}

// Card is the type representing a single laying card. It is 32 bits long, packed according to the following schematic:
// 	+--------+--------+--------+--------+
// 	|xxxbbbbb|bbbbbbbb|cdhsrrrr|xxpppppp|
// 	+--------+--------+--------+--------+
//
// 	p 	= prime number of rank (deuce=2,trey=3,four=5,...,ace=41)
// 	r 	= rank of card (deuce=0,trey=1,four=2,five=3,...,ace=12)
// 	cdhs	= suit of card (bit turned on based on suit of card)
// 	b 	= bit turned on depending on rank of card
type Card int32

// ErrBadCard is the error returned by ParseCardBytes if the passed []byte does not match the correct
// format
var ErrBadCard = errors.New("invalid card string")

// ParseCardBytes is a helper function that will convert a human-readable representation of a card to the equivalent
// 32 bit layout. It accepts strings composed of a rank concatenated with a suit. The rank may be 2-10,T,J,Q,K, or A,
// and the suit may be C, D, H, or S. It is not case sensative, and 10 and T are equivalent. It discards
// whitespace before and after. More precisely, it accepts the language L defined by:
// 	^\s*(10|[2-9]|[TJQKAtjqka])([CDHScdhs])\s*$
// Should there be a conflict between the plain english and the regular expression, the regular expression takes precedence.
// If b is not in L, ParseCardBytes returns an error value. Otherwise, it returns a Card in native format and nil
func ParseCardBytes(b []byte) (Card, error) {

	matches := cardRE.FindSubmatchIndex(b)
	if len(matches) != 6 {
		return Card(0), ErrBadCard
	}
	upc := bytes.ToUpper(b)
	rankStr := string(upc[matches[2]:matches[3]])
	suitStr := string(upc[matches[4]:matches[5]])

	rank := chrToNumRanks[rankStr]
	suit := chrToNumSuits[suitStr]
	prime := primeRanks[rank]
	bbbb := int32(0x10000) << rank
	rrrr := rank << 8

	return Card(bbbb | suit | rrrr | prime), nil
}

// MustParseCardBytes is the same as ParseCardBytes, except it panics if b is not in L
func MustParseCardBytes(b []byte) Card {
	c, err := ParseCardBytes(b)
	if err != nil {
		panic(err)
	}
	return c
}

// MustParseCardString is the same as MustParseCardBytes, except it takes a string as an arg
func MustParseCardString(s string) Card {
	c, err := ParseCardBytes([]byte(s))
	if err != nil {
		panic(err)
	}
	return c
}

// Scan wraps ParseCardBytes to satisfy the fmt.Scanner interface
// If ParseCardBytes returns an error, c is guaranteed not to change
func (c Card) Scan(state fmt.ScanState, verb rune) error {

	if verb != 'v' {
		return errors.New("custom scan formats not implemented")
	}

	b, err := state.Token(true, func(r rune) bool { return unicode.In(r, unicode.L, unicode.Nd) })
	if err != nil {
		return err
	}

	card, err := ParseCardBytes(b)
	if err != nil {
		return err
	}

	c = card
	return nil
}

// CardToString outputs a human-readable representation of a Card. Unlike Scan,
// this may be useful for creating serialized GameViews for the end user.
// The returned string will be in all uppercase, and the rank 10 will be represented using
// the letter T. Calling this method on an poorly formed card will result in undefined behavior.
func (c Card) String() string {
	rank := (int32(c) >> 8) & 0x0F
	suit := int32(c) & 0xF000

	return string(numToChrRanks[rank]) + string(numToChrSuits[suit])
}

// Deck is the basic type representing a deck of playing cards
type Deck []Card

// DefaultDeck contains all 52 cards. It may be ordered, but that is not guaranteed in the future
var DefaultDeck Deck

//TODO: Are Marshal and Unmarshal things to support long term?
// Maybe the end-user should have the freedom/responsiblity to write their own Marshaling functions

// Marshal returns an empty interface, the underlying type of which is a string, or []byte. The string is a binary
// representation of the deck, and the characters within are not guaranteed to be printable. Marshal returns an error
// if the internal write operation fails.
func (d Deck) Marshal() (interface{}, error) {
	if len(d) == 0 {
		return nil, nil
	}
	//outgoing is a binary string
	wBuf := new(bytes.Buffer)
	err := binary.Write(wBuf, binary.LittleEndian, d)
	if err != nil {
		return nil, err
	}
	return string(wBuf.Bytes()), nil
}

// Unmarshal populates a deck from a binary string previously generated by Marshal. If the empty
// interface v does not have a []byte as its underlying type, d will be set to nil.
func (d *Deck) Unmarshal(v interface{}) error {
	data, ok := v.([]byte)
	if !ok {
		*d = nil
		return nil
	}
	rBuf := bytes.NewBuffer(data)

	rDeck := make([]Card, (len(data) / 4))
	err := binary.Read(rBuf, binary.LittleEndian, &rDeck)
	if err == io.EOF {
		*d = Deck{}
	} else if err != nil {
		return err
	}
	*d = rDeck
	return nil
}

// Pop removes and returns the top card. Return 0 if stack is empty.
func (d *Deck) Pop() Card {
	if len(*d) == 0 {
		return 0
	}

	ndx := len(*d) - 1
	element := (*d)[ndx]
	*d = (*d)[:ndx]
	return element
}

// IsEmpty returns true if d has length 0, and false otherwise
func (d *Deck) IsEmpty() bool { return (len(*d) == 0) }

// Push places card onto the top of the deck. Useful for constructing known orderings, if you want.
func (d *Deck) Push(card Card) { *d = append(*d, card) }

// Shuffle resets the contents of d and performs a Fisher-Yates shuffle. Post-condition: d contains all 52 unique cards, in a normally distributed random order.
func (d *Deck) Shuffle(rand *rand.Rand) {
	*d = append([]Card{}, DefaultDeck...)
	rand.Shuffle(len(*d), func(i, j int) { (*d)[i], (*d)[j] = (*d)[j], (*d)[i] })
}

var cardRE *regexp.Regexp

var suits = [4]int32{
	0x8000, //Clubs
	0x4000, //Diamonds
	0x2000, //Hearts
	0x1000, //Spades
}

var chrToNumSuits = map[string]int32{
	"C": 0x8000, //Clubs
	"D": 0x4000, //Diamonds
	"H": 0x2000, //Hearts
	"S": 0x1000, //Spades
}
var numToChrSuits = map[int32]byte{
	0x8000: 'C', //Clubs
	0x4000: 'D', //Diamonds
	0x2000: 'H', //Hearts
	0x1000: 'S', //Spades
}

var chrToNumRanks = map[string]int32{
	"2":  0,
	"3":  1,
	"4":  2,
	"5":  3,
	"6":  4,
	"7":  5,
	"8":  6,
	"9":  7,
	"10": 8,
	"T":  8,
	"J":  9,
	"Q":  10,
	"K":  11,
	"A":  12,
}

var numToChrRanks = [13]int32{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
var primeRanks = [13]int32{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}
