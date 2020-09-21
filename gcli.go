package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {

	cli.VersionFlag = &cli.BoolFlag{
		Name: "Print version",
		Aliases: []string{"v"},
		Usage: "Print only the version number",
	}

	app := &cli.App{
		Name: "Guya CLI",
		Usage: "Runn Guya Microservices",
		Version: "v0.1.0",
		Commands: []*cli.Command{
			{
				Name: "clone",
				Aliases: []string{"c"},
				Usage: "Clone repositories",
				UsageText: "gcli clone [flag], [repository name]",
				Description: "Clone Guya Microservices git repository",
				ArgsUsage: "[myflas]",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name: "all",
						Aliases: []string{"a"},
						Usage: "Clone all repositories to the current directory",
					},
					&cli.BoolFlag{
						Name: "dev",
						Aliases: []string{"d"},
						Usage: "Clone development environemnt repository",
					},
				},
				Action: cloneGitRepositories,
			 },
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Wellcome to Guya Microservices")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}


func cloneGitRepositories(c *cli.Context) error {
	fmt.Fprintf(c.App.Writer, "|---> Start Cloning, --all\n")

	if c.Bool("all") {
		return cli.Exit("Anot b", 84)
	}

	if c.Bool("dev") {

	}

	return nil
}
