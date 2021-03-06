package cmd

import (
	"fmt"
	"log"

	"github.com/littlegreensoftware/deployavel/resources"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var id int

func init() {
	serverGetCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "Id of Server")
	createOpCacheCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "Id of Server")
	deleteOpCacheCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "Id of Server")
}

var createOpCacheCmd = &cobra.Command{
	Use:   "opcache",
	Short: "Enable opcache for a server on forge",
	Long: `Create will issue a POST request to Forge to enable opcache:
	
$deployavel create opcache --id some_id
$deployavel create opcache -i some_id
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := resources.EnableOpCache(r, id); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Successfully enabled OPCache")
	},
}

var deleteOpCacheCmd = &cobra.Command{
	Use:   "opcache",
	Short: "Disable opcache for a server on forge",
	Long: `Delete will issue a DELETE request to Forge to disable opcache:
	
$deployavel delete opcache --id some_id
$deployavel delete opcache -i some_id
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := resources.DisableOpCache(r, id); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Successfully disabled OPCache")
	},
}

var createServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Create a single server in Forge",
	Long: `Create will issue a POST request to Forge to create a server:
	
$deployavel create server --id some_id
$deployavel create server -i some_id
`,
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

var serverGetCmd = &cobra.Command{
	Use:   "server",
	Short: "Get a single server from Forge",
	Long: `Get will issue a GET request to Forge for a single server:
	
$deployavel get server --id some_id
$deployavel get server -i some_id
`,
	Run: func(cmd *cobra.Command, args []string) {
		server, err := resources.ServerRead(r, id)
		if err != nil {
			log.Fatal(err)
		}

		data := Must(server.Marshal())

		PrintMust(data)
	},
}

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
