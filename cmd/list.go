package cmd

import (
	"github.com/spf13/cobra"
)

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
