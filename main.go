package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	
	"github.com/urfave/cli"

	//"github.com/Guya-LTD/gcli/names"
	//"github.com/Guya-LTD/gcli/config"

	"gcli/names"
	"gcli/config"
)

/**
 * gcli clone <name> : default --all
 * gcli clone --all
 * gcli clone --dev
 * gcli clone admn-panel
 *
 * gcli init
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
 * 
 * gcli pv create : default --all
 * gcli pv delete : default --all
 *
 * gcli deployment create --name elk
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
						Action: deleteDatabase,
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

			// Deployment command
			{
				Name: "deployment",
				Usage: "For Managing kubernetes deployments",
				Description: "Create kubernetes deployments",
				Subcommands: []*cli.Command{
					{
						Name: "create",
						Usage: "Create deployments",
						Category: "deployment",
						Action: createNewDeployment,
						/** Flags **/
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name: "name",
								Usage: "Deployment name",
							},
						},
						/** End of Flags **/
					},
					{
						Name: "delete",
						Usage: "Delete deployments",
						Category: "deployment",
						Action: deleteDeployment,
						/** Flags **/
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name: "name",
								Usage: "Deployment name",
							},
						},
						/** End of Flags **/
					},
				},
			},
			// End of Deployment command

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
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
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
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
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
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
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
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
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
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
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
		// kubectl create -f initalize/namespaces
		cmd := exec.Command("kubectl", "create", "-f", names.INITIALIZE_NAMESPACES_DIR)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		fmt.Println(err)
	}
	return nil
}

func deleteNamespace(c *cli.Context) error {
	if c.Bool("all") {
		// kubectl delete -f initalize/namespaces
		cmd := exec.Command("kubectl", "delete", "-f", names.INITIALIZE_NAMESPACES_DIR)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		fmt.Println(err)
	}
	return nil
}

// Database

func createDb(name string, values string, db string) {
	cmd := exec.Command("helm", "install", "--namespace", names.GUYA_NAMESPACE, name, "--version", "9.1.2", db, "--values", values)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	fmt.Println(err)
}

func createDatabase(name string, repo string, values string){
	// helm install --namespace guya-ltd mongodb --version 9.1.2 bitnami/mongodb --values values.yaml
	cmd := exec.Command("helm", "install", "--namespace", names.GUYA_NAMESPACE, name, repo, "--values", values)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	fmt.Println(err)
}

func createNewDatabase(c *cli.Context) error {
	if c.String("name") == "mongodb" && !c.Bool("all") {
		createDatabase(names.MONGODB_NAME, names.MONGODB, names.MONGODB_VALUE)
	}  else {
		fmt.Println("Command Error")
	}
	/*if c.String("name") == "branch" && !c.Bool("all") {
		createDb(names.DATABASE_BRANCH_NAME, names.DATABASE_BRANCH_VALUE, DATABASE_BRANCH_DB)
	} else if c.String("name") == "cart" && !c.Bool("all") { 
		createDb(names.DATABASE_CART_NAME, names.DATABASE_CART_VALUE, DATABASE_CART_DB)
	} else if c.String("name") == "catalog" && !c.Bool("all") { 
		createDb(names.DATABASE_CATALOG_NAME, names.DATABASE_CATALOG_VALUE, DATABASE_CATALOG_DB)
	} else if c.String("name") == "chat" && !c.Bool("all") { 
		createDb(names.DATABASE_CHAT_NAME, names.DATABASE_CHAT_VALUE, DATABASE_CHAT_DB)
	} else if c.String("name") == "chipmunk" && !c.Bool("all") { 
		
	} else if c.String("name") == "dymo" && !c.Bool("all") { 
		
	} else if c.String("name") == "gatekeeper" && !c.Bool("all") { 
		createDb(names.DATABASE_GATEKEEPER_NAME, names.DATABASE_GATEKEEPER_VALUE, DATABASE_GATEKEEPER_DB)
	} else if c.String("name") == "payment" && !c.Bool("all") { 
		createDb(names.DATABASE_PAYMENT_NAME, names.DATABASE_PAYMENT_VALUES, DATABASE_PAYMENT_DB)
	} else if c.String("name") == "xpress" && !c.Bool("all") { 
		createDb(names.DATABASE_XPRESS_NAME, names.DATABASE_XPRESS_VALUE, DATABASE_XPRESS_DB)
	} else if c.String("xtrack") == "dymo" && !c.Bool("all") { 
		
	} else if c.Bool("all") {
		createDb(names.DATABASE_BRANCH_NAME, names.DATABASE_BRANCH_VALUE, DATABASE_BRANCH_DB)
		createDb(names.DATABASE_CART_NAME, names.DATABASE_CART_VALUE, DATABASE_CART_DB)
		createDb(names.DATABASE_CATALOG_NAME, names.DATABASE_CATALOG_VALUE, DATABASE_CATALOG_DB)
		createDb(names.DATABASE_CHAT_NAME, names.DATABASE_CHAT_VALUE, DATABASE_CHAT_DB)
		createDb(names.DATABASE_GATEKEEPER_NAME, names.DATABASE_GATEKEEPER_VALUE, DATABASE_GATEKEEPER_DB)
		createDb(names.DATABASE_PAYMENT_NAME, names.DATABASE_PAYMENT_VALUES, DATABASE_PAYMENT_DB)
		createDb(names.DATABASE_XPRESS_NAME, names.DATABASE_XPRESS_VALUE, DATABASE_XPRESS_DB)
	}*/
	return nil
}

func delDb(name string) {
	cmd := exec.Command("helm", "delete", name, "-n", names.GUYA_NAMESPACE)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	fmt.Println(err)
}

func deleteDatabase(c *cli.Context) error {
	if c.String("name") == "mongodb" && !c.Bool("all") {
		delDb(names.MONGODB_NAME)
	} else {
		fmt.Println("Command Error")
	}
	/*if c.String("name") == "branch" && !c.Bool("all") {
		delDb(names.DATABASE_BRANCH_NAME)
	} else if c.String("name") == "cart" && !c.Bool("all") { 
		delDb(names.DATABASE_CART_NAME)
	} else if c.String("name") == "catalog" && !c.Bool("all") { 
		delDb(names.DATABASE_CATALOG_NAME)
	} else if c.String("name") == "chat" && !c.Bool("all") { 
		delDb(names.DATABASE_CHAT_NAME)
	} else if c.String("name") == "chipmunk" && !c.Bool("all") { 
		
	} else if c.String("name") == "dymo" && !c.Bool("all") { 
		
	} else if c.String("name") == "gatekeeper" && !c.Bool("all") { 
		delDb(names.DATABASE_GATEKEEPER_NAME)
	} else if c.String("name") == "payment" && !c.Bool("all") { 
		delDb(names.DATABASE_PAYMENT_NAME)
	} else if c.String("name") == "xpress" && !c.Bool("all") { 
		delDb(names.DATABASE_XPRESS_NAME)
	} else if c.String("xtrack") == "dymo" && !c.Bool("all") { 
		
	}*/
	return nil
}


// Helm
func helm(c *cli.Context) error {
	if c.Bool("all") {
		cmd1 := exec.Command("helm", "repo", "add", "bitnami", "https://charts.bitnami.com/bitnami")
		cmd1.Stdout = os.Stdout
		cmd1.Stderr = os.Stderr
		cmd1.Run()
		//helm repo add elastic https://helm.elastic.co
		cmd2 := exec.Command("helm", "repo", "add", "elastic", "https://helm.elastic.co")
		cmd2.Stdout = os.Stdout
		cmd2.Stderr = os.Stderr
		cmd2.Run()
		//exec.Command("helm", "repo", "update").Run()
		fmt.Println("Done","\n")
	}
	return nil
}

// Deployment

func createLocalStorageForElasticsearch() {
	// kubectl apply -f elasticsearch/local-path-storage.yaml
	cmd := exec.Command("kubectl", "apply", "-f", names.ELASTICSEARCH_DEPLOYMENT_LOCAL_STORAGE_VALUE)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	fmt.Println(err)
}

func elasticsearchDeployment() {
	createLocalStorageForElasticsearch()
	// helm install --namespace guya-ltd-elk elasticsearch --version 7.9.1 elastic/elasticsearch --values elasticsearch/values.yaml
	cmd := exec.Command("helm", "install", "-n", names.GUYA_ELK_NAMESPACE, names.ELASTICSEARCH_DEPLOYMENT_NAME, "--version", names.ELASTICSEARCH_VERSION, "elastic/elasticsearch", "--values", names.ELASTICSEARCH_DEPLOYMENT_VALUE)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	fmt.Println(err)
}

func logstashDeployment() {
	// helm install --namespace guya-ltd-elk logstash --version 7.9.1 elastic/logstash --values logstash/values.yaml
	cmd := exec.Command("helm", "install", "-n", names.GUYA_ELK_NAMESPACE, names.LOGSTASH_DEPLOYMENT_NAME, "--version", names.LOGSTASH_DEPLOYMENT_VERSION, "elastic/logstash", "--values", names.LOGSTASH_DEPLOYMENT_VALUE)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	fmt.Println(err)
}

func kibanaDeploymnet() {
	// helm install --namespace guya-ltd-elk kibana --version 7.9.1 elastic/kibana
	cmd := exec.Command("helm", "install", "-n", names.GUYA_ELK_NAMESPACE, names.KIBANA_DEPLOYMENT_NAME, "--version", names.LOGSTASH_DEPLOYMENT_VERSION, "elastic/kibana")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	fmt.Println(err)
}

func deployRabbitmq() {
	// helm install mu-rabbit stable/rabbitmq --namespace guya-ltd
	cmd := exec.Command("helm", "install", "-n", names.GUYA_QUEUE_NAMESPACE, names.RABBITMQ_NAME, "stable/rabbitmq")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	fmt.Println(err)
}

func createNewDeployment(c *cli.Context) error {
	if c.String("name") == "elk" && !c.Bool("all") {
		elasticsearchDeployment()
		logstashDeployment()
		kibanaDeploymnet()
	} else if c.String("name") == "kibana" && !c.Bool("all") {
		kibanaDeploymnet()
	} else if c.String("name") == "elasticsearch" && !c.Bool("all") {
		elasticsearchDeployment()
	} else if c.String("name") == "logstash" && !c.Bool("all") {
		logstashDeployment()
	} else if c.String("name") == "rabbitmq" && !c.Bool("all") {
		deployRabbitmq()
	}
	return nil
}

func delElkDepo(name string) {
	cmd := exec.Command("helm", "delete", name, "-n", names.GUYA_ELK_NAMESPACE)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	fmt.Println(err)
}

func delRabbitMqDepo() {
	cmd := exec.Command("helm", "delete", names.RABBITMQ_NAME, "-n", names.GUYA_QUEUE_NAMESPACE)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	fmt.Println(err)
}

func deleteDeployment(c *cli.Context) error {
	if c.String("name") == "elk" && !c.Bool("all") {
		delElkDepo(names.ELASTICSEARCH_DEPLOYMENT_NAME)
		delElkDepo(names.LOGSTASH_DEPLOYMENT_NAME)
		delElkDepo(names.KIBANA_DEPLOYMENT_NAME)
	} else if c.String("name") == "elasticsearch" && !c.Bool("all") {
		delElkDepo(names.ELASTICSEARCH_DEPLOYMENT_NAME)
	} else if c.String("name") == "logstash" && !c.Bool("all") {
		delElkDepo(names.LOGSTASH_DEPLOYMENT_NAME)
	} else if c.String("name") == "kibana" && !c.Bool("all") {
		delElkDepo(names.KIBANA_DEPLOYMENT_NAME)
	} else if c.String("name") == "rabbitmq" && !c.Bool("all") {
		delRabbitMqDepo()
	} else {
		fmt.Println("Command Error")
	}
	return nil
}