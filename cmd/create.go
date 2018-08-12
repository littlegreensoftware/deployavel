package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a single resource in Forge",
	Long:  `Create will issue a POST request to Forge to create a resource`,
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createServerCmd)
	createCmd.AddCommand(createOpCacheCmd)
}
