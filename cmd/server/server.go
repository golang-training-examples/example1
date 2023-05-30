package server

import (
	"github.com/golang-training-examples/example1/cmd/root"
	"github.com/golang-training-examples/example1/pkg/server"
	"github.com/spf13/cobra"
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
	root.Cmd.AddCommand(Cmd)
	Cmd.Flags().IntVarP(
		&FlagPort,
		"port",
		"p",
		0,
		"Port to listen to",
	)
	Cmd.MarkFlagRequired("port")
}
