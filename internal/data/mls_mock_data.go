package data

import (
	"math/rand"

	"github.com/kywelch17/mock-market-intelligence/internal/models"
)

var mlsEvents = []market.Event{
	{
		ID:          "mls-fc-cincinnati-cf-montreal",
		Type:        "Game",
		League:      market.LeagueMLS,
		Description: "FC Cincinnati @ CF Montréal",
		Markets: []market.Market{
			{
				ID:          "mls-cincinnati-regular-time-win",
				EventID:     "mls-fc-cincinnati-cf-montreal",
				Description: "Cincinnati Regular Time Win",
				Status:      market.StatusOpen,
				Type:        market.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []market.Outcome{
					{
						ID:          "mls-cincinnati-win-yes",
						Description: "Yes",
						Index:       0,
					},
					{
						ID:          "mls-cincinnati-win-no",
						Description: "No",
						Index:       1,
					},
				},
			},
			{
				ID:          "mls-kenji-mboma-dem-shots-on-target",
				EventID:     "mls-fc-cincinnati-cf-montreal",
				Description: "Kenji Mboma Dem 1.5 Shots On Target",
				Status:      market.StatusOpen,
				Type:        market.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &market.Player{
					ID:       "kenji-mboma-dem",
					Name:     "Kenji Mboma Dem",
					Number:   "99",  // Assigning a placeholder number
					Position: "FW",  // Assuming forward since it's shots on target
					Team:     "CIN", // FC Cincinnati abbreviation
				},
				Outcomes: []market.Outcome{
					{
						ID:          "mls-kenji-mboma-dem-over-1-5-sots",
						Description: "Over 1.5",
						Index:       0,
					},
					{
						ID:          "mls-kenji-mboma-dem-under-1-5-sots",
						Description: "Under 1.5",
						Index:       1,
					},
				},
			},
		},
	},
}
