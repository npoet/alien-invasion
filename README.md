# alien-invasion
Mad aliens are about to invade the earth, and you are tasked with simulating the invasion.

World map inputs are stored in alien-invasion/mapsdata, assumed to be of format [CITY direction=CITY ...]

### Build

* From main directory run `go build cmd/invasioncli/main.go`

### Run

* Run main.go with flag 'aliens' for number of aliens in simulation: `./main.go -aliens=20`
* Can add optional flag 'infile' for other test maps: `./main.go -aliens=5 -infile=map.txt`
  * default map location is /mapsdata/map_test.txt
* `./main.go` with no options given will default to a 20 alien simulation on the default test map

### Test

* Includes basic tests for /internal, can run `go test` for:
  * /aliens
  * /invasion
  * /pkg
  * /worldmap