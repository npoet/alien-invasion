package worldmap

import (
	"alien-invasion/internal/pkg"
	"testing"
)

func TestReadRawMap(t *testing.T) {
	// TestReadRawMap checks that text file is properly converted for sim map generation
	filename := "mapsdata/map_test.txt"
	expected := []string{
		"Foo north=Bar west=Baz south=Qu-ux",
		"Bar south=Foo west=Bee",
	}
	output := ReadRawMap(filename)
	for i := range output {
		if output[i] != expected[i] {
			t.Errorf("Incorrect output format: " + output[i] + " want " + expected[i])
		}
	}
}

func TestGenerateMap(t *testing.T) {
	// TestGenerateMap checks that set of city objects is created correctly from raw map conversion
	rawMap := []string{
		"Foo north=Bar west=Baz south=Qu-ux",
		"Bar south=Foo west=Bee",
	}
	foo := pkg.City{
		Name:    "Foo",
		Links:   map[*pkg.City]bool{},
		Aliens:  nil,
		Initial: "north=Bar west=Baz south=Qu-ux",
	}
	bar := pkg.City{
		Name:    "Bar",
		Links:   map[*pkg.City]bool{},
		Aliens:  nil,
		Initial: "south=Foo west=Bee",
	}
	baz := pkg.City{
		Name:    "Baz",
		Links:   map[*pkg.City]bool{},
		Aliens:  nil,
		Initial: "east=Foo",
	}
	bee := pkg.City{
		Name:    "Bee",
		Links:   map[*pkg.City]bool{},
		Aliens:  nil,
		Initial: "east=Bar",
	}
	quux := pkg.City{
		Name:    "Qu-ux",
		Links:   map[*pkg.City]bool{},
		Aliens:  nil,
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
	expected := map[string]*pkg.City{
		"Foo": &foo, "Bar": &bar, "Baz": &baz, "Bee": &bee, "Qu-ux": &quux,
	}
	output := GenerateMap(rawMap)
	for i := range expected {
		if expected[i].Name != output[i].Name {
			t.Errorf("Incorrect output format: " + expected[i].Name)
		}
	}
}
