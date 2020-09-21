package main

import (
	"fmt"
	"log"
	"strings"
	"os"
	"os/exec"

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
	repos := []string{
		"https://github.com/Guya-LTD/guya.git",
		"https://github.com/Guya-LTD/bits.git",
		"https://github.com/Guya-LTD/gxdriver.git",
		"https://github.com/Guya-LTD/user.git",
		"https://github.com/Guya-LTD/gatekeeper.git",
		"https://github.com/Guya-LTD/catalog.git",
		"https://github.com/Guya-LTD/dymo.git",
		"https://github.com/Guya-LTD/gcss.git",
		"https://github.com/Guya-LTD/xpress.git",
		"https://github.com/Guya-LTD/branch.git",
		"https://github.com/Guya-LTD/chipmunk.git",
		"https://github.com/Guya-LTD/postgres.git",
		"https://github.com/Guya-LTD/shop-web.git",
		"https://github.com/Guya-LTD/xpress-web.git",
		"https://github.com/Guya-LTD/nginx.git",
		"https://github.com/Guya-LTD/redis.git",
		"https://github.com/Guya-LTD/refme.git",
		"https://github.com/Guya-LTD/xtrack.git",
		"https://github.com/Guya-LTD/admin-panel.git",
		"https://github.com/Guya-LTD/alfa-geez-node.git",
		"https://github.com/Guya-LTD/chat.git",
		"https://github.com/Guya-LTD/cart.git",
		"https://github.com/Guya-LTD/pyrat.git",
		"https://github.com/Guya-LTD/payment.git",
		"https://github.com/Guya-LTD/storybook.git",
		"https://github.com/Guya-LTD/python-logstash.git",
		"https://github.com/Guya-LTD/gcli.git"
	}

	fmt.Fprintf(c.App.Writer, "Start Cloning, --all\n")

	if c.Bool("all") {
		ab := exec.Command("git", "clone", "https://github.com/Guya-LTD/guya-dev")
		e := ab.Run()
		for i, s := range repos {
			cmd := exec.Command("git", "clone", s)
			err := cmd.Run()
			if err != nil {
				// Some thing went wrong
				// Role back
				fmt.Fprintf(c.App.Writer, "Error occured while cloning, rolling back...", i, "\n")
				rollbackCloning(repos)
			}
		}
		return cli.Exit("Anot b", 84)
	}

	if c.Bool("dev") {
		exec.Command("git", "clone", "https://github.com/Guya-LTD/guya-dev")
		for i, s := range repos {
			cmd := exec.Command("git", "clone", s, "guya-dev")
			err := cmd.Run()
			if err != nil {
				// Some thing went wrong
				// Role back
				fmt.Fprintf(c.App.Writer, "Error occured while cloning, rolling back...", i, "\n")
				rollbackCloning(repos)
			}
		}
		return cli.Exit("Anot b", 84)

	}

	return nil
}

func after(value string, a string) string {
    // Get substring after a string.
    pos := strings.LastIndex(value, a)
    if pos == -1 {
        return ""
    }
    adjustedPos := pos + len(a)
    if adjustedPos >= len(value) {
        return ""
    }
    return value[adjustedPos:len(value)]
}

func rollbackCloning(repos []string) {
	for i, s := range repos {
		st := after(s, "http://github.com/Guya-LTD/")
		cmd := exec.Command("rm", "-R", st)
		err := cmd.Run()
		if err != nil {
			// Failed to rollback
			fmt.Println(i, "Failed to clean up directory, manually remove folders before running this command agina", "\n")
		}
	}
}
