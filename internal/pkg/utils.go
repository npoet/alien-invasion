package pkg

func Check(e error) {
	if e != nil {
		panic(e)
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
