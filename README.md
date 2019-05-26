# structurecodesolver

structurecodesolver is a lightweight tool for solving artifact puzzle in [Structure Idle](http://structure.zefiris.su/)

## Installation

structurecodesolver requires Go to build the binary.
There are two ways to get the source code onto your machine. One is to simply clone the repo and navigate into the directory. The other way is to run the following go command.

```shell
go get -d github.com/galadiriancoding/structurecodesolver
```

Which should install the repo to $GOPATH/src/galadiriancoding/structurecodesolver

Once you have aquired the source, and navigated into the directory simply run

```shell
go build
```

which should build the binary inside the same directory.

If you want to use the built binary from anywhere on your machine you can run the following go command from inside the main directory

```shell
go install
```

Or when fetching the source code runn the following command instead 

```shell
go get github.com/galadiriancoding/structurecodesolver
```

If you don't want to deal with Go then you can just download the binary from the release page.

## Usage

To run the solver simply navigate to the folder with the binary and run it with the puzzle passed in as a comma-deliniated list of two character entries as the only argument.

For example if you wanted to solve a puzzle with the runes ["HT", IN", FI", "TI", "IG", "NG", "GH"] you would run

```shell
$ ./structurecodesolver HT,IN,FI,TI,IG,NG,GH
FIGHTING
```

which should print "FIGHTING" or "FINGHTIG" (if the puzzle has multiple solutions it will return a random one)
If there are no solutions to the puzzle it will return "No solution"

### Todos

- Write some Tests

## License

MIT
