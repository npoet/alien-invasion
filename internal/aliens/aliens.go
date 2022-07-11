package pkg

import (
	"alien-invasion/internal/pkg"
	"github.com/brianvoe/gofakeit/v6"
	"strings"
)

func GenAliens(numAliens int) []pkg.Alien {
	// genAliens creates numAliens Aliens with a random city choice from the names given in cities []string
	// The genAliens func should be used only to create the initial 'board' state once a map is imported
	var newAliens []pkg.Alien
	for i := 0; i < numAliens; i++ {
		// somehow the best way to create some alien-sounding names
		names := strings.Fields(gofakeit.LoremIpsumSentence(3))
		newAlien := pkg.Alien{
			Name:     names[0] + names[1],
			Location: &pkg.City{},
		}
		newAliens = append(newAliens, newAlien)
	}
	return newAliens
}

func MoveAlien(alien *pkg.Alien, city *pkg.City) {
	// TODO: handle moving aliens and updating aliens/city objects
}
