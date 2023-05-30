package cmd

import (
	_ "github.com/golang-training-examples/example1/cmd/hello"
	"github.com/golang-training-examples/example1/cmd/root"
	_ "github.com/golang-training-examples/example1/cmd/server"
	_ "github.com/golang-training-examples/example1/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
