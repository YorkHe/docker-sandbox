package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "Docker Sandbox"
	app.Usage = "Run binary or jar file in an isolated docker container with piped stdin/stdout and permission control."

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input, i",
			Usage: "Load input from `FILE`",
		},
		cli.StringFlag{
			Name:  "output, c",
			Usage: "Output are written into `FILE`",
		},
		cli.StringFlag{
			Name:  "binary, b",
			Usage: "Run Binary File -- `FILE`",
		},
		cli.StringFlag{
			Name:  "jar, j",
			Usage: "Run jar file -- `FILE`",
		},
		cli.StringFlag{
			Name:  "memory, m",
			Usage: "Memory Limit --memory 300m",
		},
	}

	app.Action = func(c *cli.Context) error {
		return nil
	}

	app.Run(os.Args)
}
