package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a single resource in Forge",
	Long: `Delete will issue a DELETE request to Forge to delete a resource:

$deployavel delete opcache -i <server id>
`,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteOpCacheCmd)
}
