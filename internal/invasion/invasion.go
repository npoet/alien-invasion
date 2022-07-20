// Package invasion implements the Alien creation and interaction methods for simulation

package invasion

import (
	aliens "alien-invasion/internal/aliens"
	utils "alien-invasion/internal/pkg"
	"fmt"
)

func AnnounceDestruction(cityName string, aliens map[string]*utils.Alien) {
	// announceDestruction handles printing fight info to console, including the oxford comma bc important
	var alienNames []string
	for i := range aliens {
		alienNames = append(alienNames, aliens[i].Name)
	}
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

func DestroyCity(city *utils.City, worldmap map[string]*utils.City, invaders map[string]utils.Alien) {
	// destroyCity handles removal of a city from the worldmap
	// remove city from all links
	for i := range worldmap {
		if worldmap[i].Links[city] {
			delete(worldmap[i].Links, city)
		}
	}
	// remove from worldmap
	delete(worldmap, city.Name)
	// remove from invader locations
	for i := range invaders {
		if invaders[i].Location == city {
			delete(invaders, i)
		}
	}
}

func SimulateInvasion(numAliens int, worldmap map[string]*utils.City) map[string]*utils.City {
	// SimulateInvasion handles creation of aliens, initial city assignment, and sim loop (up to 10k moves per alien)
	invaders := aliens.GenAliens(numAliens)
	// turn 0: init assignment of aliens to cities randomly based on worldmap input
	for i := range invaders {
		aliens.AssignAlien(invaders[i], worldmap)
	}
	// aliens move up to 10000 times
	for j := 0; j <= 10000; j++ {
		// check overlap for cities in map, destroy on impact!
		for k := range worldmap {
			if len(worldmap[k].Aliens) >= 2 {
				// if greater than 2 aliens in a city: announceDestruction and destroyCity
				AnnounceDestruction(worldmap[k].Name, worldmap[k].Aliens)
				DestroyCity(worldmap[k], worldmap, invaders)
			}
		}
		// attempt to move each alien once per turn
		// check if worldmap has been depleted:
		if len(worldmap) == 0 {
			continue
		} else {
			for i := range invaders {
				aliens.MoveAlien(invaders[i], worldmap)
			}
		}
	}
	return worldmap
}
