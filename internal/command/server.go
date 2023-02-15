package command

import (
	"github.com/urfave/cli/v2"
)

var Server = &cli.Command{
	Name:  "server",
	Usage: "BingGPT server api or ws",
	Action: func(c *cli.Context) error {
		return nil
	},
}
