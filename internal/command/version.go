package command

import (
	"BingGPT/internal/version"
	"fmt"
	"github.com/urfave/cli/v2"
)

var Version = &cli.Command{
	Name:  "version",
	Usage: "print BingGPT version",
	Action: func(c *cli.Context) error {
		fmt.Println("version:", version.Version())
		return nil
	},
}
