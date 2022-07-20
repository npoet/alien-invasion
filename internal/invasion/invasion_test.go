package invasion

import (
	"alien-invasion/internal/pkg"
	"testing"
)

func TestAnnounceDestruction(t *testing.T) {
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
	resString := "Foo was destroyed by alien and alien1!"
	resString1 := "Foo was destroyed by alien, alien1, and alien2!"
	resString2 := "Foo was destroyed by alien, alien1, alien2, and alien3!"
	if AnnounceDestruction(cityName, aliens) != resString {
		t.Errorf("Incorrect output format, want " + resString)
	}
	if AnnounceDestruction(cityName, aliens1) != resString1 {
		t.Errorf("Incorrect output format, want " + resString1)
	}
	if AnnounceDestruction(cityName, aliens2) != resString2 {
		t.Errorf("Incorrect output format, want " + resString2)
	}
}

func TestDestroyCity(t *testing.T) {}

func TestSimulateInvasion(t *testing.T) {}
