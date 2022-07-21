package pkg

import (
	"fmt"
	"strings"
)

// A City represents a place in a mapfile, holds links to other cities, some number of aliens during sim, and initial map directions
type City struct {
	Name    string            // city name, example Foo
	Links   map[*City]bool    // links to Cities created from initial directions
	Aliens  map[string]*Alien // map of Alien objects located in City
	Initial string            // initial directions from input map
}

// An Alien is an invader in the simulation, a number of which are created via command line input
type Alien struct {
	Name     string // alien name created from concatenated lourm ipsum text, example Eumnon
	Location *City  // city where alien is currently located during sim
}

func Check(e error) {
	// Check is a basic error handler removing the need for redundant if statements
	if e != nil {
		panic(e)
	}
}

func PrintMap(worldmap map[string]*City) {
	// PrintMap prints the final map in the format of the input text file, if any cities remain
	if len(worldmap) > 0 {
		fmt.Println("Final map output:")
		keys := map[string]bool{}
		for i := range worldmap {
			keys[worldmap[i].Name] = true
		}
		for j := range worldmap {
			var outputString []string
			outputString = append(outputString, worldmap[j].Name)
			initialFields := strings.Fields(worldmap[j].Initial)
			for k := range initialFields {
				s := strings.Split(initialFields[k], "=")
				if keys[s[1]] {
					outputString = append(outputString, initialFields[k])
				}
			}
			for l := range outputString {
				fmt.Print(outputString[l] + " ")
			}
			fmt.Println()
		}
	} else {
		fmt.Println("All of X has been destroyed.")
	}
}

func ReverseMap(direction string, initial string) string {
	// ReverseMap flips the direction on inferred cities from rawMap
	var returnString string
	split := strings.Split(direction, "=")
	if split[0] == "north" {
		returnString = "south=" + initial
	}
	if split[0] == "south" {
		returnString = "north=" + initial
	}
	if split[0] == "east" {
		returnString = "west=" + initial
	}
	if split[0] == "west" {
		returnString = "east=" + initial
	}
	return returnString
}
