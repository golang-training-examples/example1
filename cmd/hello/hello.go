package root

import (
	"github.com/golang-training-examples/example1/cmd/root"
	"github.com/golang-training-examples/example1/pkg/hello"
	"github.com/spf13/cobra"
)

var FlagName string

var Cmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		hello.PrintSayHello(FlagName)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagName,
		"name",
		"n",
		"World",
		"Name to say hello to",
	)
}
