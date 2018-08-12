package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a single resource in Forge",
	Long: `Get will issue a GET request to Forge for a single resource:

$deployavel get server -i <server id>
`,
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(serverGetCmd)
}
