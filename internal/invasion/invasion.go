package invasion

import (
	"alien-invasion/internal/worldmap"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
)

type Alien struct {
	name     string
	location worldmap.City
}

func genAliens(numAliens int) []Alien {
	var newAliens []Alien
	for i := 0; i < numAliens; i++ {
		newAlien := Alien{
			name:     gofakeit.HackerPhrase(),
			location: worldmap.City{},
		}
		newAliens = append(newAliens, newAlien)
		fmt.Println(newAlien)
	}
	return newAliens
}

func simluateInvasion() {
	// TODO: import worldmap and gen aliens, add gameplay logic
	//		separate func for gameplay?
}
