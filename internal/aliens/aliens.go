package pkg

import (
	"alien-invasion/internal/pkg"
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
	"strings"
)

func GenAliens(numAliens int) map[string]pkg.Alien {
	// genAliens creates and returns a map of name strings to numAliens Aliens
	newAliens := map[string]pkg.Alien{}
	for i := 0; i < numAliens; i++ {
		// somehow the best way to create some alien-sounding names
		names := strings.Fields(gofakeit.LoremIpsumSentence(3))
		newAlien := pkg.Alien{
			Name:     names[0] + names[1],
			Location: &pkg.City{},
		}
		newAliens[newAlien.Name] = newAlien
	}
	return newAliens
}

func AssignAlien(alien pkg.Alien, worldmap map[string]*pkg.City) {
	// AssignAlien handles assigning a given alien to a city and updating aliens/city objects
	// get random city from map, select random seed and get keys (arbitrary order return + random seed)
	i := rand.Intn(len(worldmap))
	keys := make([]string, 0, len(worldmap))
	for k := range worldmap {
		keys = append(keys, k)
	}
	// add alien to random city
	worldmap[keys[i]].Aliens[alien.Name] = &alien
	// add city to alien
	alien.Location = worldmap[keys[i]]
}

func MoveAlien(alien pkg.Alien, worldmap map[string]*pkg.City) {
	// MoveAlien handles moving a given alien from their current city to a random one from its links
	// get random city from Links, select random seed and get keys (arbitrary order return + random seed)
	// only move alien if links are available
	if len(alien.Location.Links) > 0 {
		i := rand.Intn(len(alien.Location.Links))
		keys := make([]pkg.City, 0, len(alien.Location.Links))
		for k := range alien.Location.Links {
			keys = append(keys, *k)
		}
		// add alien to random city from links
		worldmap[keys[i].Name].Aliens[alien.Name] = &alien
		// remove alien from current city
		delete(alien.Location.Aliens, alien.Name)
		// set location on alien
		alien.Location = worldmap[keys[i].Name]
	}
}
