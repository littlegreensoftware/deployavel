package cmd

import (
	"log"

	"github.com/littlegreensoftware/deployavel/resources"
	"github.com/spf13/cobra"
)

var serverListCmd = &cobra.Command{
	Use:   "server",
	Short: "List all servers in Forge",
	Long: `List will issue a GET request to Forge for all servers:
	
$deployavel list server
	`,
	Run: func(cmd *cobra.Command, args []string) {
		servers, err := resources.ServerList(r)
		if err != nil {
			log.Fatal(err)
		}

		data := Must(servers.Marshal())

		PrintMust(data)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all resources in Forge",
	Long: `List will issue a GET request to Forge for all resources of a specific type:
	
$deployavel list server
	`,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(serverListCmd)
}
