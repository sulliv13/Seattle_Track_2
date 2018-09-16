// Headings opens a CSV datafile and pretty prints the data fields
// in that file along with their descriptions to stdout
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const dataRoot = `../..`

var csvFilename = `LA_AIS_Data_Dec_17.csv`

func main() {
	path := filepath.Join(dataRoot, csvFilename)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println("Opened file:", path)

}
