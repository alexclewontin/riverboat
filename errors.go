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
