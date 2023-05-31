package random_pet_name

import (
	"fmt"

	"github.com/golang-training-examples/example1/cmd/root"
	"github.com/golang-training-examples/example1/internal/random_utils"
	"github.com/spf13/cobra"
)

var FlagName string

var Cmd = &cobra.Command{
	Use:     "random-pet-name",
	Aliases: []string{"rpn"},
	Short:   "Get a random pet name",
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(random_utils.RandomPetName())
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
