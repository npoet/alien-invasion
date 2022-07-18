package worldmap

import (
	"alien-invasion/internal/pkg"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ReadRawMap(filename string) []string {
	readFile, err := os.Open(filename)
	pkg.Check(err)
	defer func(readFile *os.File) {
		err := readFile.Close()
		pkg.Check(err)
	}(readFile)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var rawMap []string

	for fileScanner.Scan() {
		rawMap = append(rawMap, fileScanner.Text())
		fmt.Println(fileScanner.Text())
	}
	return rawMap
}

func GenerateMap(rawMap []string) []pkg.City {
	var worldMap []pkg.City
	for i := range rawMap {
		mapLine := strings.Fields(rawMap[i])
		for j := 1; j < len(mapLine); j++ {
			var cityNames []string
			// TODO: create cities from all uniques listed, use opposite directions for cities not
			//		expressed in a mapfile line
			n, _ := regexp.Compile("north=(.*)")
			north := n.FindStringSubmatch(mapLine[j])
			if len(north) >= 1 {
				// TODO: make city obj from each unique in city slice created above
				cityNames = append(cityNames, north[1])
			}
			s, _ := regexp.Compile("south=(.*)")
			south := s.FindStringSubmatch(mapLine[j])
			if len(south) >= 1 {
				// TODO: make city obj from each unique in city slice created above
				cityNames = append(cityNames, south[1])
			}
			e, _ := regexp.Compile("east=(.*)")
			east := e.FindStringSubmatch(mapLine[j])
			if len(east) >= 1 {
				// TODO: make city obj from each unique in city slice created above
				cityNames = append(cityNames, east[1])
			}
			w, _ := regexp.Compile("west=(.*)")
			west := w.FindStringSubmatch(mapLine[j])
			if len(west) >= 1 {
				// TODO: make city obj from each unique in city slice created above
				cityNames = append(cityNames, west[1])
			}
			for i := range cityNames {
				fmt.Print(cityNames[i] + " ")

				newCity := pkg.City{
					Name:   mapLine[0], // all before first space
					North:  nil,
					South:  nil,
					East:   nil,
					West:   nil,
					Aliens: []*pkg.Alien{},
				}
				worldMap = append(worldMap, newCity)
			}
		}
		fmt.Println("")
	}
	return worldMap
}
