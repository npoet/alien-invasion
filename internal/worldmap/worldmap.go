package worldmap

import (
	"bufio"
	"fmt"
	"os"
)

type City struct {
	name   string
	north  *City
	south  *City
	east   *City
	west   *City
	aliens []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadRawMap(filename string) []string {
	readFile, err := os.Open(filename)
	check(err)
	defer func(readFile *os.File) {
		err := readFile.Close()
		check(err)
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

func GenerateMap(rawMap []string) [][]City {
	var worldMap [][]City
	for range rawMap {
		// TODO: add city from line, split directions vs regex?
		newCity := City{
			name:   "", // all before first space
			north:  nil,
			south:  nil,
			east:   nil,
			west:   nil,
			aliens: nil,
		}
		fmt.Print(newCity)
		// TODO: insert city into map
		//		account for space needed around center or expand when necessary?
		// worldMap = append(worldMap, newCity)
	}
	return worldMap
}
