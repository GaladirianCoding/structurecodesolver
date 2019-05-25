# structurecodesolver

structurecodesolver is a lightweight tool for solving artifact puzzle in [Structure Idle](http://structure.zefiris.su/)

## Installation

structurecodesolver requires go to build the binary.
Once you have cloned the repo, and cd'ed into the directory simply run
```
$ go build
```

which should build the binary inside the same directory.

## Usage

To run the solver simply navigate to the folder with the binary and run it with the puzzle passed in as a comma-deliniated list of two character entries as the only argument.

For example if you wanted to solve a puzzle with the runes ["HT", IN", FI", "TI", "IG", "NG", "GH"] you would run
```
$ ./structurecodesolver HT,IN,FI,TI,IG,NG,GH
```

which should print "FIGHTING" or "FINGHTIG" (if the puzzle has multiple solutions it will return a random one)
If there are no solutions to the puzzle it will return "No solution"

### Todos

 - Write some Tests

License
----

MIT
