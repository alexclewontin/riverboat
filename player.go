package riverboat

import (
	. "github.com/alexclewontin/riverboat/eval"
)

type player struct {
	Ready      bool
	In         bool
	TotalBuyIn uint
	Stack      uint
	Bet        uint
	TotalBet   uint
	Cards      [2]Card
}

func (p *player) allIn() bool {
	return p.In && (p.Stack == 0)
}

func (p *player) initialize() {
	*p = player{}

	p.Ready = false
	p.In = false

}

//putInChips is simply a helper function that transfers the amounts between fields
func (p *player) putInChips(amt uint) {
	if p.Stack > amt {
		p.Bet += amt
		p.TotalBet += amt
		p.Stack -= amt
	} else {
		p.Bet += p.Stack
		p.TotalBet += p.Stack
		p.Stack = 0
	}
}

func (p *player) returnChips(amt uint) {
	if p.TotalBet > amt {
		p.TotalBet -= amt
		p.Stack -= amt
	} else {
		p.Stack += p.TotalBet
		p.TotalBet = 0
	}
}
