package root

import (
	"github.com/golang-training-examples/example1/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Cmd = &cobra.Command{
	Use:   "example1",
	Short: "example1, " + version.Version,
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("EXAMPLE")
	viper.ReadInConfig()
}
