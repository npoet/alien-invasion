package pkg

import (
	"fmt"
	"strings"
)

type City struct {
	Name    string
	Links   map[*City]bool
	Aliens  map[string]*Alien
	Initial string
}

type Alien struct {
	Name     string
	Location *City
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
