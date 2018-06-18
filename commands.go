package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"io/ioutil"
)

var Commands = []cli.Command{
	{
		Name:      "groups",
		ShortName: "g",
		Usage:     "List all groups and teams in the 2018 FIFA World Cup.",
		Action:    ListGroups,
	},
	{
		Name:      "standings",
		ShortName: "s",
		Usage:     "Get current standings for all the groups in the 2018 FIFA World Cup.",
		Action:    ListStandings,
	},
	{
		Name:      "stadiums",
		ShortName: "std",
		Usage:     "Get list of all stadiums to be used in the 2018 FIFA World Cup.",
		Action:    ListStadiums,
	},
	{
		Name:      "teams",
		ShortName: "t",
		Usage:     "Get list of all teams participating in the 2018 FIFA World Cup.",
		Action:    ListTeams,
	},
	{
		Name:      "feed",
		ShortName: "f",
		Usage:     "Get all the feed from the 2018 FIFA World Cup including past matches, upcoming matches,results, goal scorers etc.",
		Action:    ListRounds,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "date, d",
				Usage: "Filter feed by date. Format yyyy-mm-dd e.g 2016-06-18. You can use -d today , for today's feed.",
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}

func GetJson(url string) ([]byte, error) {
	r, err := netClient.Get(url)

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}

	return []byte(b), nil
}
