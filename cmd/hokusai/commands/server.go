package commands

import (
	"log"

	"github.com/netsensei/hokusai/http/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultHost = "localhost"
	defaultPort = 3000
)

func init() {
	serverStartCmd.Flags().String("host", defaultHost, "server host")
	viper.BindPFlag("host", serverStartCmd.Flags().Lookup("host"))
	viper.SetDefault("host", defaultHost)
	serverStartCmd.Flags().Int("port", defaultPort, "server port")
	viper.BindPFlag("port", serverStartCmd.Flags().Lookup("port"))
	viper.SetDefault("port", defaultPort)

	serverCmd.AddCommand(serverStartCmd)

	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server [command]",
	Short: "The Hokusai HTTP server",
}

var serverStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start the http server",
	Run: func(cmd *cobra.Command, args []string) {
		err := server.Start(server.Options{
			Host: viper.GetString("host"),
			Port: viper.GetInt("port"),
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}
