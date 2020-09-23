package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	
	"github.com/urfave/cli"

	"gcli/names"
	"gcli/config"
)

/**
 * gcli clone <name> : default --all
 * gcli clone --all
 * gcli clone --dev
 * gcli clone admn-panel
 *
 * gcli cluster create : default --type=kind
 * gcli cluster create --name <name>
 * gcli cluster delete --name <name> --type <type>
 * gcli cluster delete --name <name>
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
						Action: createNewCluster,
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
							&cli.StringFlag{
								Name: "type",
								Usage: "Kubernetes type",
							},
						},
						/** End of Flags **/
					},
					{
						Name: "delete",
						Usage: "Delete Cluster",
						Category: "cluster",
						Action: deleteCluster,
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
						Action: createNewNamespace,
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
						Action: deleteNamespace,
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
						Action: createNewDatabase,
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

			// Helm command
			{
				Name: "helm",
				Usage: "Add helm repo",
				Description: "add and update helm repos",
				Action: helm,
				/** Flags **/
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name: "all",
						Usage: "Apply all",
					},
				},
				/** End of Flags **/
			},
			// End of Hellm command

			/** End of Commands **/
		},
	}

	// Run app
	err := app.Run(os.Args)
  	if err != nil {
    	log.Fatal(err)
  	}
}

// Cloning
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

// Clone helper function
func cloneAllHelper(dir string) {
	var errors = 0
	var errorRepos = []string{}
	var errorIs = 0

	for i, s := range names.REPO_LIST {
		cmd := exec.Command("git", "clone", s)
		err := cmd.Run()
		if err != nil {
			// add error repo to chest
			errorRepos = append(errorRepos, s)
			errorIs = i
		}
	} 
	
	// Print Error or Done
	if errors == 0 {
		fmt.Println("Cloning Done.")
	}else {
		fmt.Println("Error: please remove the cloned repositories and run this comand agin", "\n")
		fmt.Println(errorRepos, errorIs, "\n")
	}
}

func cloneGitRepositories(c *cli.Context) error {
	if c.Bool("all") {
		fmt.Fprintf(c.App.Writer, "Start cloning", "\n")
		for s := range names.REPO_LIST {
			fmt.Println(s)
		}
	} else if c.Bool("dev") {
		fmt.Fprintf(c.App.Writer, "Start cloning", "\n")
		cmd := exec.Command("git", "clone", "https://github.com/Guya-LTD/guya-dev")
		err := cmd.Run()
		if err == nil {
			// Proceed
			cloneAllHelper(names.DEV_FOLDER_NAME)
		}
	}
	return nil
}

// Cluster
func createNewCluster(c *cli.Context) error {
	if c.String("name") != "" && c.Bool("all"){
		fmt.Fprintf(c.App.Writer, "Error: invalid command --name and --all doesnot go together", "\n")
	} else if c.Bool("all") && c.String("type") == "kind" {
		// Create all cluster
		fmt.Fprintf(c.App.Writer, "Creating New Cluster", "\n")
		cmd := exec.Command("kind", "create", "cluster", "--name", names.CLUSTER_NAME, "--config", config.KIND_CLUSTER_CONFIG)
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(c.App.Writer, "Error: Failed to create kind cluster", "\n", err, "\n")
		} else {
			fmt.Fprintf(c.App.Writer, "Done Kind cluster created", "\n")
		}
	}
	return nil
}

func deleteCluster(c *cli.Context) error {
	if c.Bool("all") {
		cmd := exec.Command("kind", "delete", "cluster", "--name", names.CLUSTER_NAME)
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(c.App.Writer, "Error: failed to delte clusters", "\n")
		} else {
			fmt.Fprintf(c.App.Writer, "Done clusters deleted", "\n")
		}
	}
	return nil
}

// Namespace
func createNewNamespace(c *cli.Context) error {
	if c.Bool("all") {
		cmd1 := exec.Command("kubectl", "create", "ns", names.GUYA_NAMESPACE)
		err1 := cmd1.Run()
		cmd2 := exec.Command("kubectl", "create", "ns", names.GUYA_ELK_NAMESPACE)
		err2 := cmd2.Run()

		if err1 != nil || err2 != nil {
			fmt.Fprintf(c.App.Writer, "Error: Failed to create namespaces, run delte namespace before running agin", "\n")
		} else {
			fmt.Fprintf(c.App.Writer, "Done namespaces created", "\n")
		}
	}
	return nil
}

func deleteNamespace(c *cli.Context) error {
	if c.Bool("all") {
		cmd := exec.Command("kubectl", "delete", "ns", names.GUYA_NAMESPACE, names.GUYA_ELK_NAMESPACE)
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(c.App.Writer, "Error: Failed to delete namespaces", "\n")
		} else {
			fmt.Fprintf(c.App.Writer, "Done namespaces deleted", "\n")
		}
	}
	return nil
}

// Database
func createNewDatabase(c *cli.Context) error {
	if c.String("name") != "" {
		//cmd := exec.Command("kubec")
	}
	return nil
}

func helm(c *cli.Context) error {
	if c.Bool("all") {
		exec.Command("ffmpeg", "helm", "repo", "add", "bitnami", "https://charts.bitnami.com/bitnami").Run()
		//exec.Command("helm", "repo", "update").Run()
		fmt.Println("Done", "\n")
	}
	return nil
}