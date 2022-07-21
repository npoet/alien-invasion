# alien-invasion
Mad aliens are about to invade the earth, and you are tasked with simulating the invasion.

A command line interface is provided to simulate the alien invasion of the world of X. The program outputs details of 
the simulation and the final map state after it completes.
* World map inputs are stored in alien-invasion/internal/worldmap/mapsdata, assumed to be of format 
[CITY direction=CITY ...]
* Names created with https://github.com/brianvoe/gofakeit

### Build

* From main directory run `go build cmd/invasioncli/main.go`

### Run

* Run main with flag 'aliens' for number of aliens in simulation: `./main -aliens=20`
* Can add optional flag 'infile' for other test maps: `./main -aliens=5 -infile=map.txt`
  * default map location is /mapsdata/map_test.txt
* `./main` with no options given will default to a 20 alien simulation on the default test map

### Test

* Includes basic tests for /internal, can run `go test` for:
  * /aliens
  * /invasion
  * /pkg
  * /worldmap

### Sample Output:

* `./main -aliens=10`

```
10 aliens are invading!
Qu-ux has been destroyed by Quamvoluptas and Consequaturet!
Foo has been destroyed by Minimaoptio, Nihilvelit, and Eumnon!
Baz has been destroyed by Quiillo, Etet, and Porrosed!
Final map output:
Bar west=Bee 
Bee east=Bar 
```