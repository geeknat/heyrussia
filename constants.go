package main

import (
	"net/http"
	"time"
)

const Name = "HeyRussia"
const Version = "1.0.0"
const Description = "Welcome to The FIFA 2018 World Cup CLI.\n\n\n" +
	"Access all the groups, stadiums, teams, standings and live feed from Russia.\n\n\n" +
	"Just call Russia!!!\n\n\n" +
	"See 'heyrussia --help' to see all the available options."

const Author = "Geek Nat"
const Email = "geeknat7@gmail.com"
const Copyright = "(c) 2018 Geek Nat"

const BaseUrl = "https://raw.githubusercontent.com/openfootball/world-cup.json/master/2018/"


type APIInterface struct {
	GetGroups    string
	GetFeed      string
	GetStadiums  string
	GetStandings string
	GetTeams     string
}

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

var API = APIInterface{
	GetGroups:    BaseUrl + "worldcup.groups.json",
	GetFeed:      BaseUrl + "worldcup.json",
	GetStadiums:  BaseUrl + "worldcup.stadiums.json",
	GetStandings: BaseUrl + "worldcup.standings.json",
	GetTeams:     BaseUrl + "worldcup.teams.json",
}
