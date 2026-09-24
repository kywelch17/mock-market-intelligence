package data

import (
	"github.com/kywelch17/mock-market-intelligence/internal/models"
	"math/rand"
)

var NBAEvents = []models.Event{
	{
		ID:          "nba-celtics-lakers",
		Type:        "Game",
		League:      models.LeagueNBA,
		Description: "Boston Celtics @ Los Angeles Lakers",
		Markets: []models.Market{
			{
				ID:          "nba-celtics-lakers-ml",
				EventID:     "nba-celtics-lakers",
				Description: "Celtics vs Lakers Moneyline",
				Status:      models.StatusOpen,
				Type:        models.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-celtics-win",
						Description: "BOS",
						Index:       0,
					},
					{
						ID:          "nba-lakers-win",
						Description: "LAL",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-celtics-lakers-spread",
				EventID:     "nba-celtics-lakers",
				Description: "Celtics -5.5",
				Status:      models.StatusOpen,
				Type:        models.TypeSpread,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-celtics-minus-5-5",
						Description: "Celtics -5.5",
						Index:       0,
					},
					{
						ID:          "nba-lakers-plus-5-5",
						Description: "Lakers +5.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-celtics-lakers-total",
				EventID:     "nba-celtics-lakers",
				Description: "Total Points 228.5",
				Status:      models.StatusOpen,
				Type:        models.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-over-228-5",
						Description: "Over 228.5",
						Index:       0,
					},
					{
						ID:          "nba-under-228-5",
						Description: "Under 228.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-jayson-tatum-points",
				EventID:     "nba-celtics-lakers",
				Description: "Jayson Tatum 32.5 Points",
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "jayson-tatum",
					Name:     "Jayson Tatum",
					Number:   "0",
					Position: "SF",
					Team:     "BOS",
				},
				Outcomes: []models.Outcome{
					{
						ID:          "nba-jayson-tatum-over-32-5-points",
						Description: "Over 32.5",
						Index:       0,
					},
					{
						ID:          "nba-jayson-tatum-under-32-5-points",
						Description: "Under 32.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nba-warriors-nuggets",
		Type:        "Game",
		League:      models.LeagueNBA,
		Description: "Golden State Warriors @ Denver Nuggets",
		Markets: []models.Market{
			{
				ID:          "nba-warriors-nuggets-ml",
				EventID:     "nba-warriors-nuggets",
				Description: "Warriors vs Nuggets Moneyline",
				Status:      models.StatusOpen,
				Type:        models.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-warriors-win",
						Description: "GSW",
						Index:       0,
					},
					{
						ID:          "nba-nuggets-win",
						Description: "DEN",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-warriors-nuggets-spread",
				EventID:     "nba-warriors-nuggets",
				Description: "Warriors -3.0",
				Status:      models.StatusOpen,
				Type:        models.TypeSpread,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-warriors-minus-3-0",
						Description: "Warriors -3.0",
						Index:       0,
					},
					{
						ID:          "nba-nuggets-plus-3-0",
						Description: "Nuggets +3.0",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-warriors-nuggets-total",
				EventID:     "nba-warriors-nuggets",
				Description: "Total Points 234.5",
				Status:      models.StatusOpen,
				Type:        models.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-over-234-5",
						Description: "Over 234.5",
						Index:       0,
					},
					{
						ID:          "nba-under-234-5",
						Description: "Under 234.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-stephen-curry-3pm",
				EventID:     "nba-warriors-nuggets",
				Description: "Stephen Curry 4.5 3-Pointers Made",
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "stephen-curry",
					Name:     "Stephen Curry",
					Number:   "30",
					Position: "PG",
					Team:     "GSW",
				},
				Outcomes: []models.Outcome{
					{
						ID:          "nba-stephen-curry-over-4-5-3pm",
						Description: "Over 4.5",
						Index:       0,
					},
					{
						ID:          "nba-stephen-curry-under-4-5-3pm",
						Description: "Under 4.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nba-heat-knicks",
		Type:        "Game",
		League:      models.LeagueNBA,
		Description: "Miami Heat @ New York Knicks",
		Markets: []models.Market{
			{
				ID:          "nba-heat-knicks-ml",
				EventID:     "nba-heat-knicks",
				Description: "Heat vs Knicks Moneyline",
				Status:      models.StatusOpen,
				Type:        models.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-heat-win",
						Description: "MIA",
						Index:       0,
					},
					{
						ID:          "nba-knicks-win",
						Description: "NYK",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-heat-knicks-spread",
				EventID:     "nba-heat-knicks",
				Description: "Heat -4.5",
				Status:      models.StatusOpen,
				Type:        models.TypeSpread,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-heat-minus-4-5",
						Description: "Heat -4.5",
						Index:       0,
					},
					{
						ID:          "nba-knicks-plus-4-5",
						Description: "Knicks +4.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-heat-knicks-total",
				EventID:     "nba-heat-knicks",
				Description: "Total Points 212.5",
				Status:      models.StatusOpen,
				Type:        models.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-over-212-5",
						Description: "Over 212.5",
						Index:       0,
					},
					{
						ID:          "nba-under-212-5",
						Description: "Under 212.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-bam-adebayo-points",
				EventID:     "nba-heat-knicks",
				Description: "Bam Adebayo 18.5 Points",
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "bam-adebayo",
					Name:     "Bam Adebayo",
					Number:   "13",
					Position: "C",
					Team:     "MIA",
				},
				Outcomes: []models.Outcome{
					{
						ID:          "nba-bam-adebayo-over-18-5-points",
						Description: "Over 18.5",
						Index:       0,
					},
					{
						ID:          "nba-bam-adebayo-under-18-5-points",
						Description: "Under 18.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nba-mavericks-suns",
		Type:        "Game",
		League:      models.LeagueNBA,
		Description: "Dallas Mavericks @ Phoenix Suns",
		Markets: []models.Market{
			{
				ID:          "nba-mavericks-suns-ml",
				EventID:     "nba-mavericks-suns",
				Description: "Mavericks vs Suns Moneyline",
				Status:      models.StatusOpen,
				Type:        models.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-mavericks-win",
						Description: "DAL",
						Index:       0,
					},
					{
						ID:          "nba-suns-win",
						Description: "PHX",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-mavericks-suns-spread",
				EventID:     "nba-mavericks-suns",
				Description: "Mavericks -2.5",
				Status:      models.StatusOpen,
				Type:        models.TypeSpread,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-mavericks-minus-2-5",
						Description: "Mavericks -2.5",
						Index:       0,
					},
					{
						ID:          "nba-suns-plus-2-5",
						Description: "Suns +2.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-mavericks-suns-total",
				EventID:     "nba-mavericks-suns",
				Description: "Total Points 226.5",
				Status:      models.StatusOpen,
				Type:        models.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-over-226-5",
						Description: "Over 226.5",
						Index:       0,
					},
					{
						ID:          "nba-under-226-5",
						Description: "Under 226.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-cooper-flagg-points",
				EventID:     "nba-mavericks-suns",
				Description: "Cooper Flagg 21.5 Points",
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "cooper-flagg",
					Name:     "Cooper Flagg",
					Number:   "32",
					Position: "F",
					Team:     "DAL",
				},
				Outcomes: []models.Outcome{
					{
						ID:          "nba-cooper-flagg-over-21-5-points",
						Description: "Over 21.5",
						Index:       0,
					},
					{
						ID:          "nba-cooper-flagg-under-21-5-points",
						Description: "Under 21.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nba-76ers-bulls",
		Type:        "Game",
		League:      models.LeagueNBA,
		Description: "Philadelphia 76ers @ Chicago Bulls",
		Markets: []models.Market{
			{
				ID:          "nba-76ers-bulls-ml",
				EventID:     "nba-76ers-bulls",
				Description: "76ers vs Bulls Moneyline",
				Status:      models.StatusOpen,
				Type:        models.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-76ers-win",
						Description: "PHI",
						Index:       0,
					},
					{
						ID:          "nba-bulls-win",
						Description: "CHI",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-76ers-bulls-spread",
				EventID:     "nba-76ers-bulls",
				Description: "76ers -6.0",
				Status:      models.StatusOpen,
				Type:        models.TypeSpread,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-76ers-minus-6-0",
						Description: "76ers -6.0",
						Index:       0,
					},
					{
						ID:          "nba-bulls-plus-6-0",
						Description: "Bulls +6.0",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-76ers-bulls-total",
				EventID:     "nba-76ers-bulls",
				Description: "Total Points 220.5",
				Status:      models.StatusOpen,
				Type:        models.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-over-220-5",
						Description: "Over 220.5",
						Index:       0,
					},
					{
						ID:          "nba-under-220-5",
						Description: "Under 220.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-joel-embiid-points",
				EventID:     "nba-76ers-bulls",
				Description: "Joel Embiid 33.5 Points",
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "joel-embiid",
					Name:     "Joel Embiid",
					Number:   "21",
					Position: "C",
					Team:     "PHI",
				},
				Outcomes: []models.Outcome{
					{
						ID:          "nba-joel-embiid-over-33-5-points",
						Description: "Over 33.5",
						Index:       0,
					},
					{
						ID:          "nba-joel-embiid-under-33-5-points",
						Description: "Under 33.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nba-clippers-timberwolves",
		Type:        "Game",
		League:      models.LeagueNBA,
		Description: "Los Angeles Clippers @ Minnesota Timberwolves",
		Markets: []models.Market{
			{
				ID:          "nba-clippers-timberwolves-ml",
				EventID:     "nba-clippers-timberwolves",
				Description: "Clippers vs Timberwolves Moneyline",
				Status:      models.StatusOpen,
				Type:        models.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-clippers-win",
						Description: "LAC",
						Index:       0,
					},
					{
						ID:          "nba-timberwolves-win",
						Description: "MIN",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-clippers-timberwolves-spread",
				EventID:     "nba-clippers-timberwolves",
				Description: "Clippers -4.5",
				Status:      models.StatusOpen,
				Type:        models.TypeSpread,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-clippers-minus-4-5",
						Description: "Clippers -4.5",
						Index:       0,
					},
					{
						ID:          "nba-timberwolves-plus-4-5",
						Description: "Timberwolves +4.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-clippers-timberwolves-total",
				EventID:     "nba-clippers-timberwolves",
				Description: "Total Points 218.5",
				Status:      models.StatusOpen,
				Type:        models.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-over-218-5",
						Description: "Over 218.5",
						Index:       0,
					},
					{
						ID:          "nba-under-218-5",
						Description: "Under 218.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-darius-garland-points",
				EventID:     "nba-clippers-timberwolves",
				Description: "Darius Garland 20.5 Points",
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "darius-garland",
					Name:     "Darius Garland",
					Number:   "10",
					Position: "PG",
					Team:     "LAC",
				},
				Outcomes: []models.Outcome{
					{
						ID:          "nba-darius-garland-over-20-5-points",
						Description: "Over 20.5",
						Index:       0,
					},
					{
						ID:          "nba-darius-garland-under-20-5-points",
						Description: "Under 20.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nba-heat-bucks",
		Type:        "Game",
		League:      models.LeagueNBA,
		Description: "Miami Heat @ Milwaukee Bucks",
		Markets: []models.Market{
			{
				ID:          "nba-heat-bucks-ml",
				EventID:     "nba-heat-bucks",
				Description: "Heat vs Bucks Moneyline",
				Status:      models.StatusOpen,
				Type:        models.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-heat-win",
						Description: "MIA",
						Index:       0,
					},
					{
						ID:          "nba-bucks-win",
						Description: "MIL",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-heat-bucks-spread",
				EventID:     "nba-heat-bucks",
				Description: "Bucks -8.5",
				Status:      models.StatusOpen,
				Type:        models.TypeSpread,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-heat-minus-8-5",
						Description: "Heat -8.5",
						Index:       0,
					},
					{
						ID:          "nba-bucks-plus-8-5",
						Description: "Bucks +8.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-heat-bucks-total",
				EventID:     "nba-heat-bucks",
				Description: "Total Points 224.5",
				Status:      models.StatusOpen,
				Type:        models.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-over-224-5",
						Description: "Over 224.5",
						Index:       0,
					},
					{
						ID:          "nba-under-224-5",
						Description: "Under 224.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-tyler-herro-points",
				EventID:     "nba-heat-bucks",
				Description: "Tyler Herro 22.5 Points",
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "tyler-herro",
					Name:     "Tyler Herro",
					Number:   "14",
					Position: "SG",
					Team:     "MIL",
				},
				Outcomes: []models.Outcome{
					{
						ID:          "nba-tyler-herro-over-22-5-points",
						Description: "Over 22.5",
						Index:       0,
					},
					{
						ID:          "nba-tyler-herro-under-22-5-points",
						Description: "Under 22.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nba-kings-pistons",
		Type:        "Game",
		League:      models.LeagueNBA,
		Description: "Sacramento Kings @ Detroit Pistons",
		Markets: []models.Market{
			{
				ID:          "nba-kings-pistons-ml",
				EventID:     "nba-kings-pistons",
				Description: "Kings vs Pistons Moneyline",
				Status:      models.StatusOpen,
				Type:        models.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-kings-win",
						Description: "SAC",
						Index:       0,
					},
					{
						ID:          "nba-pistons-win",
						Description: "DET",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-kings-pistons-spread",
				EventID:     "nba-kings-pistons",
				Description: "Kings -3.5",
				Status:      models.StatusOpen,
				Type:        models.TypeSpread,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-kings-minus-3-5",
						Description: "Kings -3.5",
						Index:       0,
					},
					{
						ID:          "nba-pistons-plus-3-5",
						Description: "Pistons +3.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-kings-pistons-total",
				EventID:     "nba-kings-pistons",
				Description: "Total Points 230.5",
				Status:      models.StatusOpen,
				Type:        models.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-over-230-5",
						Description: "Over 230.5",
						Index:       0,
					},
					{
						ID:          "nba-under-230-5",
						Description: "Under 230.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-domantas-sabonis-rebounds",
				EventID:     "nba-kings-pistons",
				Description: "Domantas Sabonis 12.5 Rebounds",
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "domantas-sabonis",
					Name:     "Domantas Sabonis",
					Number:   "11",
					Position: "C",
					Team:     "SAC",
				},
				Outcomes: []models.Outcome{
					{
						ID:          "nba-domantas-sabonis-over-12-5-rebounds",
						Description: "Over 12.5",
						Index:       0,
					},
					{
						ID:          "nba-domantas-sabonis-under-12-5-rebounds",
						Description: "Under 12.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nba-hornets-magic",
		Type:        "Game",
		League:      models.LeagueNBA,
		Description: "Charlotte Hornets @ Orlando Magic",
		Markets: []models.Market{
			{
				ID:          "nba-hornets-magic-ml",
				EventID:     "nba-hornets-magic",
				Description: "Hornets vs Magic Moneyline",
				Status:      models.StatusOpen,
				Type:        models.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-hornets-win",
						Description: "CHO",
						Index:       0,
					},
					{
						ID:          "nba-magic-win",
						Description: "ORL",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-hornets-magic-spread",
				EventID:     "nba-hornets-magic",
				Description: "Hornets -6.5",
				Status:      models.StatusOpen,
				Type:        models.TypeSpread,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-hornets-minus-6-5",
						Description: "Hornets -6.5",
						Index:       0,
					},
					{
						ID:          "nba-magic-plus-6-5",
						Description: "Magic +6.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-hornets-magic-total",
				EventID:     "nba-hornets-magic",
				Description: "Total Points 210.5",
				Status:      models.StatusOpen,
				Type:        models.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-over-210-5",
						Description: "Over 210.5",
						Index:       0,
					},
					{
						ID:          "nba-under-210-5",
						Description: "Under 210.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-brandon-miller-points",
				EventID:     "nba-hornets-magic",
				Description: "Brandon Miller 21.5 Points",
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "brandon-miller",
					Name:     "Brandon Miller",
					Number:   "24",
					Position: "SF",
					Team:     "CHO",
				},
				Outcomes: []models.Outcome{
					{
						ID:          "nba-brandon-miller-over-21-5-points",
						Description: "Over 21.5",
						Index:       0,
					},
					{
						ID:          "nba-brandon-miller-under-21-5-points",
						Description: "Under 21.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nba-thunder-trailblazers",
		Type:        "Game",
		League:      models.LeagueNBA,
		Description: "Oklahoma City Thunder @ Portland Trail Blazers",
		Markets: []models.Market{
			{
				ID:          "nba-thunder-trailblazers-ml",
				EventID:     "nba-thunder-trailblazers",
				Description: "Thunder vs Trail Blazers Moneyline",
				Status:      models.StatusOpen,
				Type:        models.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-thunder-win",
						Description: "OKC",
						Index:       0,
					},
					{
						ID:          "nba-trailblazers-win",
						Description: "POR",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-thunder-trailblazers-spread",
				EventID:     "nba-thunder-trailblazers",
				Description: "Thunder -4.0",
				Status:      models.StatusOpen,
				Type:        models.TypeSpread,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-thunder-minus-4-0",
						Description: "Thunder -4.0",
						Index:       0,
					},
					{
						ID:          "nba-trailblazers-plus-4-0",
						Description: "Trail Blazers +4.0",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-thunder-trailblazers-total",
				EventID:     "nba-thunder-trailblazers",
				Description: "Total Points 222.5",
				Status:      models.StatusOpen,
				Type:        models.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-over-222-5",
						Description: "Over 222.5",
						Index:       0,
					},
					{
						ID:          "nba-under-222-5",
						Description: "Under 222.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-shai-gilgeous-alexander-points",
				EventID:     "nba-thunder-trailblazers",
				Description: "Shai Gilgeous-Alexander 28.5 Points",
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "shai-gilgeous-alexander",
					Name:     "Shai Gilgeous-Alexander",
					Number:   "2",
					Position: "PG",
					Team:     "OKC",
				},
				Outcomes: []models.Outcome{
					{
						ID:          "nba-shai-gilgeous-alexander-over-28-5-points",
						Description: "Over 28.5",
						Index:       0,
					},
					{
						ID:          "nba-shai-gilgeous-alexander-under-28-5-points",
						Description: "Under 28.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nba-nets-hawks",
		Type:        "Game",
		League:      models.LeagueNBA,
		Description: "Brooklyn Nets @ Atlanta Hawks",
		Markets: []models.Market{
			{
				ID:          "nba-nets-hawks-ml",
				EventID:     "nba-nets-hawks",
				Description: "Nets vs Hawks Moneyline",
				Status:      models.StatusOpen,
				Type:        models.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-nets-win",
						Description: "BKN",
						Index:       0,
					},
					{
						ID:          "nba-hawks-win",
						Description: "ATL",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-nets-hawks-spread",
				EventID:     "nba-nets-hawks",
				Description: "Nets -2.5",
				Status:      models.StatusOpen,
				Type:        models.TypeSpread,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-nets-minus-2-5",
						Description: "Nets -2.5",
						Index:       0,
					},
					{
						ID:          "nba-hawks-plus-2-5",
						Description: "Hawks +2.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-nets-hawks-total",
				EventID:     "nba-nets-hawks",
				Description: "Total Points 225.5",
				Status:      models.StatusOpen,
				Type:        models.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-over-225-5",
						Description: "Over 225.5",
						Index:       0,
					},
					{
						ID:          "nba-under-225-5",
						Description: "Under 225-5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-michael-porter-jr-points",
				EventID:     "nba-nets-hawks",
				Description: "Michael Porter Jr. 22.5 Points",
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "michael-porter-jr",
					Name:     "Michael Porter Jr.",
					Number:   "17",
					Position: "F",
					Team:     "BKN",
				},
				Outcomes: []models.Outcome{
					{
						ID:          "nba-michael-porter-jr-over-22-5-points",
						Description: "Over 22.5",
						Index:       0,
					},
					{
						ID:          "nba-michael-porter-jr-under-22-5-points",
						Description: "Under 22.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nba-lakers-suns",
		Type:        "Game",
		League:      models.LeagueNBA,
		Description: "Los Angeles Lakers @ Phoenix Suns",
		Markets: []models.Market{
			{
				ID:          "nba-lakers-suns-ml",
				EventID:     "nba-lakers-suns",
				Description: "Lakers vs Suns Moneyline",
				Status:      models.StatusOpen,
				Type:        models.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-lakers-win",
						Description: "LAL",
						Index:       0,
					},
					{
						ID:          "nba-suns-win",
						Description: "PHX",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-lakers-suns-spread",
				EventID:     "nba-lakers-suns",
				Description: "Lakers -1.5",
				Status:      models.StatusOpen,
				Type:        models.TypeSpread,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-lakers-minus-1-5",
						Description: "Lakers -1.5",
						Index:       0,
					},
					{
						ID:          "nba-suns-plus-1-5",
						Description: "Suns +1.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-lakers-suns-total",
				EventID:     "nba-lakers-suns",
				Description: "Total Points 232.5",
				Status:      models.StatusOpen,
				Type:        models.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-over-232-5",
						Description: "Over 232.5",
						Index:       0,
					},
					{
						ID:          "nba-under-232-5",
						Description: "Under 232.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-luka-doncic-points",
				EventID:     "nba-lakers-suns",
				Description: "Luka Dončić 32.5 Points",
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "luka-doncic",
					Name:     "Luka Dončić",
					Number:   "77",
					Position: "PG",
					Team:     "LAL",
				},
				Outcomes: []models.Outcome{
					{
						ID:          "nba-luka-doncic-over-32-5-points",
						Description: "Over 32.5",
						Index:       0,
					},
					{
						ID:          "nba-luka-doncic-under-32-5-points",
						Description: "Under 32.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nba-bulls-cavaliers",
		Type:        "Game",
		League:      models.LeagueNBA,
		Description: "Chicago Bulls @ Cleveland Cavaliers",
		Markets: []models.Market{
			{
				ID:          "nba-bulls-cavaliers-ml",
				EventID:     "nba-bulls-cavaliers",
				Description: "Bulls vs Cavaliers Moneyline",
				Status:      models.StatusOpen,
				Type:        models.TypeMoney,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-bulls-win",
						Description: "CHI",
						Index:       0,
					},
					{
						ID:          "nba-cavaliers-win",
						Description: "CLE",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-bulls-cavaliers-spread",
				EventID:     "nba-bulls-cavaliers",
				Description: "Bulls -3.5",
				Status:      models.StatusOpen,
				Type:        models.TypeSpread,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-bulls-minus-3-5",
						Description: "Bulls -3.5",
						Index:       0,
					},
					{
						ID:          "nba-cavaliers-plus-3-5",
						Description: "Cavaliers +3.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-bulls-cavaliers-total",
				EventID:     "nba-bulls-cavaliers",
				Description: "Total Points 218.5",
				Status:      models.StatusOpen,
				Type:        models.TypeTotal,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-over-218-5",
						Description: "Over 218.5",
						Index:       0,
					},
					{
						ID:          "nba-under-218-5",
						Description: "Under 218.5",
						Index:       1,
					},
				},
			},
			{
				ID:          "nba-josh-giddey-points",
				EventID:     "nba-bulls-cavaliers",
				Description: "Josh Giddey 17.5 Points",
				Status:      models.StatusOpen,
				Type:        models.TypePlayer,
				Volume:      1000 + rand.Float64()*999000,
				Player: &models.Player{
					ID:       "josh-giddey",
					Name:     "Josh Giddey",
					Number:   "3",
					Position: "G",
					Team:     "CHI",
				},
				Outcomes: []models.Outcome{
					{
						ID:          "nba-josh-giddey-over-17-5-points",
						Description: "Over 17.5",
						Index:       0,
					},
					{
						ID:          "nba-josh-giddey-under-17-5-points",
						Description: "Under 17.5",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nba-warriors-future",
		Type:        models.TypeFuture,
		League:      models.LeagueNBA,
		Description: "Golden State Warriors",
		Markets: []models.Market{
			{
				ID:          "nba-warriors-win-championship",
				EventID:     "nba-warriors-future",
				Description: "Win NBA Championship",
				Status:      models.StatusOpen,
				Type:        models.TypeFuture,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-warriors-win-title",
						Description: "Golden State Warriors",
						Index:       0,
					},
					{
						ID:          "nba-warriors-lose-title",
						Description: "Field",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nba-celtics-future",
		Type:        models.TypeFuture,
		League:      models.LeagueNBA,
		Description: "Boston Celtics",
		Markets: []models.Market{
			{
				ID:          "nba-celtics-win-championship",
				EventID:     "nba-celtics-future",
				Description: "Win NBA Championship",
				Status:      models.StatusOpen,
				Type:        models.TypeFuture,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-celtics-win-title",
						Description: "Boston Celtics",
						Index:       0,
					},
					{
						ID:          "nba-celtics-lose-title",
						Description: "Field",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nba-nuggets-future",
		Type:        models.TypeFuture,
		League:      models.LeagueNBA,
		Description: "Denver Nuggets",
		Markets: []models.Market{
			{
				ID:          "nba-nuggets-win-championship",
				EventID:     "nba-nuggets-future",
				Description: "Win NBA Championship",
				Status:      models.StatusOpen,
				Type:        models.TypeFuture,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-nuggets-win-title",
						Description: "Denver Nuggets",
						Index:       0,
					},
					{
						ID:          "nba-nuggets-lose-title",
						Description: "Field",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nba-lakers-future",
		Type:        models.TypeFuture,
		League:      models.LeagueNBA,
		Description: "Los Angeles Lakers",
		Markets: []models.Market{
			{
				ID:          "nba-lakers-win-championship",
				EventID:     "nba-lakers-future",
				Description: "Win NBA Championship",
				Status:      models.StatusOpen,
				Type:        models.TypeFuture,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-lakers-win-title",
						Description: "Los Angeles Lakers",
						Index:       0,
					},
					{
						ID:          "nba-lakers-lose-title",
						Description: "Field",
						Index:       1,
					},
				},
			},
		},
	},
	{
		ID:          "nba-suns-future",
		Type:        models.TypeFuture,
		League:      models.LeagueNBA,
		Description: "Phoenix Suns",
		Markets: []models.Market{
			{
				ID:          "nba-suns-win-championship",
				EventID:     "nba-suns-future",
				Description: "Win NBA Championship",
				Status:      models.StatusOpen,
				Type:        models.TypeFuture,
				Volume:      1000 + rand.Float64()*999000,
				Player:      nil,
				Outcomes: []models.Outcome{
					{
						ID:          "nba-suns-win-title",
						Description: "Phoenix Suns",
						Index:       0,
					},
					{
						ID:          "nba-suns-lose-title",
						Description: "Field",
						Index:       1,
					},
				},
			},
		},
	},
}
