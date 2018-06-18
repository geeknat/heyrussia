package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"encoding/json"
	"github.com/olekukonko/tablewriter"
)


func ListGroups(c *cli.Context) {

	results := APIResponse{}

	res, err := GetJson(API.GetGroups)

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

		table.SetHeader([]string{group.Name})

		for _, team := range group.Teams {
			table.Append([]string{team.Name})
		}

		table.Render()

	}



}
