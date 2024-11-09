package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/elnazdev/network-cli/command"
)

func main() {
	app := cli.NewApp()
	app.Name = "ncli"
	app.Commands = getCliCommands()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getCliCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "iplookup",
			Usage:  "looks up host IP address using the local resolver",
			Flags:  []cli.Flag{cli.StringFlag{Name: "host"}},
			Action: command.LookupIPFunc("host"),
		},
		{
			Name:   "mxlookup",
			Usage:  "Looks up DNS MX records for the given domain name",
			Flags:  []cli.Flag{cli.StringFlag{Name: "host"}},
			Action: command.LookupMXFunc("host"),
		},
		{
			Name:   "nslookup",
			Usage:  "Looks up DNS NS records for a given domain name",
			Flags:  []cli.Flag{cli.StringFlag{Name: "host"}},
			Action: command.LookupNSFunc("host"),
		},
	}
}
