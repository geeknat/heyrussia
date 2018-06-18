package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"strconv"
)

func ListStandings(c *cli.Context) {

	results := APIResponse{}

	res, err := GetJson(API.GetStandings)

	if err != nil {
		fmt.Println("Something went wrong : " + err.Error())
		os.Exit(2)
	}

	err = json.Unmarshal(res, &results)

	if err != nil {
		fmt.Println("Something went wrong : " + err.Error())
		os.Exit(2)
	}

	for _, group := range results.Groups {

		table := tablewriter.NewWriter(c.App.Writer)

		table.SetHeader([]string{group.Name, "P", "W", "D", "L", "GF", "GA", "PTS"})

		for _, standing := range group.Standings {
			table.Append([]string{standing.Team.Name,
				strconv.Itoa(standing.Played),
				strconv.Itoa(standing.Won),
				strconv.Itoa(standing.Drawn),
				strconv.Itoa(standing.Lost),
				strconv.Itoa(standing.GoalsFor),
				strconv.Itoa(standing.GoalsAgainst),
				strconv.Itoa(standing.Pts)})
		}

		table.Render()

	}

}
