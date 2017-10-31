package main

import (
	"fmt"
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
			Name:  "memory, m",
			Usage: "Memory Limit --memory 300m",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "jar",
			Aliases: []string{"a"},
			Usage:   "Run a jar file",
			Action: func(c *cli.Context) error {
				err := runJarFile(c)

				return err
			},
		},
		{
			Name:    "binary",
			Aliases: []string{"b"},
			Usage:   "Run a binary file",
			Action: func(c *cli.Context) error {
				err := runBinaryFile(c)

				return err
			},
		},
	}

	app.Run(os.Args)
}

func runJarFile(c *cli.Context) error {
	fmt.Println("Run a jar file : ", c.Args().First())
	return nil
}

func runBinaryFile(c *cli.Context) error {
	fmt.Println("Run a binary file: ", c.Args().First())
	return nil
}
