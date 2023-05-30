package root

import (
	"fmt"

	"github.com/golang-training-examples/example1/cmd/root"
	"github.com/spf13/cobra"
)

var FlagName string

var Cmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("Hello %s!\n", FlagName)
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
