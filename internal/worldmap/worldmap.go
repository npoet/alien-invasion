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
			fmt.Println(mapLine[j])
		}
		north, _ := regexp.Compile("north=*s")
		south, _ := regexp.Compile("south=*s")
		east, _ := regexp.Compile("east=*s")
		west, _ := regexp.Compile("west=*s")
		// TODO: create cities from all uniques listed, use opposite directions for cities not
		//		expressed in a mapfile line

		newCity := pkg.City{
			Name:   mapLine[0], // all before first space
			North:  nil,
			South:  nil,
			East:   nil,
			West:   nil,
			Aliens: []*pkg.Alien{},
		}
		fmt.Println(newCity)
		worldMap = append(worldMap, newCity)
	}
	return worldMap
}
