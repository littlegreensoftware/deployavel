package cmd

import (
	"log"

	"github.com/littlegreensoftware/deployavel/resources"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var createServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Create a single server in Forge",
	Long: `Create will issue a POST request to Forge to create a server:
	
	deployavel get server --id some_id
	deployavel get server -i some_id`,
	Run: func(cmd *cobra.Command, args []string) {
		var serverCnf resources.ServerRequest

		err := yaml.Unmarshal(yamlFile, &serverCnf)
		if err != nil {
			log.Fatal(err)
		}

		server, err := resources.ServerCreate(r, serverCnf.Server)
		if err != nil {
			log.Fatal(err)
		}

		data := Must(server.Marshal())

		PrintMust(data)
	},
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a single resource in Forge",
	Long:  `Create will issue a POST request to Forge to create a resource`,
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createServerCmd)
}
