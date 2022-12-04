package command

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// LookupIPFunc returns a functions to look up host IP address using the local resolver
func LookupIPFunc(name string) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		ips, err := net.LookupIP(ctx.String(name))
		if err != nil {
			return err
		}
		for ip := range ips {
			fmt.Println(ips[ip].String())
		}
		return nil
	}
}
