package version

import (
	"fmt"

	"github.com/golang-training-examples/example1/cmd/root"
	"github.com/golang-training-examples/example1/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "version",
	Short:   "Print the version number of example1",
	Aliases: []string{"v"},
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
