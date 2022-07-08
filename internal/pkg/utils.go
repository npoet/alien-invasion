package pkg

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

type City struct {
	Name   string
	North  *City
	South  *City
	East   *City
	West   *City
	Aliens []*Alien
}

type Alien struct {
	Name     string
	Location *City
}
