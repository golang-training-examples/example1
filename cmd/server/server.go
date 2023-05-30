package server

import (
	"github.com/golang-training-examples/example1/cmd/root"
	"github.com/golang-training-examples/example1/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var FlagPort int

var Cmd = &cobra.Command{
	Use:   "server",
	Short: "Simple Golang HTTP Server",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		server.Server(FlagPort)
	},
}

func init() {
	viper.BindEnv("PORT")
	port := viper.GetInt("PORT")
	root.Cmd.AddCommand(Cmd)
	Cmd.Flags().IntVarP(
		&FlagPort,
		"port",
		"p",
		port,
		"Port to listen to",
	)
	if port == 0 {
		Cmd.MarkFlagRequired("port")
	}
}
