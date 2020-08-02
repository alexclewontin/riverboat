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
	"testing"
)

func TestIntegration_Scenarios(t *testing.T) {

	t.Run("Scenario 1", func(t *testing.T) {
		g := NewGame()

		pn_a := g.AddPlayer()
		g.AddPlayer()
		g.AddPlayer()

		err := Deal(g, pn_a, 0)

		if err != ErrIllegalAction {
			t.Error("Test failed - Deal must return ErrIllegalAction as 0 players are marked ready.")
		}
	})

	t.Run("Scenario 2", func(t *testing.T) {
		g := NewGame()

		pn_a := g.AddPlayer()
		g.AddPlayer()
		g.AddPlayer()

		err := ToggleReady(g, pn_a, 0)

		if err != ErrIllegalAction {
			t.Error("Test failed - ToggleReady must return ErrIllegalAction as player 0 has not bought in.")
		}
	})

	t.Run("Scenario 3", func(t *testing.T) {
		var err error
		g := NewGame()

		pn_a := g.AddPlayer()
		pn_b := g.AddPlayer()
		pn_c := g.AddPlayer()

		err = BuyIn(g, pn_a, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_b, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_c, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = ToggleReady(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}
	})

	t.Run("Scenario 4", func(t *testing.T) {
		var err error
		g := NewGame()

		pn_a := g.AddPlayer()
		pn_b := g.AddPlayer()
		pn_c := g.AddPlayer()

		err = BuyIn(g, pn_a, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_b, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_c, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = ToggleReady(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = Deal(g, pn_b, 0)

		if err != ErrIllegalAction {
			t.Errorf("Test failed - must return ErrIllegalAction as pn_b is not the dealer")
		}
	})

	t.Run("Scenario 5", func(t *testing.T) {
		var err error
		g := NewGame()

		pn_a := g.AddPlayer()
		pn_b := g.AddPlayer()
		pn_c := g.AddPlayer()

		err = BuyIn(g, pn_a, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_b, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_c, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = ToggleReady(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}

		err = Bet(g, pn_a, 25)

		if err != nil {

			t.Errorf("Test failed - error betting: %s", err)
		}

		if g.players[pn_a].Bet != 25 {
			t.Errorf("Betting mechanic not working.")
		}
	})

	t.Run("Scenario 6", func(t *testing.T) {
		var err error
		g := NewGame()

		pn_a := g.AddPlayer()
		pn_b := g.AddPlayer()
		pn_c := g.AddPlayer()

		err = BuyIn(g, pn_a, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_b, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_c, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = ToggleReady(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		// Preflop

		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}

		err = Bet(g, pn_a, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_b, 15)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		// Flop
		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}

		err = Bet(g, pn_b, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_c, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		// Turn
		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}

		err = Bet(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		//River

		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}

		err = Bet(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

	})

	t.Run("Scenario 7", func(t *testing.T) {
		var err error
		g := NewGame()

		pn_a := g.AddPlayer()
		pn_b := g.AddPlayer()
		pn_c := g.AddPlayer()

		err = BuyIn(g, pn_a, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_b, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = BuyIn(g, pn_c, 100)

		if err != nil {
			t.Errorf("Test failed - Error buying in: %s", err)
		}

		err = ToggleReady(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		err = ToggleReady(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - Error marking ready: %s", err)
		}

		// Preflop

		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}

		err = Bet(g, pn_a, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_b, 15)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		// Flop
		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}

		err = Bet(g, pn_b, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Fold(g, pn_c, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 25)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		// Turn
		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}

		err = Bet(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		//River

		err = Deal(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error dealing: %s", err)
		}

		err = Bet(g, pn_b, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

		err = Bet(g, pn_a, 0)

		if err != nil {
			t.Errorf("Test failed - error betting: %s", err)
		}

	})
}
