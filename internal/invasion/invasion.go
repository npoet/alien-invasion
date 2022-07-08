package invasion

import (
	"alien-invasion/internal/pkg"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
)

func genAliens(numAliens int) []pkg.Alien {
	var newAliens []pkg.Alien
	for i := 0; i < numAliens; i++ {
		newAlien := pkg.Alien{
			Name:     gofakeit.HackerPhrase(),
			Location: &pkg.City{},
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
