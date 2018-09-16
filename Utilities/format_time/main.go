// Format_time converts time entries in the csv datafiles to
// the preferred format to import into ESRI
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
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
	var ti []int // index of the time fields in the csv headers
	formatTimes := make(map[int]string)
	count := 0
	for scanner.Scan() {
		count++
		if count%500000 == 0 {
			log.Println("Parsing line:", count)
		}
		if !firstLineRead {
			firstLine = scanner.Text()
			if !csvAIS.Debug {
				f2.WriteString(fmt.Sprintln(firstLine)) // write the headers
			}
			h = csvAIS.NewHeaders(firstLine)
			// d(func() { fmt.Println(h) })
			for i, header := range h {
				if reField.MatchString(header) {
					ti = append(ti, i) //find indeces for the time fields
					fmt.Println("Parsing", header)
				}
			}
		}
		if firstLineRead {
			line := scanner.Text()
			fields := strings.Split(line, ",")
			for _, tIndex := range ti {
				sampleTime := fields[tIndex]
				t, err := time.Parse("2006-01-02T15:04:05", sampleTime)
				if err != nil {
					panic(err)
				}
				tString := t.Format("01/02/2006 03:04:05 MST")
				formatTimes[tIndex] = tString
			}

			// write the new fields
			var transformedLine string
			var lastIndex = len(fields) - 1
			for j, field := range fields {
				if !inSlice(j, ti) {
					transformedLine += field
				} else { // it is the time field and needs the transform
					if v, ok := formatTimes[j]; ok {
						transformedLine += v
					} else {
						panic(fmt.Errorf("key: %d not in map formatTimes: %v", j, formatTimes))
					}
				}
				if j != lastIndex {
					transformedLine += ","
				} else {
					transformedLine += fmt.Sprintln()
					d(func() { //only debug if this is the last field
						// fmt.Printf("sampleTime: %s\t\tformatTime:%s\n", sampleTime, tString)
						fmt.Print(transformedLine)
					})
				}
			}
			if !csvAIS.Debug {
				f2.WriteString(transformedLine) // write data to the new file
			}
		}
		if !firstLineRead {
			firstLineRead = true
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func inSlice(j int, indexList []int) bool {
	for _, i := range indexList {
		if j == i {
			return true
		}
	}
	return false
}
