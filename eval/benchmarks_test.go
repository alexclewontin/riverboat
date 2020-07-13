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
	"encoding/json"
	"testing"

	"github.com/chehsunliu/poker"
	"github.com/notnil/joker/hand"
)

var dataPoker1 = map[int32]string{
	6252: `["As", "Ks", "Jc", "7h", "5d"]`, // high card
	3448: `["As", "Ac", "Jc", "7h", "5d"]`, // pair
	2497: `["As", "Ac", "Jc", "Jd", "5d"]`, // two pair
	1636: `["As", "Ac", "Ad", "Jd", "5d"]`, // three of a kind
	1600: `["As", "Ks", "Qd", "Jh", "Td"]`, // straight
	1542: `["Ts", "7s", "4s", "3s", "2s"]`, // flush
	298:  `["4s", "4c", "4d", "2s", "2h"]`, // full house
	19:   `["As", "Ac", "Ad", "Ah", "5h"]`, // four of a kind
	1:    `["As", "Ks", "Qs", "Js", "Ts"]`, // straight flush
}

var dataPoker2 = map[int32]string{
	6252: `["3d", "As", "Ks", "Jc", "7h", "5d"]`, // high card
	3448: `["3d", "As", "Ac", "Jc", "7h", "5d"]`, // pair
	2497: `["3d", "As", "Ac", "Jc", "Jd", "5d"]`, // two pair
	1636: `["3d", "As", "Ac", "Ad", "Jd", "5d"]`, // three of a kind
	1600: `["3d", "As", "Ks", "Qd", "Jh", "Td"]`, // straight
	1542: `["3d", "Ts", "7s", "4s", "3s", "2s"]`, // flush
	298:  `["3d", "4s", "4c", "4d", "2s", "2h"]`, // full house
	19:   `["3d", "As", "Ac", "Ad", "Ah", "5h"]`, // four of a kind
	1:    `["3d", "As", "Ks", "Qs", "Js", "Ts"]`, // straight flush
}

var dataPoker3 = map[int32]string{
	6252: `["2d", "3d", "As", "Ks", "Jc", "7h", "5d"]`, // high card
	3448: `["2d", "3d", "As", "Ac", "Jc", "7h", "5d"]`, // pair
	2497: `["2d", "3d", "As", "Ac", "Jc", "Jd", "5d"]`, // two pair
	1636: `["2c", "3d", "As", "Ac", "Ad", "Jd", "5d"]`, // three of a kind
	1600: `["2d", "3d", "As", "Ks", "Qd", "Jh", "Td"]`, // straight
	1542: `["2d", "3d", "Ts", "7s", "4s", "3s", "2s"]`, // flush
	298:  `["2d", "3d", "4s", "4c", "4d", "2s", "2h"]`, // full house
	19:   `["2d", "3d", "As", "Ac", "Ad", "Ah", "5h"]`, // four of a kind
	1:    `["2d", "3d", "As", "Ks", "Qs", "Js", "Ts"]`, // straight flush
}

var dataJoker1 = []string{
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

var dataJoker2 = []string{
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

var dataJoker3 = []string{
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

var dataRiverboat1 = [][][]byte{
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

var dataRiverboat2 = [][][]byte{
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

var dataRiverboat3 = [][][]byte{
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

func BenchmarkFivePoker(b *testing.B) {
	var cardsPoker1 [][]poker.Card
	for score := range dataPoker1 {
		var cards []poker.Card

		json.Unmarshal([]byte(dataPoker1[score]), &cards)
		cardsPoker1 = append(cardsPoker1, cards)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, cards := range cardsPoker1 {
			poker.Evaluate(cards)
		}
	}
}

func BenchmarkFiveJoker(b *testing.B) {

	var cardsJoker1 [][]hand.Card
	for score := range dataJoker1 {
		var cards []hand.Card

		json.Unmarshal([]byte(dataJoker1[score]), &cards)
		cardsJoker1 = append(cardsJoker1, cards)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, cards := range cardsJoker1 {
			hand.New(cards)
		}
	}
}

func BenchmarkFiveRiverboat(b *testing.B) {
	var cardsRiverboat1 [][]Card

	for _, s := range dataRiverboat1 {
		var cards []Card

		for _, ss := range s {
			c := ParseCardStr(ss)
			cards = append(cards, c)
		}

		cardsRiverboat1 = append(cardsRiverboat1, cards)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, cards := range cardsRiverboat1 {
			HandValue(cards[0], cards[1], cards[2], cards[3], cards[4])
		}
	}
}

// func BenchmarkSixPoker(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		for _, cards := range cardsPoker2 {
// 			poker.Evaluate(cards)
// 		}
// 	}
// }

// func BenchmarkSixJoker(b *testing.B) {

// 	for i := 0; i < b.N; i++ {
// 		for _, cards := range cardsJoker2 {
// 			hand.New(cards)
// 		}
// 	}
// }

// func BenchmarkSixRiverboat(b *testing.B) {

// 	for i := 0; i < b.N; i++ {
// 		for _, cards := range cardsRiverboat3 {
// 			BestFiveOfSix(cards[0], cards[1], cards[2], cards[3], cards[4], cards[5])
// 		}
// 	}
// }

// func BenchmarkSevenPoker(b *testing.B) {

// 	for i := 0; i < b.N; i++ {
// 		for _, cards := range cardsPoker3 {
// 			poker.Evaluate(cards)
// 		}
// 	}
// }

// func BenchmarkSevenJoker(b *testing.B) {

// 	for i := 0; i < b.N; i++ {
// 		for _, cards := range cardsJoker3 {
// 			hand.New(cards)
// 		}
// 	}
// }

// func BenchmarkSevenRiverboat(b *testing.B) {

// 	for i := 0; i < b.N; i++ {
// 		for _, cards := range cardsRiverboat3 {
// 			BestFiveOfSeven(cards[0], cards[1], cards[2], cards[3], cards[4], cards[5], cards[6])
// 		}
// 	}
// }
