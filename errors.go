package riverboat

import (
	"errors"
)

var ErrMidRoundAction = errors.New("This action cannot be performed at this time")
var ErrBuyTooBig = errors.New("This would exceed the maximum configured purchased stack size.")
var ErrBadPlayerNum = errors.New("This player number does not exist.")
var ErrNoMoney = errors.New("Not enough money for this.")
var ErrWrongPlayer = errors.New("You are not the appropriate player to perform this action.")
var ErrNotEnoughPlayers = errors.New("Need more players to start the round.")
var ErrBadBetAmt = errors.New("This is not a legal bet amount.")

var ErrInternalBadGameStage = errors.New("Internal Error: Bad game stage")
