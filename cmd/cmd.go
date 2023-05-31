package cmd

import (
	_ "github.com/golang-training-examples/example1/cmd/hello"
	_ "github.com/golang-training-examples/example1/cmd/random_pet_name"
	"github.com/golang-training-examples/example1/cmd/root"
	_ "github.com/golang-training-examples/example1/cmd/server"
	_ "github.com/golang-training-examples/example1/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
