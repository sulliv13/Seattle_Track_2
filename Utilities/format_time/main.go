// Format_time converts time entries in the csv datafiles to
// the preferred format to import into ESRI
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"gopl.io/mine/ais_headings/Seattle_Track_2/Utilities/csvAIS"
)

const dataRoot = `/Users/zac/go/src/gopl.io/mine/ais_headings/Seattle_Track_2/Data`
const timeFieldRE = `BaseDateTime.*`
const existingForm = "2006-01-02T03:04:05"

var jsonFilename = `Field_Descriptions.json`

// csvAIS package includes a global variable Debug that controls whether
// calls to d are executedd.
var d = csvAIS.D

// Flag for --file --output to the command line utility.
// The order of initialization is undefined, so both long and short flags
// must be set up with an init function.
var inFilename string
var outFilename string

func init() {
	const (
		defaultInFile  = "Subset_LA_AIS_Data_Dec_17.csv"
		defaultOutFile = "TimeFormattedOutfile.csv"
		usageFile      = "Filename to transform"
		usageOutput    = "Filename to save transform"
	)
	flag.StringVar(&inFilename, "file", defaultInFile, usageFile)
	flag.StringVar(&inFilename, "f", defaultInFile, usageFile+" (shorthand)")
	flag.StringVar(&outFilename, "out", defaultOutFile, usageOutput)
	flag.StringVar(&outFilename, "o", defaultOutFile, usageOutput+" (shorthand)")
}

func main() {
	// Control the global Debug flag
	csvAIS.Debug = false

	// Implemented flag is [-f]
	flag.Parse()

	// Get the regexp for the the time field header
	reField := regexp.MustCompile(timeFieldRE)

	// Open the input file
	path := filepath.Join(dataRoot, inFilename)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println("Opened file:", path)

	// Open the output file
	path2 := filepath.Join(dataRoot, outFilename)
	f2, err := os.Create(path2)
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	fmt.Println("Created file for output:", path2)

	// Read the first line into memory and break into headers
	// Then read the time for each subsequent line
	var firstLine string
	scanner := bufio.NewScanner(f)
	firstLineRead := false
	var h csvAIS.Headers
	var ti int // index of the time field in the csv headers
	for scanner.Scan() {
		if firstLineRead { // For all lines after the first line
			line := scanner.Text()
			fields := strings.Split(line, ",")
			sampleTime := fields[ti]
			t, err := time.Parse("2006-01-02T15:04:05", sampleTime)
			if err != nil {
				panic(err)
			}
			tString := t.Format("01/02/2006 03:04:05 MST")

			// write the new field
			var transformedLine string
			var lastIndex = len(fields) - 1
			for j, field := range fields {
				if j != ti {
					transformedLine += field
				} else { // it is the header field and needs the transform
					transformedLine += tString
				}
				if j != lastIndex {
					transformedLine += ","
				} else { //only debug if this is the last field
					transformedLine += fmt.Sprintln()
					d(func() {
						fmt.Printf("sampleTime: %s\t\tformatTime:%s\n", sampleTime, tString)
						fmt.Print(transformedLine)
					})
				}
			}
			if !csvAIS.Debug {
				f2.WriteString(transformedLine) // write data to the new file
			}
		} else {
			firstLine = scanner.Text()
			if !csvAIS.Debug {
				f2.WriteString(firstLine) // write the headers
			}
			firstLineRead = true
			h = csvAIS.NewHeaders(firstLine)
			d(func() { fmt.Println(h) })
			for i, header := range h {
				if reField.MatchString(header) {
					ti = i //find the index for the time field
					fmt.Println("Parsing", header)
				}
			}
		}

	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
