package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"gitlab.com/elgol/network-cli/command"
)

func main() {
	app := cli.NewApp()
	app.Name = "Network CLI"
	app.Commands = getCliCommands()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getCliCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "ipLookup",
			Usage:  "looks up host IP address using the local resolver",
			Flags:  []cli.Flag{cli.StringFlag{Name: "host"}},
			Action: command.LookupIPFunc("host"),
		},
	}
}
