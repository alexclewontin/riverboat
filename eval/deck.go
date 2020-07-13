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
	"bytes"
	"encoding/binary"
	"io"
	"math/rand"
	"time"
	//"log"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	for _, s := range suits {
		for j := 0; j < 13; j++ {
			DefaultDeck.Push(Card((1 << (16 + j)) | s | (int32(j) << 8) | primeRanks[j]))
		}
	}
}

type Deck []Card

var DefaultDeck Deck

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

func (d *Deck) Unmarshal(v interface{}) error {
	data, ok := v.([]byte)
	if !ok {
		*d = nil
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

func (d Deck) MarshalDB() (interface{}, error) { return d.Marshal() }
func (d Deck) UnmarshalDB(v interface{}) error { return d.Unmarshal(v) }

// Remove and return top card. Return 0 if stack is empty.
func (d *Deck) Pop() Card {
	if len(*d) == 0 {
		return 0
	} else {
		ndx := len(*d) - 1
		element := (*d)[ndx]
		*d = (*d)[:ndx]
		return element
	}
}

func (d *Deck) IsEmpty() bool { return (len(*d) == 0) }

func (d *Deck) Push(card Card) { *d = append(*d, card) }

func (d *Deck) Shuffle() {
	*d = append([]Card{}, DefaultDeck...)
	rand.Shuffle(len(*d), func(i, j int) { (*d)[i], (*d)[j] = (*d)[j], (*d)[i] })
}
