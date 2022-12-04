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

// LookupMXFunc returns a function to look up DNS MX records for a domain name
func LookupMXFunc(name string) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		mxs, err := net.LookupMX(ctx.String(name))
		if err != nil {
			return err
		}
		for mx := range mxs {
			fmt.Println(mxs[mx].Host)
		}
		return nil
	}
}
