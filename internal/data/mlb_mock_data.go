package data

import (
	"math/rand"

	"github.com/kywelch17/mock-market-intelligence/internal/models"
)

var mlbEvents = []market.Event{
	{
		ID:          "mlb-yankees-redsox",
		Type:        "Game",
		League:      market.LeagueMLB,
		Description: "New York Yankees @ Boston Red Sox",
		Markets: []market.Market{
			{
				ID:          "mlb-yankees-runline",
				EventID:     "mlb-yankees-redsox",
				Description: "New York Yankees Run Line -1.5",
				Status:      market.StatusOpen,
				Type:        market.TypeSpread,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []market.Outcome{
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
				Status:      market.StatusOpen,
				Type:        market.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []market.Outcome{
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
				Status:      market.StatusOpen,
				Type:        market.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &market.Player{
					ID:       "aaron-judge",
					Name:     "Aaron Judge",
					Number:   "99",
					Position: "RF",
					Team:     "NYY",
				},
				Outcomes: []market.Outcome{
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
		League:      market.LeagueMLB,
		Description: "Los Angeles Dodgers @ San Francisco Giants",
		Markets: []market.Market{
			{
				ID:          "mlb-mookie-betts-hits",
				EventID:     "mlb-dodgers-giants",
				Description: "Mookie Betts 1.5 Hits",
				Status:      market.StatusOpen,
				Type:        market.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &market.Player{
					ID:       "mookie-betts",
					Name:     "Mookie Betts",
					Number:   "50",
					Position: "RF",
					Team:     "LAD",
				},
				Outcomes: []market.Outcome{
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
				Description: "Walker Buehler 6.5 Strikeouts",
				Status:      market.StatusOpen,
				Type:        market.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &market.Player{
					ID:       "shohei-ohtani",
					Name:     "Shoehei Ohtani",
					Number:   "17",
					Position: "SP",
					Team:     "LAD",
				},
				Outcomes: []market.Outcome{
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
		League:      market.LeagueMLB,
		Description: "Pittsburgh Pirates @ Colorado Rockies",
		Markets: []market.Market{
			{
				ID:          "mlb-pirates-ml",
				EventID:     "mlb-pirates-rockies",
				Description: "Pittsburgh Pirates ML",
				Status:      market.StatusOpen,
				Type:        market.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []market.Outcome{
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
				Status:      market.StatusOpen,
				Type:        market.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []market.Outcome{
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
				Status:      market.StatusOpen,
				Type:        market.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &market.Player{
					ID:       "bryan-reynolds",
					Name:     "Bryan Reynolds",
					Number:   "22",
					Position: "CF",
					Team:     "PIT",
				},
				Outcomes: []market.Outcome{
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
		Type:        market.TypeFuture,
		League:      market.LeagueMLB,
		Description: "New York Yankees",
		Markets: []market.Market{
			{
				ID:          "mlb-yankees-win-world-series",
				EventID:     "mlb-yankees-futures",
				Description: "Win the World Series",
				Status:      market.StatusOpen,
				Type:        market.TypeFuture,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []market.Outcome{
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
		Type:        market.TypeFuture,
		League:      market.LeagueMLB,
		Description: "Los Angeles Dodgers",
		Markets: []market.Market{
			{
				ID:          "mlb-dodgers-win-world-series",
				EventID:     "mlb-dodgers-futures",
				Description: "Win the World Series",
				Status:      market.StatusOpen,
				Type:        market.TypeFuture,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []market.Outcome{
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
