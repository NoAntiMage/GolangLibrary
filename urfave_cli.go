package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "wujimaster_cli"
	app.Version = "1.0.0"
	app.Usage = "for_cli_command"
	app.Description = "This is the cli-tool for containerd which is the core module of docker compents"

	var name string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "n,name",
			Value:       "wujimaster",
			Usage:       "say my name",
			Destination: &name,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:     "add",
			Aliases:  []string{"a"},
			Usage:    "calc 1 + 1",
			Category: "math",
			Action: func(c *cli.Context) error {
				fmt.Println("1 + 1 = ", 1+1)
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		if c.String("name") == "Millais" {
			fmt.Println("greetings, my lady", name)
		} else {
			fmt.Println("hello", name)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
