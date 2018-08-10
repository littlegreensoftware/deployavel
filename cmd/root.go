package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/littlegreensoftware/deployavel/api"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var (
	cfgFile  string
	yamlFile []byte
	r        api.ForgeRequest
	tokenCnf api.Token
)

var rootCmd = &cobra.Command{
	Use:   "deployavel",
	Short: "Provision Laravel Servers on Digital Ocean",
	Long: `deployavel is a CLI application written in GO to 
easily provision and manage Digital Ocean servers for Laravel based applications.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Create global HTTP Client
	api.NewHTTPClient()

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $PWD/config.yml)")
}

func initConfig() {
	if cfgFile == "" {
		home, err := filepath.Abs("./")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		cfgFile = home + "/config.yml"
	}

	yamlFile, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamlFile, &tokenCnf)
	if err != nil {
		log.Fatal(err)
	}

	r = api.ForgeRequest{
		Token:  tokenCnf.Auth.Value,
		Client: *api.GlobalHTTPClient(),
	}
}
