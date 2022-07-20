package pkg

import "testing"

func TestPrintMap(t *testing.T) {}

func TestReverseMap(t *testing.T) {
	initial := "Bar"
	directions := map[string]string{"north": "south", "south": "north", "east": "west", "west": "east"}
	for i := range directions {
		expected := directions[i] + "=" + initial
		if ReverseMap(directions[i], initial) != expected {
			t.Errorf("Incorrect output format, want " + expected)
		}
	}
}
