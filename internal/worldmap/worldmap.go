package worldmap

import (
	"alien-invasion/internal/pkg"
	"bufio"
	"os"
	"regexp"
	"strings"
)

func ReadRawMap(filename string) []string {
	// ReadRawMap takes a filename input and reads map from text, returning slice of strings for sim
	readFile, err := os.Open(filename)
	pkg.Check(err)
	// defer file closure, handle errors
	defer func(readFile *os.File) {
		err := readFile.Close()
		pkg.Check(err)
	}(readFile)
	// scan file by line
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var rawMap []string
	// convert lines to string and add to rawMap
	for fileScanner.Scan() {
		rawMap = append(rawMap, fileScanner.Text())
	}
	return rawMap
}

func GenerateMap(rawMap []string) map[string]*pkg.City {
	// GenerateMap takes an input of string slice from map file and returns map of strings to City objects for sim
	worldMap := map[string]*pkg.City{}
	// 'set' to track existing cities
	cityNames := map[string]bool{}
	// iterate rawMap and create cities, add links:
	for i := range rawMap {
		// split each line of rawMap by whitespace
		mapLine := strings.Fields(rawMap[i])
		// for each entry in mapLine:
		for j := 1; j < len(mapLine); j++ {
			// add cityname from line start if unique
			isCreated := cityNames[mapLine[0]]
			if !isCreated {
				// create city from line start name if unique
				initCity := pkg.City{
					Name:   mapLine[0],
					Links:  map[*pkg.City]bool{},
					Aliens: map[string]*pkg.Alien{},
				}
				// add name to 'set' and worldmap after city creation
				cityNames[mapLine[0]] = true
				worldMap[mapLine[0]] = &initCity
			}
			// compile regex to find remaining cities in line from directions
			r, _ := regexp.Compile("=(.*)")
			// match after equals sign, the directions themselves don't necessarily matter
			match := r.FindStringSubmatch(mapLine[j])
			for k := 1; k < len(match); k++ {
				// create cities from uniques, skip first entry
				isCreated := cityNames[match[k]]
				if !isCreated {
					// create new
					newCity := pkg.City{
						Name:   match[k],
						Links:  map[*pkg.City]bool{worldMap[mapLine[0]]: true},
						Aliens: map[string]*pkg.Alien{},
					}
					// add to worldmap, names
					worldMap[newCity.Name] = &newCity
					cityNames[match[k]] = true
					// add link to initial city in line
					worldMap[mapLine[0]].Links[&newCity] = true
				} else {
					// add link to initial city in line
					worldMap[mapLine[0]].Links[worldMap[match[k]]] = true
				}
			}
		}
	}
	return worldMap
}
