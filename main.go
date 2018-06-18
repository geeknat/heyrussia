package main

import (
	"github.com/urfave/cli"
	"os"
	"log"
)

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Usage = Description
	app.Author = Author
	app.Email = Email
	app.Version = Version
	app.Copyright = Copyright

	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
