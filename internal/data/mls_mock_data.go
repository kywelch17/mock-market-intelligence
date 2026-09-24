package data

import (
	"github.com/kywelch17/mock-market-intelligence/internal/models"
	"math/rand"
)

var mlsEvents = []models.Event{
	{
		ID:          "mls-fc-cincinnati-cf-montreal",
		Type:        "Game",
		League:      models.LeagueMLS,
		Description: "FC Cincinnati @ CF Montréal",
		Markets: []models.Market{
			{
				ID:          "mls-cincinnati-regular-time-win",
				EventID:     "mls-fc-cincinnati-cf-montreal",
				Description: "Cincinnati Regular Time Win",
				Status:      models.StatusOpen,
				Type:        models.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
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
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "kenji-mboma-dem",
					Name:     "Kenji Mboma Dem",
					Number:   "99",  // Assigning a placeholder number
					Position: "FW",  // Assuming forward since it's shots on target
					Team:     "CIN", // FC Cincinnati abbreviation
				},
				Outcomes: []models.Outcome{
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
