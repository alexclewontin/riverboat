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

//* This file incorporates modified benchmark data from https://github.com/chehsunliu/poker
//* covered by the following license:

// * MIT License
// *
// * Copyright (c) 2018 Che-Hsun Liu
// *
// * Permission is hereby granted, free of charge, to any person obtaining a copy
// * of this software and associated documentation files (the "Software"), to deal
// * in the Software without restriction, including without limitation the rights
// * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// * copies of the Software, and to permit persons to whom the Software is
// * furnished to do so, subject to the following conditions:
// *
// * The above copyright notice and this permission notice shall be included in all
// * copies or substantial portions of the Software.
// *
// * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// * SOFTWARE.

package eval

import (
	"encoding/json"
	"testing"

	"github.com/chehsunliu/poker"
	"github.com/notnil/joker/hand"
)

var dataPoker5 = []string{
	`["As", "Ks", "Jc", "7h", "5d"]`, // high card
	`["As", "Ac", "Jc", "7h", "5d"]`, // pair
	`["As", "Ac", "Jc", "Jd", "5d"]`, // two pair
	`["As", "Ac", "Ad", "Jd", "5d"]`, // three of a kind
	`["As", "Ks", "Qd", "Jh", "Td"]`, // straight
	`["Ts", "7s", "4s", "3s", "2s"]`, // flush
	`["4s", "4c", "4d", "2s", "2h"]`, // full house
	`["As", "Ac", "Ad", "Ah", "5h"]`, // four of a kind
	`["As", "Ks", "Qs", "Js", "Ts"]`, // straight flush
}

var dataPoker6 = []string{
	`["3d", "As", "Ks", "Jc", "7h", "5d"]`, // high card
	`["3d", "As", "Ac", "Jc", "7h", "5d"]`, // pair
	`["3d", "As", "Ac", "Jc", "Jd", "5d"]`, // two pair
	`["3d", "As", "Ac", "Ad", "Jd", "5d"]`, // three of a kind
	`["3d", "As", "Ks", "Qd", "Jh", "Td"]`, // straight
	`["3d", "Ts", "7s", "4s", "3s", "2s"]`, // flush
	`["3d", "4s", "4c", "4d", "2s", "2h"]`, // full house
	`["3d", "As", "Ac", "Ad", "Ah", "5h"]`, // four of a kind
	`["3d", "As", "Ks", "Qs", "Js", "Ts"]`, // straight flush
}

var dataPoker7 = []string{
	`["2d", "3d", "As", "Ks", "Jc", "7h", "5d"]`, // high card
	`["2d", "3d", "As", "Ac", "Jc", "7h", "5d"]`, // pair
	`["2d", "3d", "As", "Ac", "Jc", "Jd", "5d"]`, // two pair
	`["2c", "3d", "As", "Ac", "Ad", "Jd", "5d"]`, // three of a kind
	`["2d", "3d", "As", "Ks", "Qd", "Jh", "Td"]`, // straight
	`["2d", "3d", "Ts", "7s", "4s", "3s", "2s"]`, // flush
	`["2d", "3d", "4s", "4c", "4d", "2s", "2h"]`, // full house
	`["2d", "3d", "As", "Ac", "Ad", "Ah", "5h"]`, // four of a kind
	`["2d", "3d", "As", "Ks", "Qs", "Js", "Ts"]`, // straight flush
}

var dataJoker5 = []string{
	`["A♠", "K♠", "J♣", "7♥", "5♦"]`, // high card
	`["A♠", "A♣", "J♣", "7♥", "5♦"]`, // pair
	`["A♠", "A♣", "J♣", "J♦", "5♦"]`, // two pair
	`["A♠", "A♣", "A♦", "J♦", "5♦"]`, // three of a kind
	`["A♠", "K♠", "Q♦", "J♥", "T♦"]`, // straight
	`["T♠", "7♠", "4♠", "3♠", "2♠"]`, // flush
	`["4♠", "4♣", "4♦", "2♠", "2♥"]`, // full house
	`["A♠", "A♣", "A♦", "A♥", "5♥"]`, // four of a kind
	`["A♠", "K♠", "Q♠", "J♠", "T♠"]`, // straight flush
}

var dataJoker6 = []string{
	`["3♦", "A♠", "K♠", "J♣", "7♥", "5♦"]`, // high card
	`["3♦", "A♠", "A♣", "J♣", "7♥", "5♦"]`, // pair
	`["3♦", "A♠", "A♣", "J♣", "J♦", "5♦"]`, // two pair
	`["3♦", "A♠", "A♣", "A♦", "J♦", "5♦"]`, // three of a kind
	`["3♦", "A♠", "K♠", "Q♦", "J♥", "T♦"]`, // straight
	`["3♦", "T♠", "7♠", "4♠", "3♠", "2♠"]`, // flush
	`["3♦", "4♠", "4♣", "4♦", "2♠", "2♥"]`, // full house
	`["3♦", "A♠", "A♣", "A♦", "A♥", "5♥"]`, // four of a kind
	`["3♦", "A♠", "K♠", "Q♠", "J♠", "T♠"]`, // straight flush
}

var dataJoker7 = []string{
	`["2♦", "3♦", "A♠", "K♠", "J♣", "7♥", "5♦"]`, // high card
	`["2♦", "3♦", "A♠", "A♣", "J♣", "7♥", "5♦"]`, // pair
	`["2♦", "3♦", "A♠", "A♣", "J♣", "J♦", "5♦"]`, // two pair
	`["2♣", "3♦", "A♠", "A♣", "A♦", "J♦", "5♦"]`, // three of a kind
	`["2♦", "3♦", "A♠", "K♠", "Q♦", "J♥", "T♦"]`, // straight
	`["2♦", "3♦", "T♠", "7♠", "4♠", "3♠", "2♠"]`, // flush
	`["2♦", "3♦", "4♠", "4♣", "4♦", "2♠", "2♥"]`, // full house
	`["2♦", "3♦", "A♠", "A♣", "A♦", "A♥", "5♥"]`, // four of a kind
	`["2♦", "3♦", "A♠", "K♠", "Q♠", "J♠", "T♠"]`, // straight flush
}

var dataRiverboat5 = [][][]byte{
	{[]byte("As"), []byte("Ks"), []byte("Jc"), []byte("7h"), []byte("5d")}, // high card
	{[]byte("As"), []byte("Ac"), []byte("Jc"), []byte("7h"), []byte("5d")}, // pair
	{[]byte("As"), []byte("Ac"), []byte("Jc"), []byte("Jd"), []byte("5d")}, // two pair
	{[]byte("As"), []byte("Ac"), []byte("Ad"), []byte("Jd"), []byte("5d")}, // three of a kind
	{[]byte("As"), []byte("Ks"), []byte("Qd"), []byte("Jh"), []byte("Td")}, // straight
	{[]byte("Ts"), []byte("7s"), []byte("4s"), []byte("3s"), []byte("2s")}, // flush
	{[]byte("4s"), []byte("4c"), []byte("4d"), []byte("2s"), []byte("2h")}, // full house
	{[]byte("As"), []byte("Ac"), []byte("Ad"), []byte("Ah"), []byte("5h")}, // four of a kind
	{[]byte("As"), []byte("Ks"), []byte("Qs"), []byte("Js"), []byte("Ts")}, // straight flush
}

var dataRiverboat6 = [][][]byte{
	{[]byte("3d"), []byte("As"), []byte("Ks"), []byte("Jc"), []byte("7h"), []byte("5d")}, // high card
	{[]byte("3d"), []byte("As"), []byte("Ac"), []byte("Jc"), []byte("7h"), []byte("5d")}, // pair
	{[]byte("3d"), []byte("As"), []byte("Ac"), []byte("Jc"), []byte("Jd"), []byte("5d")}, // two pair
	{[]byte("3d"), []byte("As"), []byte("Ac"), []byte("Ad"), []byte("Jd"), []byte("5d")}, // three of a kind
	{[]byte("3d"), []byte("As"), []byte("Ks"), []byte("Qd"), []byte("Jh"), []byte("Td")}, // straight
	{[]byte("3d"), []byte("Ts"), []byte("7s"), []byte("4s"), []byte("3s"), []byte("2s")}, // flush
	{[]byte("3d"), []byte("4s"), []byte("4c"), []byte("4d"), []byte("2s"), []byte("2h")}, // full house
	{[]byte("3d"), []byte("As"), []byte("Ac"), []byte("Ad"), []byte("Ah"), []byte("5h")}, // four of a kind
	{[]byte("3d"), []byte("As"), []byte("Ks"), []byte("Qs"), []byte("Js"), []byte("Ts")}, // straight flush
}

var dataRiverboat7 = [][][]byte{
	{[]byte("2d"), []byte("3d"), []byte("As"), []byte("Ks"), []byte("Jc"), []byte("7h"), []byte("5d")}, // high card
	{[]byte("2d"), []byte("3d"), []byte("As"), []byte("Ac"), []byte("Jc"), []byte("7h"), []byte("5d")}, // pair
	{[]byte("2d"), []byte("3d"), []byte("As"), []byte("Ac"), []byte("Jc"), []byte("Jd"), []byte("5d")}, // two pair
	{[]byte("2c"), []byte("3d"), []byte("As"), []byte("Ac"), []byte("Ad"), []byte("Jd"), []byte("5d")}, // three of a kind
	{[]byte("2d"), []byte("3d"), []byte("As"), []byte("Ks"), []byte("Qd"), []byte("Jh"), []byte("Td")}, // straight
	{[]byte("2d"), []byte("3d"), []byte("Ts"), []byte("7s"), []byte("4s"), []byte("3s"), []byte("2s")}, // flush
	{[]byte("2d"), []byte("3d"), []byte("4s"), []byte("4c"), []byte("4d"), []byte("2s"), []byte("2h")}, // full house
	{[]byte("2d"), []byte("3d"), []byte("As"), []byte("Ac"), []byte("Ad"), []byte("Ah"), []byte("5h")}, // four of a kind
	{[]byte("2d"), []byte("3d"), []byte("As"), []byte("Ks"), []byte("Qs"), []byte("Js"), []byte("Ts")}, // straight flush
}

func BenchmarkFiveJoker(b *testing.B) {
	var cardsJoker5 [][]hand.Card
	for score := range dataJoker5 {
		var cards []hand.Card
		json.Unmarshal([]byte(dataJoker5[score]), &cards)
		cardsJoker5 = append(cardsJoker5, cards)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, cards := range cardsJoker5 {
			hand.New(cards)
		}
	}
}

func BenchmarkFivePoker(b *testing.B) {
	var cardsPoker5 [][]poker.Card
	for score := range dataPoker5 {
		var cards []poker.Card
		json.Unmarshal([]byte(dataPoker5[score]), &cards)
		cardsPoker5 = append(cardsPoker5, cards)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, cards := range cardsPoker5 {
			poker.Evaluate(cards)
		}
	}
}

func BenchmarkFiveRiverboat(b *testing.B) {
	var cardsRiverboat5 [][]Card
	for _, s := range dataRiverboat5 {
		var cards []Card
		for _, ss := range s {
			c := ParseCardStr(ss)
			cards = append(cards, c)
		}
		cardsRiverboat5 = append(cardsRiverboat5, cards)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, cards := range cardsRiverboat5 {
			HandValue(cards[0], cards[1], cards[2], cards[3], cards[4])
		}
	}
}

func BenchmarkSixJoker(b *testing.B) {
	var cardsJoker6 [][]hand.Card
	for score := range dataJoker6 {
		var cards []hand.Card
		json.Unmarshal([]byte(dataJoker6[score]), &cards)
		cardsJoker6 = append(cardsJoker6, cards)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, cards := range cardsJoker6 {
			hand.New(cards)
		}
	}
}

func BenchmarkSixPoker(b *testing.B) {
	var cardsPoker6 [][]poker.Card
	for score := range dataPoker6 {
		var cards []poker.Card
		json.Unmarshal([]byte(dataPoker6[score]), &cards)
		cardsPoker6 = append(cardsPoker6, cards)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, cards := range cardsPoker6 {
			poker.Evaluate(cards)
		}
	}
}

func BenchmarkSixRiverboat(b *testing.B) {
	var cardsRiverboat6 [][]Card
	for _, s := range dataRiverboat6 {
		var cards []Card
		for _, ss := range s {
			c := ParseCardStr(ss)
			cards = append(cards, c)
		}
		cardsRiverboat6 = append(cardsRiverboat6, cards)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, cards := range cardsRiverboat6 {
			BestFiveOfSix(cards[0], cards[1], cards[2], cards[3], cards[4], cards[5])
		}
	}
}

func BenchmarkSevenJoker(b *testing.B) {
	var cardsJoker7 [][]hand.Card
	for score := range dataJoker7 {
		var cards []hand.Card
		json.Unmarshal([]byte(dataJoker7[score]), &cards)
		cardsJoker7 = append(cardsJoker7, cards)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, cards := range cardsJoker7 {
			hand.New(cards)
		}
	}
}

func BenchmarkSevenPoker(b *testing.B) {
	var cardsPoker7 [][]poker.Card
	for score := range dataPoker7 {
		var cards []poker.Card
		json.Unmarshal([]byte(dataPoker7[score]), &cards)
		cardsPoker7 = append(cardsPoker7, cards)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, cards := range cardsPoker7 {
			poker.Evaluate(cards)
		}
	}
}

func BenchmarkSevenRiverboat(b *testing.B) {
	var cardsRiverboat7 [][]Card
	for _, s := range dataRiverboat7 {
		var cards []Card
		for _, ss := range s {
			c := ParseCardStr(ss)
			cards = append(cards, c)
		}
		cardsRiverboat7 = append(cardsRiverboat7, cards)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, cards := range cardsRiverboat7 {
			BestFiveOfSeven(cards[0], cards[1], cards[2], cards[3], cards[4], cards[5], cards[6])
		}
	}
}
