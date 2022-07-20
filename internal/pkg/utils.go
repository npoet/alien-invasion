package pkg

import "fmt"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

type City struct {
	Name   string
	Links  map[*City]bool
	Aliens []*Alien
}

func (c *City) PrintName() {
	fmt.Println(c.Name)
}

func (c *City) PrintLinks() {
	// var links []string
	for i := range c.Links {
		print(i)
		print(" ")
	}
	fmt.Println()
}

type Alien struct {
	Name     string
	Location *City
}

func (a Alien) Print() {
	fmt.Print(a.Name)
}
