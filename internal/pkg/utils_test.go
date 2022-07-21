package pkg

import "testing"

func TestReverseMap(t *testing.T) {
	// TestReverseMap checks that directions are properly reversed and output structure retained
	initial := "Bar"
	directions := map[string]string{"north": "south", "south": "north", "east": "west", "west": "east"}
	for i := range directions {
		expected := directions[i] + "=" + initial
		if ReverseMap(directions[i], initial) != expected {
			t.Errorf("Incorrect output format, want " + expected)
		}
	}
}
