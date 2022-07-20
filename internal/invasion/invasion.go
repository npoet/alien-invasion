// Package invasion implements the Alien creation and interaction methods for simulation

package invasion

import (
	aliens "alien-invasion/internal/aliens"
	utils "alien-invasion/internal/pkg"
	"fmt"
)

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

func destroyCity(worldmap []utils.City) []utils.City {
	// TODO: handle destruction of a city on the worldmap
	//		step 1: check each city for multiple aliens
	//for i := range worldmap {
	//	aliens := worldmap[i].Aliens
	//	if len(worldmap[i].Aliens) >= 2 {
	//
	//	}
	//}
	return []utils.City{}
}

func SimluateInvasion(numAliens int, worldmap map[string]*utils.City) {
	// TODO: import worldmap and gen aliens, add gameplay logic
	//		separate func for gameplay?

	invaders := aliens.GenAliens(numAliens)
	// turn 0: init assignment of aliens to cities randomly based on worldmap input
	for i := range invaders {
		aliens.AssignAlien(&invaders[i], worldmap)
	}
	// aliens move up to 10000 times
	for j := 0; j <= 10000; j++ {
		// turn 1: check overlap on aliens generation, destroy on impact!
		for k := range worldmap {
			if len(worldmap[k].Aliens) >= 2 {
				// TODO: announceDestruction and destroyCity
			}
			for i := range invaders {
				aliens.MoveAlien(&invaders[i], worldmap)
			}
			// TODO: turn n: all cities/aliens destroyed or 10,000 moves for each living aliens
		}
		// TODO: finally, print remaining cities in format of og file
	}
}
