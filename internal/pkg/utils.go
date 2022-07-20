package pkg

import "fmt"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func PrintMap(worldmap map[string]*City, rawMap []string) {
	for i := range worldmap {
		name := worldmap[i].Name
		// TODO: get initial line from rawMap if main, check links
		var links []string
		for j := range worldmap[i].Links {
			links = append(links, j.Name)
		}
		fmt.Print(name + " ")
		for k := range links {
			// TODO: get initial direction from rawMap
			fmt.Print(links[k] + " ")
		}
	}
}

type City struct {
	Name   string
	Links  map[*City]bool
	Aliens map[string]*Alien
}

type Alien struct {
	Name     string
	Location *City
}
