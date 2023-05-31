package gorm

import (
	"fmt"
	"math/rand"

	"github.com/golang-training-examples/example1/cmd/root"
	"github.com/golang-training-examples/example1/internal/random_utils"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var FlagHost string

var Cmd = &cobra.Command{
	Use:   "gorm",
	Short: "Gorm example",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		example(FlagHost)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagHost,
		"host",
		"H",
		"localhost",
		"Postgres host",
	)
}

type Pet struct {
	gorm.Model
	Name string
	Age  uint
}

func example(host string) {
	db, err := gorm.Open(postgres.Open("host="+host+" user=postgres password=pg dbname=postgres port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Pet{})

	// Create
	db.Create(&Pet{
		Name: random_utils.RandomPetName(),
		Age:  uint(rand.Intn(8) + 1),
	})

	// Read

	var pets []Pet
	db.Find(&pets)
	for _, pet := range pets {
		fmt.Println(pet.Name, pet.Age)
	}
}
