package riverboat

import (
	. "github.com/alexclewontin/riverboat/eval"
)

type GameView struct {
	DealerNum      uint
	ActionNum      uint
	UTGNum         uint
	SBNum          uint
	BBNum          uint
	CommunityCards []Card
	Flags          GameFlags
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
	// Something like reflection, but evaulated at compile-time would be ideal
	// Probably using go generate.

	//WARNING: This needs to be the deepest of deep copies. If adding a field, 
	//make sure that it is. An example: copying a slice of structs, where the struct
	//has a field that is a slice: this doesn't work by default. Write a helper function.
	view := &GameView{
		DealerNum: g.dealerNum,
		ActionNum: g.actionNum,
		UTGNum: g.utgNum,
		SBNum: g.sbNum,
		BBNum: g.bbNum,
		CommunityCards: append([]Card{}, g.communityCards...),
		Flags: g.flags,

		//TODO: is this necessary?
		Config: GameConfig{
			MaxBuy: g.config.MaxBuy,
			BigBlind: g.config.BigBlind,
			SmallBlind: g.config.SmallBlind,
		},
		Players: append([]player{}, g.players...),
		Deck: append([]Card{}, g.deck...),
		//TODO: I can already tell that this is broken
		Pots: append([]Pot{}, g.pots...),
		MinRaise: g.minRaise,
		ReadyCount: g.readyCount(),
	}

	return view
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
	g.flags = gv.Flags

	//TODO: is this necessary?
	g.config = GameConfig{
		MaxBuy: gv.Config.MaxBuy,
		BigBlind: gv.Config.BigBlind,
		SmallBlind: gv.Config.SmallBlind,
	}

	g.players = append([]player{}, gv.Players...)
	g.deck = append([]Card{}, gv.Deck...)

	//TODO: I can already tell that this is broken
	g.pots = append([]Pot{}, g.pots...)
	g.minRaise = gv.MinRaise




	// ActionNum: g.actionNum,
	// UTGNum: g.utgNum,
	// SBNum: g.sbNum,
	// BBNum: g.bbNum,
	// CommunityCards: append([]Card{}, g.communityCards...),
	// Flags: g.flags,
	// Config: GameConfig{
	// 	MaxBuy: g.config.MaxBuy,
	// 	BigBlind: g.config.BigBlind,
	// 	SmallBlind: g.config.SmallBlind,
	// },
	// Players: append([]player{}, g.players...),
	// Deck: append([]Card{}, g.deck...),
	// Pots: append([]Pot{}, g.pots...),
	// MinRaise: g.minRaise,
	// ReadyCount: g.readyCount(),

}
// GenerateView is primarily for creating a view that can be serialized for delivery to a specific player
func (g *Game) GenerateView(pn uint) *GameView {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	gv := g.copyToView()
	gv.Deck = nil

	// D. R. Y.!
	hideCards := func(pn2 uint) { gv.Players[pn2].Cards = [2]Card{ 0, 0 } }
	showCards := func(pn2 uint) { gv.Players[pn2].Cards = [2]Card{g.players[pn2].Cards[0], g.players[pn2].Cards[1]} }

	allInCount := 0
	inCount := 0

	for i, p := range g.players {
		if uint(i) != pn {
			hideCards(uint(i))
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

	if g.getStage() == Over {
		for _, pot := range g.pots {
			//TODO: technically, this should start with the called player and go around.
			// For the moment, we're just gonna show the winners
			for i := range pot.WinningPlayerNums {
				showCards(uint(i))
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



