package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"time"
	"strconv"
)

func ListRounds(c *cli.Context) {

	results := APIResponse{}

	res, err := GetJson(API.GetFeed)

	if err != nil {
		fmt.Println("Something went wrong : " + err.Error())
		os.Exit(2)
	}

	err = json.Unmarshal(res, &results)

	if err != nil {
		fmt.Println("Something went wrong : " + err.Error())
		os.Exit(2)
	}

	dateFilter := c.String("date")

	table := tablewriter.NewWriter(c.App.Writer)

	table.SetHeader([]string{
		"EVENT",
		"DATE",
		"GROUP",
		"HOME",
		"HOME SCORERS",
		"AWAY",
		"AWAY SCORERS",
		"STADIUM"})

	for _, round := range results.Rounds {

		for _, match := range round.Matches {

			if dateFilter == "" {
				table.Append(ParseMatch(round, match))
			} else {

				if dateFilter == "today" {
					dateFilter = time.Now().Format("2006-01-02")
				}

				if match.Date == dateFilter {
					table.Append(ParseMatch(round, match))
				} else {
					table.SetHeader([]string{})
				}
			}

		}

	}

	table.SetAutoMergeCells(true)

	table.SetRowLine(true)

	table.Render()

}

func ParseMatch(round Round, match Matches) []string {

	var teamOneInfo,teamTwoInfo = "",""

	for _, goal := range match.Goals1 {
		teamOneInfo += goal.Name + " '" + strconv.Itoa(goal.Minute) + "\n"
	}

	for _, goal := range match.Goals2 {
		teamTwoInfo += goal.Name + " '" + strconv.Itoa(goal.Minute) + "\n"
	}

	//for _, goal := range match.Goals2 {
	//	teamOneInfo += goal.Name + " '" + strconv.Itoa(goal.Minute) + "\n"
	//}

	return []string{
		round.Name,
		match.Date+"\n"+match.Time,
		match.Group,
		match.TeamOne.Name + "\n" + strconv.Itoa(match.Score1),
		teamOneInfo,
		match.TeamTwo.Name + "\n" + strconv.Itoa(match.Score2),
		teamTwoInfo,
		match.Stadium.Name,
	}
}
