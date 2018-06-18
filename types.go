package main

type Team struct {
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Continent string    `json:"continent,omitempty"`
	Assoc     TeamAssoc `json:"assoc,omitempty"`
}

type TeamAssoc struct {
	Key         string               `json:"key"`
	Name        string               `json:"name"`
	Continental TeamAssocContinental `json:"continental"`
}

type TeamAssocContinental struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Group struct {
	Name      string      `json:"name"`
	Teams     []Team      `json:"teams"`
	Standings []Standings `json:"standings,omitempty"`
}

type APIResponse struct {
	Name     string    `json:"name"`
	Groups   []Group   `json:"groups,omitempty"`
	Stadiums []Stadium `json:"stadiums,omitempty"`
	Teams    []Team    `json:"teams,omitempty"`
	Rounds   []Round   `json:"rounds,omitempty"`
}

type Round struct {
	Name    string    `json:"name"`
	Matches []Matches `json:"matches"`
}

type Matches struct {
	Num      int     `json:"num,omitempty"`
	Date     string  `json:"date,omitempty"`
	Time     string  `json:"time,omitempty"`
	TeamOne  Team    `json:"team1,omitempty"`
	TeamTwo  Team    `json:"team2,omitempty"`
	Score1   int     `json:"score1,omitempty"`
	Score2   int     `json:"score2,omitempty"`
	Score1i  int     `json:"score1i,omitempty"`
	Score2i  int     `json:"score2i,omitempty"`
	Goals1   []Goal  `json:"goals1,omitempty"`
	Goals2   []Goal  `json:"goals2,omitempty"`
	Group    string  `json:"group,omitempty"`
	City     string  `json:"city,omitempty"`
	Timezone string  `json:"timezone,omitempty"`
	Stadium  Stadium `json:"stadium,omitempty"`
}

type Goal struct {
	Name   string `json:"name"`
	Minute int    `json:"minute"`
	Score1 int    `json:"score1"`
	Score2 int    `json:"score2"`
}

type Standings struct {
	Team         Team `json:"team"`
	Position     int  `json:"pos"`
	Played       int  `json:"played"`
	Won          int  `json:"won"`
	Drawn        int  `json:"drawn"`
	Lost         int  `json:"lost"`
	GoalsFor     int  `json:"goals_for"`
	GoalsAgainst int  `json:"goals_against"`
	Pts          int  `json:"pts"`
}

type Stadium struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	City     string `json:"city"`
	Timezone string `json:"timezone"`
}
