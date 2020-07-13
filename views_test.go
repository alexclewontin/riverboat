package riverboat

import (
	"reflect"
	"sync"
	"testing"
	. "github.com/alexclewontin/riverboat/eval"
)

func TestGame_GenerateOmniView_Driver(t * testing.T) {

	g := &Game{}
	t.Run("Test Functionality", func(t *testing.T) {g.GenerateOmniView()})
}

func TestGame_GenerateOmniView(t *testing.T) {
	type fields struct {
		mtx            *sync.Mutex
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
	}
	tests := []struct {
		name    string
		fields  fields
		want    *GameView
		wantErr bool
	}{
		{
			name: "Basic (empty structs and slices)",
			fields: fields{
				mtx:    &sync.Mutex{},
				DealerNum:     3,
				ActionNum:     4,
				UTGNum:         5,
				SBNum:          6,
				BBNum:          7,
				CommunityCards: []Card{
					8394515,
					16783383,
					33564957,
					67115551,
					134224677,
					1601,
				},
				Flags:         0xF0,
				Config:        GameConfig{
				},
				Players:        []player{},
				Deck:           DefaultDeck,
				Pots:           []Pot{},
				MinRaise:       25,
			},
			want: &GameView{
				DealerNum:     3,
				ActionNum:     4,
				UTGNum:         5,
				SBNum:          6,
				BBNum:          7,
				CommunityCards: []Card{
					8394515,
					16783383,
					33564957,
					67115551,
					134224677,
					1601,
				},
				Flags:         0xF0,
				Config:        GameConfig{

				},
				Players:        []player{},
				Deck:           DefaultDeck,
				Pots:           []Pot{},
				MinRaise:       25,
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				//ignore warning here (the only important functionality is that the mutex is ignored during the deep copy)
				mtx:            *tt.fields.mtx,
				dealerNum:      tt.fields.DealerNum,
				actionNum:      tt.fields.ActionNum,
				utgNum:         tt.fields.UTGNum,
				sbNum:          tt.fields.SBNum,
				bbNum:          tt.fields.BBNum,
				communityCards: tt.fields.CommunityCards,
				flags:          tt.fields.Flags,
				config:         tt.fields.Config,
				players:        tt.fields.Players,
				deck:           tt.fields.Deck,
				pots:           tt.fields.Pots,
				minRaise:       tt.fields.MinRaise,
			}
			got := g.GenerateOmniView()
			var err error = nil
			if (err != nil) != tt.wantErr {
				t.Errorf("Game.GenerateOmniView() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Game.GenerateOmniView() = %+v\nwant %+v", got, tt.want)
			}
		})
	}
}


func TestGame_GenerateOmniViewChangedVals(t *testing.T) {
	type fields struct {
		mtx            *sync.Mutex
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
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Game
		wantErr bool
	}{
		{
			name: "Copy",
			fields: fields{
				mtx:    &sync.Mutex{},
				DealerNum:     3,
				ActionNum:     4,
				UTGNum:         5,
				SBNum:          6,
				BBNum:          7,
				CommunityCards: []Card{
					8394515,
					16783383,
					33564957,
					67115551,
					134224677,
					1601,
				},
				Flags:         0xF0,
				Config:        GameConfig{
					BigBlind: 25,
					SmallBlind: 10,
				},
				Players:        []player{
					{
						Ready: true,
						In: false,
						TotalBuyIn: 100,
						Stack: 105,
						Bet: 10,
						TotalBet: 20,
						Cards: [2]Card{
							33564957,
							67115551,
						},
					},
				},
				Deck:           Deck{},
				Pots:           []Pot{},
				MinRaise:       25,
			},
			want: &Game{
				dealerNum:     3,
				actionNum:     4,
				utgNum:         5,
				sbNum:          6,
				bbNum:          7,
				communityCards: []Card{
					8394515,
					16783383,
					33564957,
					67115551,
					134224677,
					1601,
				},
				flags:         0xF0,
				config:        GameConfig{
					BigBlind: 25,
					SmallBlind: 10,
				},
				players:        []player{
					{
						Ready: true,
						In: false,
						TotalBuyIn: 100,
						Stack: 105,
						Bet: 10,
						TotalBet: 20,
						Cards: [2]Card{
							33564957,
							67115551,
						},
					},
				},
				deck:           Deck{},
				pots:           []Pot{},
				minRaise:       25,
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				//ignore warning here (the only important functionality is that the mutex is ignored during the deep copy)
				mtx:            *tt.fields.mtx,
				dealerNum:      tt.fields.DealerNum,
				actionNum:      tt.fields.ActionNum,
				utgNum:         tt.fields.UTGNum,
				sbNum:          tt.fields.SBNum,
				bbNum:          tt.fields.BBNum,
				communityCards: tt.fields.CommunityCards,
				flags:          tt.fields.Flags,
				config:         tt.fields.Config,
				players:        tt.fields.Players,
				deck:           tt.fields.Deck,
				pots:           tt.fields.Pots,
				minRaise:       tt.fields.MinRaise,
			}
			var err error = nil
			view := g.GenerateOmniView()

			view.DealerNum = view.DealerNum + 1
			view.CommunityCards[2] = 0
			view.Config.BigBlind = 50
			view.Players[0].Cards[0] = 0
			view.Players[0].Bet = 20


			if (err != nil) != tt.wantErr {
				t.Errorf("Game.GenerateOmniView() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(g, tt.want) {
				t.Errorf("Game.GenerateOmniView() = %+v, want %+v", g, tt.want)
			}
		})
	}
}
