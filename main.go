/*
Example1 is a simple app for Golang Training

Usage:

	example1 hello

For more information, see:

	example1 --help
*/

package main

import (
	"github.com/golang-training-examples/example1/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("EXAMPLE")
	viper.ReadInConfig()
	cmd.Execute()
}
