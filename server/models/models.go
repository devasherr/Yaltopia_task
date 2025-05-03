package models

type PrematchData struct {
	Success int `json:"success"`
	Results []struct {
		EventID string `json:"event_id"`
		Main    struct {
			SP struct {
				ToWinTheMatch struct {
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"to_win_the_match"`
			} `json:"sp"`
		} `json:"main"`
		Innings1 struct {
			SP struct {
				FirstInningsScore struct {
					Odds []struct {
						ID     string `json:"id"`
						Odds   string `json:"odds"`
						Name   string `json:"name"`
						Header string `json:"header"`
					} `json:"odds"`
				} `json:"1st_innings_score"`
			} `json:"sp"`
		} `json:"innings_1"`
		Match struct {
			SP struct {
				TeamToMakeHighest1st6OversScore struct {
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"team_to_make_highest_1st_6_overs_score"`
			} `json:"sp"`
		} `json:"match"`
	} `json:"results"`
}

type ResultData struct {
	Success int `json:"success"`
	Results []struct {
		ID         string `json:"id"`
		SportID    string `json:"sport_id"`
		Time       string `json:"time"`
		TimeStatus string `json:"time_status"`
		League     struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"league"`
		Home struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			ImageID string `json:"image_id"`
		} `json:"home"`
		Away struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			ImageID string `json:"image_id"`
		} `json:"away"`
		SS string `json:"ss"`
	} `json:"results"`
}

type VolleyPrematchData struct {
	Success int `json:"success"`
	Results []struct {
		FixtureID string `json:"FI"`
		EventID   string `json:"event_id"`
		Main      struct {
			Sp struct {
				GameLines struct {
					Odds []struct {
						ID       string `json:"id"`
						Name     string `json:"name"`
						Odds     string `json:"odds"`
						Header   string `json:"header"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"game_lines"`
				CorrectSetScore struct {
					Odds []struct {
						ID     string `json:"id"`
						Odds   string `json:"odds"`
						Name   string `json:"name"`
						Header string `json:"header"`
					} `json:"odds"`
				} `json:"correct_set_score"`
			} `json:"sp"`
		} `json:"main"`
		Others []struct {
			Sp struct {
				Set1Lines struct {
					Odds []struct {
						ID       string `json:"id"`
						Odds     string `json:"odds"`
						Name     string `json:"name"`
						Header   string `json:"header"`
						Handicap string `json:"handicap"`
					} `json:"odds"`
				} `json:"set_1_lines,omitempty"`
			} `json:"sp"`
		} `json:"others"`
	} `json:"results"`
}

type VolleyResultData struct {
	Success int `json:"success"`
	Results []struct {
		ID         string `json:"id"`
		Time       string `json:"time"`
		TimeStatus string `json:"time_status"`
		League     struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"league"`
		Home struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"home"`
		Away struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"away"`
		SS     string `json:"ss"`
		Scores map[string]struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"scores"`
	} `json:"results"`
}

type Fixture struct {
	ID       string `json:"id"`
	HomeId   string `json:"homeId"`
	HomeName string `json:"homeName"`
	AwayId   string `json:"awayId"`
	AwayName string `json:"awayName"`
}

type Odd struct {
	ID       string `json:"id"`
	Odds     string `json:"odds"`
	Name     string `json:"name"`
	Header   string `json:"header"`
	Handicap string `json:"handicap"`
}

type CricketRequest struct {
	ID     string `json:"id"`
	Odds   string `json:"odds"`
	Name   string `json:"name"`
	Header string `json:"header"`
}

type VolleyballRequest struct {
	ID       string `json:"id"`
	Odds     string `json:"odds"`
	Name     string `json:"name"`
	Header   string `json:"header"`
	Handicap string `json:"handicap"`
}
