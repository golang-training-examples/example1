package random_utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomPetName() string {
	list := []string{
		"Dela",
		"Nela",
		"Debora",
		"Loki",
		"Kuna",
	}
	return list[rand.Intn(len(list))]
}
