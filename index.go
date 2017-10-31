package main

import (
	"fmt"
	"os"
	"os/exec"

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

				if err != nil {
					fmt.Errorf("%v", err)
				}

				return err
			},
		},
		{
			Name:    "binary",
			Aliases: []string{"b"},
			Usage:   "Run a binary file",
			Action: func(c *cli.Context) error {
				err := runBinaryFile(c)

				if err != nil {
					fmt.Errorf("%v", err)
				}

				return err
			},
		},
	}

	app.Run(os.Args)
}

func runJarFile(c *cli.Context) error {
	exPath, err := os.Getwd()

	cmd := exec.Command("docker", "run", "--rm", "--network", "none", "-it", "-a", "stdin", "-a", "stdout", "-a", "stderr", "-v", exPath+":/test", "openjdk:latest", "/bin/bash", "-c", "cd /test && java -jar "+c.Args().First())

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()

	return err
}

func runBinaryFile(c *cli.Context) error {
	exPath, err := os.Getwd()

	cmd := exec.Command("docker", "run", "--rm", "--network", "none", "-it", "-a", "stdin", "-a", "stdout", "-a", "stderr", "-v", exPath+":/test", "ubuntu:latest", "/bin/bash", "-c", "cd /test && ./"+c.Args().First())

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()

	return err
}
