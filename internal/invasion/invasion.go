// Package invasion implements the Alien creation and interaction methods for simulation

package invasion

import (
	"alien-invasion/internal/pkg"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
)

func genAliens(numAliens int) []pkg.Alien {
	// genAliens creates numAliens Aliens with a random city choice from the names given in cities []string
	// The genAliens func should be used only to create the initial 'board' state once a map is imported
	var newAliens []pkg.Alien
	for i := 0; i < numAliens; i++ {
		newAlien := pkg.Alien{
			Name:     gofakeit.NounCollectiveThing() + " " + gofakeit.BeerName(),
			Location: &pkg.City{},
		}
		newAliens = append(newAliens, newAlien)
		fmt.Println(newAlien)
	}
	return newAliens
}

func announceDestruction(cityName string, alienNames []string) {
	// announceDestruction handles printing fight info to console, including the oxford comma bc important
	var aliensList string
	if len(alienNames) == 2 {
		aliensList = alienNames[0] + " and " + alienNames[1]
	} else if len(alienNames) >= 3 {
		for i := 0; i < len(alienNames)-1; i++ {
			aliensList = aliensList + alienNames[i] + ", "
		}
		aliensList = aliensList + "and " + alienNames[len(alienNames)-1]
	} else {
		panic("Not enough alienNames to fight...")
	}
	fmt.Println(cityName + " has been destroyed by " + aliensList + "!")
}

func cityDestruction() {
	// TODO: handle destruction of a city on the worldmap
}

func SimluateInvasion(numAliens int, worldmap [][]pkg.City) {
	// TODO: import worldmap and gen aliens, add gameplay logic
	//		separate func for gameplay?

	aliens := genAliens(numAliens)
	for range aliens {
		// TODO: assign aliens to cities randomly based on worldmap input
		// TODO: turn 1: check overlap on alien generation, destroy on impact!
	}
	// TODO: turn 2: aliens move, check and destroy (announceDestruction())
	// TODO: turn n: all cities/aliens destroyed or 10,000 moves for each living alien
	// TODO: finally, print remaining cities in format of og file
	fmt.Print(aliens)
}
