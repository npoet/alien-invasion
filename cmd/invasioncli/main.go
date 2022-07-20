package main

import (
	"alien-invasion/internal/invasion"
	"alien-invasion/internal/pkg"
	"alien-invasion/internal/worldmap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// get args from cli
	args := os.Args[1:]
	fmt.Println(args[0] + " aliens are invading!")
	// read map from file
	numAliens, err := strconv.Atoi(args[0])
	pkg.Check(err)
	// create string slice from map file and convert to simulation map
	rawMap := worldmap.ReadRawMap("mapsdata/map_test.txt")
	simMap := worldmap.GenerateMap(rawMap)
	// simulate invasion from read map
	invasion.SimluateInvasion(numAliens, simMap)
}
