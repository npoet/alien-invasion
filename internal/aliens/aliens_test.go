package pkg

import (
	"alien-invasion/internal/pkg"
	"strconv"
	"testing"
)

func TestGenAliens(t *testing.T) {
	// TestGenAliens tests that the correct number of aliens are created
	num := []int{1, 3, 5, 11}
	for i := range num {
		res := GenAliens(num[i])
		if len(res) != num[i] {
			t.Errorf("Incorrect number of aliens, want " + strconv.Itoa(num[i]))
		}
	}
}

func TestAssignAlien(t *testing.T) {
	// Test AssignAlien tests that initial alien assignment is working properly
	alien := pkg.Alien{
		Name:     "alien",
		Location: &pkg.City{},
	}
	foo := pkg.City{
		Name:    "Foo",
		Links:   map[*pkg.City]bool{},
		Aliens:  map[string]*pkg.Alien{},
		Initial: "north=Bar",
	}
	bar := pkg.City{
		Name:    "Bar",
		Links:   map[*pkg.City]bool{},
		Aliens:  map[string]*pkg.Alien{},
		Initial: "south=Foo",
	}
	foo.Links[&bar] = true
	bar.Links[&foo] = true
	worldmap := map[string]*pkg.City{
		"Foo": &foo, "Bar": &bar,
	}
	AssignAlien(alien, worldmap)
	num := 0
	for i := range worldmap {
		if len(worldmap[i].Aliens) > 0 {
			num += 1
		}
	}
	if num == 0 {
		t.Errorf("Alien not assigned")
	}
}

func TestMoveAlien(t *testing.T) {
	// TestMoveAliens checks that aliens are properly moved in sim logic
	alien := pkg.Alien{
		Name:     "alien",
		Location: &pkg.City{},
	}
	foo := pkg.City{
		Name:    "Foo",
		Links:   map[*pkg.City]bool{},
		Aliens:  map[string]*pkg.Alien{},
		Initial: "north=Bar",
	}
	bar := pkg.City{
		Name:    "Bar",
		Links:   map[*pkg.City]bool{},
		Aliens:  map[string]*pkg.Alien{},
		Initial: "south=Foo",
	}
	foo.Links[&bar] = true
	bar.Links[&foo] = true
	alien.Location = &bar
	bar.Aliens[alien.Name] = &alien
	worldmap := map[string]*pkg.City{
		"Foo": &foo, "Bar": &bar,
	}
	MoveAlien(alien, worldmap)
	if len(foo.Aliens) == 0 {
		t.Errorf("Alien not moved")
	}
}
