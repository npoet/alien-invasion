package worldmap

import (
	"alien-invasion/internal/pkg"
	"bufio"
	"fmt"
	"os"
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

func GenerateMap(rawMap []string) [][]pkg.City {
	var worldMap [][]pkg.City
	for range rawMap {
		// TODO: add city from line, split directions vs regex?
		newCity := pkg.City{
			Name:   "", // all before first space
			North:  nil,
			South:  nil,
			East:   nil,
			West:   nil,
			Aliens: nil,
		}
		fmt.Print(newCity)
		// TODO: insert city into map
		//		account for space needed around center or expand when necessary?
		// worldMap = append(worldMap, newCity)
	}
	return worldMap
}
