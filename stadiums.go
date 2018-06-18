package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"strconv"
)

func ListStadiums(c *cli.Context) {

	results := APIResponse{}

	res, err := GetJson(API.GetStadiums)

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

	table.SetHeader([]string{"NAME", "CAPACITY", "CITY", "TIMEZONE"})

	for _, stadium := range results.Stadiums {

		table.Append([]string{stadium.Name, strconv.Itoa(stadium.Capacity), stadium.City, stadium.Timezone})

	}

	table.Render()

}
