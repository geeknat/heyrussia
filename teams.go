package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"encoding/json"
	"github.com/olekukonko/tablewriter"
)

func ListTeams(c *cli.Context) {

	results := APIResponse{}

	res, err := GetJson(API.GetTeams)

	if err != nil {
		fmt.Println("Something went wrong : " + err.Error())
		os.Exit(2)
	}

	err = json.Unmarshal(res, &results)

	if err != nil {
		fmt.Println("Something went wrong : " + err.Error())
		os.Exit(2)
	}

	table := tablewriter.NewWriter(c.App.Writer)

	table.SetHeader([]string{"CONTINENT", "TEAM", "CODE", "COUNTRY ASSOC.", "CONTINENTAL ASSOC."})

	for _, team := range results.Teams {

		table.Append([]string{team.Continent, team.Name, team.Code, team.Assoc.Name, team.Assoc.Continental.Name})

	}

	table.Render()

}
