package main

import (
	"alien-invasion/internal/invasion"
	"alien-invasion/internal/pkg"
	"alien-invasion/internal/worldmap"
	"flag"
	"fmt"
	"strconv"
)

func main() {
	// define flags from cli for optional testing file, numAliens
	// assumed that file is accessible in current directory (/alien-invasion)
	filename := flag.String("infile", "internal/worldmap/mapsdata/map_test.txt", "filename for map input")
	// assumed that aliens will be given as a positive integer
	aliens := flag.String("aliens", "20", "number of aliens for simulation")
	// parse flags
	flag.Parse()
	// announce invasion
	fmt.Println(*aliens + " aliens are invading X!")
	// read map from file
	numAliens, err := strconv.Atoi(*aliens)
	pkg.Check(err)
	// create string slice from map file and convert to simulation map
	rawMap := worldmap.ReadRawMap(*filename)
	simMap := worldmap.GenerateMap(rawMap)
	// simulate invasion from read map
	finalMap := invasion.SimulateInvasion(numAliens, simMap)
	// print out remaining world state in original format 'Foo north=Bar west=Baz south=Qu-ux'
	pkg.PrintMap(finalMap)
}
