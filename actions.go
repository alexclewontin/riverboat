package riverboat

//The generic type of all actions, made to better allow external agents to interact with the game
type Action func(g *Game, pn uint, data uint) error

func Bet(g *Game, pn uint, data uint) error {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	if !g.getBetting() {
		return ErrMidRoundAction
	}

	if g.actionNum != pn {
		return ErrWrongPlayer
	}

	p, err := g.getPlayer(pn)
	if err != nil {
		return err
	}

	//rename this for readability
	betVal := data

	var minBet uint = g.toCall()
	var maxBet uint = g.getLimit()

	var isLegal bool

	//TODO: I don't love this if-else if chain, but I was originally using
	// a lambda with multiple returns as a control flow structure (which
	// really just avoided using the elses?), which definitely
	//hurts readability. Refactor to better express
	if !g.canOpen(pn) {
		//Won't hit now, reserved for future implementations
		isLegal = false
	} else if betVal >= maxBet {
		//You can always go all-in
		isLegal = true
	} else if betVal < (minBet - p.Bet) {
		//Not calling the minimum needed
		isLegal = false
	} else if betVal == (minBet - p.Bet) {
		//Calling exactly
		isLegal = true
	} else if betVal < (minBet + g.minRaise - p.Bet) {
		// More than calling, but less than minimum raise
		isLegal = false
	} else {
		// More than calling, and at least the minimum raise
		isLegal = true
		g.minRaise = betVal + p.Bet - minBet
	}

	if !isLegal {
		//I could just return this in every spot, but i suspect the structure of what is legal
		//will change as more betting schemes are introduced, so seems more extensible to keep it here
		return ErrBadBetAmt
	}

	return g.updateRoundInfo()
}

// BuyIn buys more chips for the player. Here, data is the amount to buy in for
func BuyIn(g *Game, pn uint, data uint) error {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	p, err := g.getPlayer(pn)
	if err != nil {
		return err
	}

	//Can't buy in while playing
	if p.In {
		return ErrMidRoundAction
	}

	//Can't buy more than the maximum buy, if it's configured
	if g.config.MaxBuy != 0 && p.Stack+data > g.config.MaxBuy {
		return ErrBuyTooBig
	}

	//Otherwise, add it to the stack
	p.Stack = p.Stack + data

	//And add it to your total
	p.TotalBuyIn = p.TotalBuyIn + data

	return nil
}

func Deal(g *Game, pn uint, data uint) error {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	if pn != g.dealerNum {
		return ErrWrongPlayer
	}

	stage, betting := g.getStageAndBetting()

	if betting {
		return ErrMidRoundAction
	}

	g.initStage()

	switch stage {
	case PreDeal:

		// Zero all the community cards from last round
		for i := range g.communityCards {
			g.communityCards[i] = 0
		}

		for !g.players[g.dealerNum].Ready {
			g.dealerNum = g.dealerNum + 1
		}

		err := g.updateBlindNums()
		if err != nil {
			return err
		}

		g.actionNum = g.utgNum

		for i := 0; i < 3; i++ {
			g.deck.Shuffle()
		}

		for i, p := range g.players {
			if p.Ready {
				g.players[i].Cards[0] = g.deck.Pop()
				g.players[i].Cards[1] = g.deck.Pop()
				g.players[i].In = true
			}
		}

		g.players[g.sbNum].putInChips(g.config.SmallBlind)
		g.players[g.bbNum].putInChips(g.config.BigBlind)

	case PreFlop:
		g.communityCards[0] = g.deck.Pop()
		g.communityCards[1] = g.deck.Pop()
		g.communityCards[2] = g.deck.Pop()

	case Flop:
		g.communityCards[3] = g.deck.Pop()

	case Turn:
		g.communityCards[4] = g.deck.Pop()

	case River:
		return ErrMidRoundAction
	default:
		return ErrInternalBadGameStage
	}

	g.setStageAndBetting(stage+1, true)

	return nil
}

func Fold(g *Game, pn uint, data uint) error {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	p, err := g.getPlayer(pn)
	if err != nil {
		return err
	}

	if g.actionNum != pn {
		return ErrWrongPlayer
	}

	p.In = false

	return g.updateRoundInfo()

}

func ToggleReady(g *Game, pn uint) error {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	p, err := g.getPlayer(pn)
	if err != nil {
		return err
	}

	if p.In {
		return ErrMidRoundAction
	}

	if p.Ready {
		p.Ready = false
		p.Cards[0] = 0
		p.Cards[1] = 0
	} else {
		if p.Stack == 0 {
			return ErrNoMoney
		}
		p.Ready = true
	}

	if pn == g.dealerNum {
		for !(g.players[g.dealerNum].Ready) {
			g.dealerNum = g.dealerNum + 1
		}
	}

	g.updateBlindNums()

	return nil
}
