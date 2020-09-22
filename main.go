package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

/**
 * gcli clone <name> : default --all
 * gcli clone --all
 * gcli clone --dev
 * gcli clone admn-panel
 *
 * gcli cluster create : default --type=kind
 * gcli cluster delete --type=<name>
 *
 * gcli namespace create --name <name> : default --all
 * gcli namespace create --all
 * gcli namespace create guya-ltd 
 * gcli namespace delete --name <name> : default all
 * gcli namespace delete --all
 * gcli namespace delete guyya-ltd
 *
 * gcli database create <name> : default --all
 * gcli database create --all
 * gcli database create userdb
 * gcli database delete <name> : default --all
 * gcli database delete --all
 * gcli database delete --all --prune
 * gcli database delete userdb
 * 
 * gcli pvc delete : default --all
 * gcli pvc delete --name <name>
 */

func main() {
	app := &cli.App{
		// For bash default auto completion
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			/** Commands **/

			// Clone command
			{
				Name: "clone",
				Usage: "Clone repositories",
				UsageText: "gcli clone [flag], [repository name]",
				Description: "Clone Guya Microservices git repository",
				Action: cloneGitRepositories,
				/** Flags **/
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name: "all",
						Usage: "Apply all",
					},
					&cli.BoolFlag{
						Name: "dev",
						Usage: "Apply development env",
					},
					&cli.BoolFlag{
						Name: "prod",
						Usage: "Apply production env",
					},
					&cli.BoolFlag{
						Name: "sta",
						Usage: "Apply staging env",
					},
				},
				/** End of Flags **/
			},
			// End of Clone command

			// Cluster command
			{
				Name: "cluster",
				Usage: "For Managing kubernetes cluster",
				Description: "Create kubernetes cluster based on the type",
				Action: cloneGitRepositories,
				Subcommands: []*cli.Command{
					{
						Name: "create",
						Usage: "Create new Cluster",
						Category: "cluster",
						//Action: createNewCluster,
						/** Flags **/
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name: "name",
								Usage: "Name of the cluster",
							},
							&cli.BoolFlag{
								Name: "all",
								Usage: "Apply all",
							},
						},
						/** End of Flags **/
					},
					{
						Name: "delete",
						Usage: "Delete Cluster",
						Category: "cluster",
						//Action: deleteCluster,
						/** Flags **/
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name: "name",
								Usage: "Name of the cluster",
							},
							&cli.BoolFlag{
								Name: "all",
								Usage: "Apply all",
							},
						},
						/** End of Flags **/
					},
				},
			},
			// End of Cluster command

			// Namespace command
			{
				Name: "namespace",
				Usage: "For Managing kubernetes namespace",
				Description: "Create kubernetes namespace",
				Action: cloneGitRepositories,
				Subcommands: []*cli.Command{
					{
						Name: "create",
						Usage: "Create new Cluster",
						Category: "namespace",
						//Action: createNewCluster,
						/** Flags **/
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name: "name",
								Usage: "Namespace name",
							},
							&cli.BoolFlag{
								Name: "all",
								Usage: "Apply all",
							},
						},
						/** End of Flags **/
					},
					{
						Name: "delete",
						Usage: "Delete Cluster",
						Category: "namespace",
						//Action: deleteCluster,
						/** Flags **/
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name: "name",
								Usage: "Namespace name",
							},
							&cli.BoolFlag{
								Name: "all",
								Usage: "Apply all",
							},
						},
						/** End of Flags **/
					},
				},
			},
			// End of Namespace command

			// Database command
			{
				Name: "database",
				Usage: "For Managing kubernetes databases",
				Description: "Create kubernetes database",
				Action: cloneGitRepositories,
				Subcommands: []*cli.Command{
					{
						Name: "create",
						Usage: "Create database",
						Category: "database",
						//Action: createNewCluster,
						/** Flags **/
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name: "name",
								Usage: "Namespace name",
							},
							&cli.BoolFlag{
								Name: "all",
								Usage: "Apply all",
							},
						},
						/** End of Flags **/
					},
					{
						Name: "delete",
						Usage: "Delete Cluster",
						Category: "database",
						//Action: deleteCluster,
						/** Flags **/
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name: "name",
								Usage: "Database name",
							},
							&cli.BoolFlag{
								Name: "prune",
								Usage: "Delete pvc",
							},
							&cli.BoolFlag{
								Name: "all",
								Usage: "Apply all",
							},
						},
						/** End of Flags **/
					},
				},
			},
			// End of Database command

			// pvc command
			{
				Name: "pvc",
				Usage: "Remove presistance volume",
				UsageText: "gcli pvc [flag], [repository name]",
				Description: "Clone Guya Microservices git repository",
				Action: cloneGitRepositories,
				/** Flags **/
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name: "all",
						Usage: "Apply all",
					},
					&cli.StringFlag{
						Name: "name",
						Usage: "name of the pvc",
					},
				},
				/** End of Flags **/
			},
			// End of pvc command

			/** End of Commands **/
		},
	}

	// Run app
	err := app.Run(os.Args)
  	if err != nil {
    	log.Fatal(err)
  	}
}

func cloneGitRepositories(c *cli.Context) error {
	if c.Bool("all") {
		fmt.Fprintf(c.App.Writer, "all")
	} else if c.Bool("dev") {
		fmt.Fprintf(c.App.Writer, "dev")
	}
	return nil
}