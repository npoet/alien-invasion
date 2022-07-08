package main

import (
	"alien-invasion/internal/worldmap"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	numAliens := args[0]
	fmt.Println(numAliens + " aliens are invading!")
	rawMap := worldmap.ReadRawMap("mapsdata/map_test.txt")
	gameMap := worldmap.GenerateMap(rawMap)
	fmt.Print(gameMap)
}
