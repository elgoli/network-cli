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
		{
			Name:   "mxLookup",
			Usage:  "Looks up DNS MX records for the given domain name",
			Flags:  []cli.Flag{cli.StringFlag{Name: "host"}},
			Action: command.LookupMXFunc("host"),
		},
		{
			Name:   "nsLookup",
			Usage:  "Looks up DNS NS records for a given domain name",
			Flags:  []cli.Flag{cli.StringFlag{Name: "host"}},
			Action: command.LookupNSFunc("host"),
		},
	}
}
