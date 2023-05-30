package root

import (
	"github.com/golang-training-examples/example1/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "example1",
	Short: "example1, " + version.Version,
}
