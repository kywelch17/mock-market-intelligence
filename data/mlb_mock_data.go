package data

import (
	"github.com/kywelch17/mock-market-intelligence/internal/models"
	"math/rand"
)

var MLBEvents = []models.Event{
	{
		ID:          "mlb-yankees-redsox",
		Type:        "Game",
		League:      models.LeagueMLB,
		Description: "New York Yankees @ Boston Red Sox",
		Markets: []models.Market{
			{
				ID:          "mlb-yankees-runline",
				EventID:     "mlb-yankees-redsox",
				Description: "New York Yankees Run Line -1.5",
				Status:      models.StatusOpen,
				Type:        models.TypeSpread,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "mlb-yankees-minus-1-5",
						Description: "Yankees -1.5",
						Index:       0,
					},
					{
						ID:          "mlb-redsox-plus-1-5",
						Description: "Red Sox +1.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "mlb-total-runs",
				EventID:     "mlb-yankees-redsox",
				Description: "Total Runs 8.5",
				Status:      models.StatusOpen,
				Type:        models.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "mlb-over-8-5-runs",
						Description: "Over 8.5",
						Index:       0,
					},
					{
						ID:          "mlb-under-8-5-runs",
						Description: "Under 8.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "mlb-aaron-judge-homeruns",
				EventID:     "mlb-yankees-redsox",
				Description: "Aaron Judge 0.5 Home Runs",
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "aaron-judge",
					Name:     "Aaron Judge",
					Number:   "99",
					Position: "RF",
					Team:     "NYY",
				},
				Outcomes: []models.Outcome{
					{
						ID:          "mlb-aaron-judge-over-0-5-hr",
						Description: "Over 0.5",
						Index:       0,
					},
					{
						ID:          "mlb-aaron-judge-under-0-5-hr",
						Description: "Under 0.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "mlb-dodgers-giants",
		Type:        "Game",
		League:      models.LeagueMLB,
		Description: "Los Angeles Dodgers @ San Francisco Giants",
		Markets: []models.Market{
			{
				ID:          "mlb-mookie-betts-hits",
				EventID:     "mlb-dodgers-giants",
				Description: "Mookie Betts 1.5 Hits",
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "mookie-betts",
					Name:     "Mookie Betts",
					Number:   "50",
					Position: "RF",
					Team:     "LAD",
				},
				Outcomes: []models.Outcome{
					{
						ID:          "mlb-mookie-betts-over-1-5-hits",
						Description: "Over 1.5",
						Index:       0,
					},
					{
						ID:          "mlb-mookie-betts-under-1-5-hits",
						Description: "Under 1.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "mlb-dodgers-pitcher-strikeouts",
				EventID:     "mlb-dodgers-giants",
				Description: "Shohei Ohtani 6.5 Strikeouts",
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "shohei-ohtani",
					Name:     "Shohei Ohtani",
					Number:   "17",
					Position: "SP",
					Team:     "LAD",
				},
				Outcomes: []models.Outcome{
					{
						ID:          "mlb-shohei-ohtani-over-6-5-k",
						Description: "Over 6.5",
						Index:       0,
					},
					{
						ID:          "mlb-shohei-ohtani-under-6-5-k",
						Description: "Under 6.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "mlb-pirates-rockies",
		Type:        "Game",
		League:      models.LeagueMLB,
		Description: "Pittsburgh Pirates @ Colorado Rockies",
		Markets: []models.Market{
			{
				ID:          "mlb-pirates-ml",
				EventID:     "mlb-pirates-rockies",
				Description: "Pittsburgh Pirates ML",
				Status:      models.StatusOpen,
				Type:        models.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "mlb-pirates-ml-yes",
						Description: "Yes",
						Index:       0,
					},
					{
						ID:          "mlb-pirates-ml-no",
						Description: "No",
						Index:       1,
					},
				},
			},
			{
				ID:          "mlb-rockies-over-under",
				EventID:     "mlb-pirates-rockies",
				Description: "Total Runs 10.5",
				Status:      models.StatusOpen,
				Type:        models.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "mlb-over-10-5-runs",
						Description: "Over 10.5",
						Index:       0,
					},
					{
						ID:          "mlb-under-10-5-runs",
						Description: "Under 10.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "mlb-bryan-reynolds-hits",
				EventID:     "mlb-pirates-rockies",
				Description: "Bryan Reynolds 1.5 Hits",
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "bryan-reynolds",
					Name:     "Bryan Reynolds",
					Number:   "22",
					Position: "CF",
					Team:     "PIT",
				},
				Outcomes: []models.Outcome{
					{
						ID:          "mlb-bryan-reynolds-over-1-5-hits",
						Description: "Over 1.5",
						Index:       0,
					},
					{
						ID:          "mlb-bryan-reynolds-under-1-5-hits",
						Description: "Under 1.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "mlb-yankees-futures",
		Type:        models.TypeFuture,
		League:      models.LeagueMLB,
		Description: "New York Yankees",
		Markets: []models.Market{
			{
				ID:          "mlb-yankees-win-world-series",
				EventID:     "mlb-yankees-futures",
				Description: "Win the World Series",
				Status:      models.StatusOpen,
				Type:        models.TypeFuture,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "mlb-yankees-win-world-series-yes",
						Description: "Yes",
						Index:       0,
					},
					{
						ID:          "mlb-yankees-win-world-series-no",
						Description: "No",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "mlb-dodgers-futures",
		Type:        models.TypeFuture,
		League:      models.LeagueMLB,
		Description: "Los Angeles Dodgers",
		Markets: []models.Market{
			{
				ID:          "mlb-dodgers-win-world-series",
				EventID:     "mlb-dodgers-futures",
				Description: "Win the World Series",
				Status:      models.StatusOpen,
				Type:        models.TypeFuture,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "mlb-dodgers-win-world-series-yes",
						Description: "Yes",
						Index:       0,
					},
					{
						ID:          "mlb-dodgers-win-world-series-no",
						Description: "No",
						Index:       1,
					},
				},
			},
		},
	},
}
