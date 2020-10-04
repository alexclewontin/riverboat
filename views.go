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

package riverboat

import (
	. "github.com/alexclewontin/riverboat/eval"
)

// GameView is the type that represents a snapshot of a Game's state.
type GameView struct {
	DealerNum      uint
	ActionNum      uint
	UTGNum         uint
	SBNum          uint
	BBNum          uint
	CommunityCards []Card
	Stage          GameStage
	Betting        bool
	Config         GameConfig
	Players        []player
	Deck           Deck
	Pots           []Pot
	MinRaise       uint
	ReadyCount     uint
}

func (g *Game) copyToView() *GameView {
	//TODO: Is there some way to do this programatically? I considered using
	// reflection, but since that happens at runtime it is less performant.
	// Something like reflection, but evaluated at compile-time would be ideal
	// Probably using go generate.

	//WARNING: This needs to be the deepest of deep copies. If adding a field,
	//make sure that it is. An example: copying a slice of structs, where the struct
	//has a field that is a slice: this doesn't work by default. Write a helper function.
	view := &GameView{
		DealerNum:      g.dealerNum,
		ActionNum:      g.actionNum,
		UTGNum:         g.utgNum,
		SBNum:          g.sbNum,
		BBNum:          g.bbNum,
		CommunityCards: append([]Card{}, g.communityCards...),
		Stage:          g.getStage(),
		Betting:        g.getBetting(),
		Config:         g.config,
		Players:        append([]player{}, g.players...),
		Deck:           append([]Card{}, g.deck...),
		Pots:           copyPots(g.pots),
		MinRaise:       g.minRaise,
		ReadyCount:     g.readyCount(),
	}

	return view
}

func copyPots(src []Pot) []Pot {
	ret := make([]Pot, len(src))
	for i := range src {
		ret[i].Amt = src[i].Amt
		ret[i].TopShare = src[i].TopShare
		ret[i].WinningScore = src[i].WinningScore
		ret[i].EligiblePlayerNums = append([]uint{}, src[i].EligiblePlayerNums...)
		ret[i].WinningPlayerNums = append([]uint{}, src[i].WinningPlayerNums...)
		ret[i].WinningHand = append([]Card{}, src[i].WinningHand...)
	}

	return ret
}

// FillFromView is primarily for loading a stored view from a persistence layer
func (g *Game) FillFromView(gv *GameView) {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	g.dealerNum = gv.DealerNum
	g.actionNum = gv.ActionNum
	g.utgNum = gv.UTGNum
	g.bbNum = gv.BBNum
	g.sbNum = gv.SBNum
	g.communityCards = append([]Card{}, gv.CommunityCards...)
	g.setStageAndBetting(gv.Stage, gv.Betting)
	g.config = gv.Config
	g.players = append([]player{}, gv.Players...)
	g.deck = append([]Card{}, gv.Deck...)
	g.pots = copyPots(gv.Pots)
	g.minRaise = gv.MinRaise
}

// GeneratePlayerView is primarily for creating a view that can be serialized for delivery to a specific player
// The generated view holds only the information that the player denoted by pn is entitled to see at the moment it is generated.
func (g *Game) GeneratePlayerView(pn uint) *GameView {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	gv := g.copyToView()
	gv.Deck = nil

	// D. R. Y.!
	hideCards := func(pn2 uint) { gv.Players[pn2].Cards = [2]Card{0, 0} }
	showCards := func(pn2 uint) { gv.Players[pn2].Cards = [2]Card{g.players[pn2].Cards[0], g.players[pn2].Cards[1]} }

	allInCount := 0
	inCount := 0

	for i, p := range g.players {
		if uint(i) != pn {
			hideCards(uint(i))
		} else {
			if !g.players[pn].In {
				hideCards(uint(i))
			}
		}

		if p.allIn() {
			allInCount++
		}

		if p.In {
			inCount++
		}

	}

	// If in a heads-up situation
	if allInCount == inCount {
		for i, p := range g.players {
			if p.In {
				showCards(uint(i))
			}
		}
	}

	if g.getStage() == PreDeal {

		showCards(g.calledNum)
		_, scoreToBeat := BestFiveOfSeven(
			g.players[g.calledNum].Cards[0],
			g.players[g.calledNum].Cards[1],
			g.communityCards[0],
			g.communityCards[1],
			g.communityCards[2],
			g.communityCards[3],
			g.communityCards[4],
		)

		for i := range g.players {
			pni := (g.calledNum + uint(i)) % uint(len(g.players))
			_, iScore := BestFiveOfSeven(
				g.players[pni].Cards[0],
				g.players[pni].Cards[1],
				g.communityCards[0],
				g.communityCards[1],
				g.communityCards[2],
				g.communityCards[3],
				g.communityCards[4],
			)

			if (iScore < scoreToBeat) && g.players[pni].In {
				showCards(pni)
				scoreToBeat = iScore
			}
		}

		for _, pot := range g.pots {

			for _, j := range pot.WinningPlayerNums {
				showCards(j)
			}
		}
	}

	return gv
}

// GenerateOmniView is primarily for creating a view that can be serialized for delivery to a persistance layer, like a db or in-memory store
// Nothing is censored, not even the contents of the deck
func (g *Game) GenerateOmniView() *GameView {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	return g.copyToView()

}
