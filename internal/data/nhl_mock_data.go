package data

import (
	"math/rand"

	"github.com/kywelch17/mock-market-intelligence/internal/market"
)

var nhlEvents = []market.Event{
	{
		ID:          "nhl-leafs-bruins",
		Type:        "Game",
		League:      market.LeagueNHL,
		Description: "Toronto Maple Leafs @ Boston Bruins",
		Markets: []market.Market{
			{
				ID:          "nhl-leafs-moneyline",
				EventID:     "nhl-leafs-bruins",
				Description: "Toronto Maple Leafs ML",
				Status:      market.StatusOpen,
				Type:        market.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []market.Outcome{
					{
						ID:          "nhl-leafs-ml-yes",
						Description: "Yes",
						Index:       0,
					},
					{
						ID:          "nhl-leafs-ml-no",
						Description: "No",
						Index:       1,
					},
				},
			},
			{
				ID:          "nhl-bruins-puck-line",
				EventID:     "nhl-leafs-bruins",
				Description: "Boston Bruins Puck Line -1.5",
				Status:      market.StatusOpen,
				Type:        market.TypeSpread,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []market.Outcome{
					{
						ID:          "nhl-bruins-minus-1-5",
						Description: "Bruins -1.5",
						Index:       0,
					},
					{
						ID:          "nhl-leafs-plus-1-5",
						Description: "Leafs +1.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nhl-total-goals",
				EventID:     "nhl-leafs-bruins",
				Description: "Total Goals 6.5",
				Status:      market.StatusOpen,
				Type:        market.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []market.Outcome{
					{
						ID:          "nhl-over-6-5-goals",
						Description: "Over 6.5",
						Index:       0,
					},
					{
						ID:          "nhl-under-6-5-goals",
						Description: "Under 6.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nhl-auston-matthews-shots",
				EventID:     "nhl-leafs-bruins",
				Description: "Auston Matthews 3.5 Shots On Target",
				Status:      market.StatusOpen,
				Type:        market.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &market.Player{
					ID:       "auston-matthews",
					Name:     "Auston Matthews",
					Number:   "34",
					Position: "C",
					Team:     "TOR",
				},
				Outcomes: []market.Outcome{
					{
						ID:          "nhl-auston-matthews-over-3-5-sots",
						Description: "Over 3.5",
						Index:       0,
					},
					{
						ID:          "nhl-auston-matthews-under-3-5-sots",
						Description: "Under 3.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nhl-oilers-flames",
		Type:        "Game",
		League:      market.LeagueNHL,
		Description: "Edmonton Oilers @ Calgary Flames",
		Markets: []market.Market{
			{
				ID:          "nhl-oilers-moneyline",
				EventID:     "nhl-oilers-flames",
				Description: "Edmonton Oilers ML",
				Status:      market.StatusOpen,
				Type:        market.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []market.Outcome{
					{
						ID:          "nhl-oilers-ml-yes",
						Description: "Yes",
						Index:       0,
					},
					{
						ID:          "nhl-oilers-ml-no",
						Description: "No",
						Index:       1,
					},
				},
			},
			{
				ID:          "nhl-oilers-puck-line",
				EventID:     "nhl-oilers-flames",
				Description: "Edmonton Oilers Puck Line -1.5",
				Status:      market.StatusOpen,
				Type:        market.TypeSpread,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []market.Outcome{
					{
						ID:          "nhl-oilers-minus-1-5",
						Description: "Oilers -1.5",
						Index:       0,
					},
					{
						ID:          "nhl-flames-plus-1-5",
						Description: "Flames +1.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nhl-total-goals-alt",
				EventID:     "nhl-oilers-flames",
				Description: "Total Goals 5.5",
				Status:      market.StatusOpen,
				Type:        market.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []market.Outcome{
					{
						ID:          "nhl-over-5-5-goals",
						Description: "Over 5.5",
						Index:       0,
					},
					{
						ID:          "nhl-under-5-5-goals",
						Description: "Under 5.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nhl-mcdavid-assists",
				EventID:     "nhl-oilers-flames",
				Description: "Connor McDavid 1.5 Assists",
				Status:      market.StatusOpen,
				Type:        market.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &market.Player{
					ID:       "connor-mcdavid",
					Name:     "Connor McDavid",
					Number:   "97",
					Position: "C",
					Team:     "EDM",
				},
				Outcomes: []market.Outcome{
					{
						ID:          "nhl-mcdavid-over-1-5-assists",
						Description: "Over 1.5",
						Index:       0,
					},
					{
						ID:          "nhl-mcdavid-under-1-5-assists",
						Description: "Under 1.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nhl-flames-goalie-saves",
				EventID:     "nhl-oilers-flames",
				Description: "Jacob Markstrom 28.5 Saves",
				Status:      market.StatusOpen,
				Type:        market.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &market.Player{
					ID:       "jacob-markstrom",
					Name:     "Jacob Markstrom",
					Number:   "25",
					Position: "G",
					Team:     "CGY",
				},
				Outcomes: []market.Outcome{
					{
						ID:          "nhl-markstrom-over-28-5-saves",
						Description: "Over 28.5",
						Index:       0,
					},
					{
						ID:          "nhl-markstrom-under-28-5-saves",
						Description: "Under 28.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nhl-rangers-futures",
		Type:        market.TypeFuture,
		League:      market.LeagueNHL,
		Description: "New York Rangers",
		Markets: []market.Market{
			{
				ID:          "nhl-rangers-win-stanley-cup",
				EventID:     "nhl-rangers-futures",
				Description: "Win the Stanley Cup",
				Status:      market.StatusOpen,
				Type:        market.TypeFuture,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []market.Outcome{
					{
						ID:          "nhl-rangers-win-stanley-cup-yes",
						Description: "Yes",
						Index:       0,
					},
					{
						ID:          "nhl-rangers-win-stanley-cup-no",
						Description: "No",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nhl-avalanche-futures",
		Type:        market.TypeFuture,
		League:      market.LeagueNHL,
		Description: "Colorado Avalanche",
		Markets: []market.Market{
			{
				ID:          "nhl-avalanche-win-stanley-cup",
				EventID:     "nhl-avalanche-futures",
				Description: "Win the Stanley Cup",
				Status:      market.StatusOpen,
				Type:        market.TypeFuture,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []market.Outcome{
					{
						ID:          "nhl-avalanche-win-stanley-cup-yes",
						Description: "Yes",
						Index:       0,
					},
					{
						ID:          "nhl-avalanche-win-stanley-cup-no",
						Description: "No",
						Index:       1,
					},
				},
			},
		},
	},
}
