package cmd

import (
	"log"

	"github.com/littlegreensoftware/deployavel/resources"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var serverID int

func init() {
	createSiteCmd.PersistentFlags().IntVarP(&serverID, "id", "i", 0, "Id of Server")
}

var createSiteCmd = &cobra.Command{
	Use:   "site",
	Short: "Create a site for a server on forge",
	Long: `Create will issue a POST request to Forge to create a site:
	
$deployavel create site --id some_id
$deployavel create site -i some_id
`,
	Run: func(cmd *cobra.Command, args []string) {
		var serverCnf resources.SiteRequest

		err := yaml.Unmarshal(yamlFile, &serverCnf)
		if err != nil {
			log.Fatal(err)
		}

		site, err := resources.SiteCreate(r, serverID, serverCnf)
		if err != nil {
			log.Fatal(err)
		}

		data := Must(site.Marshal())

		PrintMust(data)
	},
}
