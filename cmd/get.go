package cmd

import (
	"log"

	"bitbucket.org/littlegreensoftware/deployavel/resources"
	"github.com/spf13/cobra"
)

var id int

var serverGetCmd = &cobra.Command{
	Use:   "server",
	Short: "Get a single server from Forge",
	Long: `Get will issue a GET request to Forge for a single server:
	
	deployavel get server --id some_id
	deployavel get server -i some_id`,
	Run: func(cmd *cobra.Command, args []string) {
		server, err := resources.ServerRead(r, id)
		if err != nil {
			log.Fatal(err)
		}

		data := Must(server.Marshal())

		PrintMust(data)
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a single resource in Forge",
	Long:  `Get will issue a GET request to Forge for a single resource`,
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(serverGetCmd)

	getCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "Id of Server")
}
