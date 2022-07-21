package invasion

import (
	"alien-invasion/internal/pkg"
	"testing"
)

func TestAnnounceDestruction(t *testing.T) {
	// TestAnnounceDestruction tests the sentence creation with commas etc
	cityName := "Foo"
	aliens := map[string]*pkg.Alien{
		"alien":  {Name: "alien", Location: nil},
		"alien1": {Name: "alien1", Location: nil},
	}
	aliens1 := map[string]*pkg.Alien{
		"alien":  {Name: "alien", Location: nil},
		"alien1": {Name: "alien1", Location: nil},
		"alien2": {Name: "alien2", Location: nil},
	}
	aliens2 := map[string]*pkg.Alien{
		"alien":  {Name: "alien", Location: nil},
		"alien1": {Name: "alien1", Location: nil},
		"alien2": {Name: "alien2", Location: nil},
		"alien3": {Name: "alien3", Location: nil},
	}
	resString := "Foo has been destroyed by alien and alien1!"
	resString1 := "Foo has been destroyed by alien, alien1, and alien2!"
	resString2 := "Foo has been destroyed by alien, alien1, alien2, and alien3!"
	if len(AnnounceDestruction(cityName, aliens)) != len(resString) {
		t.Errorf("Incorrect output format, want " + resString)
	}
	if len(AnnounceDestruction(cityName, aliens1)) != len(resString1) {
		t.Errorf("Incorrect output format, want " + resString1)
	}
	if len(AnnounceDestruction(cityName, aliens2)) != len(resString2) {
		t.Errorf("Incorrect output format, want " + resString2)
	}
}

func TestDestroyCity(t *testing.T) {
	// TestDestroyCity checks that a city and any aliens are properly removed
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
	invaders := map[string]pkg.Alien{
		"alien": alien,
	}
	DestroyCity(&bar, worldmap, invaders)
	if foo.Links[&bar] {
		t.Errorf("City not destroyed")
	}
}

func TestSimulateInvasion(t *testing.T) {
	// TestSimulateInvasion checks that invasion logic runs properly
	numAliens := 10
	foo := pkg.City{
		Name:    "Foo",
		Links:   map[*pkg.City]bool{},
		Aliens:  map[string]*pkg.Alien{},
		Initial: "north=Bar west=Baz south=Qu-ux",
	}
	bar := pkg.City{
		Name:    "Bar",
		Links:   map[*pkg.City]bool{},
		Aliens:  map[string]*pkg.Alien{},
		Initial: "south=Foo west=Bee",
	}
	baz := pkg.City{
		Name:    "Baz",
		Links:   map[*pkg.City]bool{},
		Aliens:  map[string]*pkg.Alien{},
		Initial: "east=Foo",
	}
	bee := pkg.City{
		Name:    "Bee",
		Links:   map[*pkg.City]bool{},
		Aliens:  map[string]*pkg.Alien{},
		Initial: "east=Bar",
	}
	quux := pkg.City{
		Name:    "Qu-ux",
		Links:   map[*pkg.City]bool{},
		Aliens:  map[string]*pkg.Alien{},
		Initial: "north=Foo",
	}
	foo.Links[&bar] = true
	foo.Links[&baz] = true
	foo.Links[&quux] = true
	bar.Links[&foo] = true
	bar.Links[&bee] = true
	baz.Links[&foo] = true
	quux.Links[&foo] = true
	bee.Links[&bar] = true
	worldmap := map[string]*pkg.City{
		"Foo": &foo, "Bar": &bar, "Baz": &baz, "Bee": &bee, "Qu-ux": &quux,
	}
	output := SimulateInvasion(numAliens, worldmap)
	if output == nil {
		t.Errorf("Incorrect output format, expected map")
	}
}
