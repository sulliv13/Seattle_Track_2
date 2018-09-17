// Headings opens a CSV datafile and pretty prints the data fields
// in that file along with their descriptions to stdout
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
	"unicode"
)

const debug = false // logs debug statements to stdout
const dataRoot = `/Users/zac/go/src/gopl.io/mine/ais_headings/Seattle_Track_2/Data`
const pad = ' ' //padding character for prety print

var jsonFilename = `AIS_field_descriptions.json`

type csvField struct {
	Name        string
	Description string
}

// Headers are the column names of a csv data file
type Headers []string

// Flags for filename to check headings.
// The order of initialization is undefined, so they must be set up
// with an init function.
var filename string

func init() {
	const (
		defaultFile = "AIS_LA_SD_Jan_1_to_15_2016_Filtered_by_Proximity.csv"
		usage       = "Filename to check for headers"
	)
	flag.StringVar(&filename, "file", defaultFile, usage)
	flag.StringVar(&filename, "f", defaultFile, usage+" (shorthand)")
}

func main() {
	//Parse for flags
	flag.Parse()

	// Open the test file
	path := filepath.Join(dataRoot, filename)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println("Opened file:", path)

	// Read the first line into memory and break into headers
	var firstLine string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		firstLine = scanner.Text()
		//Exit if the first row begins with a number
		r := rune(firstLine[0])
		if unicode.IsDigit(r) {
			fmt.Println("The file begins with a number and does not appear to have headers")
			os.Exit(0)
		}
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	h := NewHeaders(firstLine)
	d(func() { fmt.Println(h) })

	// Load the JSON field descriptions and create a map from them
	descriptionsBlob, err := ioutil.ReadFile(filepath.Join(dataRoot, jsonFilename))
	if err != nil {
		panic(err)
	}
	var fields []csvField
	err = json.Unmarshal(descriptionsBlob, &fields)
	if err != nil {
		panic(err)
	}

	descMap := make(map[string]string)
	for _, af := range fields {
		descMap[strings.TrimSpace(af.Name)] = strings.TrimSpace(af.Description)
	}
	d(func() { fmt.Println("Description map has length:", len(descMap)) })
	d(func() {
		for k, v := range descMap {
			fmt.Printf("\t%s len(k)=%d:\t%v\n", k, len(k), v)
		}
	})

	// For each header pretty print its name and description
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, pad, 0)
	for _, header := range h {
		header = strings.TrimSpace(header)
		if d, ok := descMap[header]; ok {
			fmt.Fprintf(w, "%s: \t%s\n", header, d)
		} else {
			log.Println("no match for key: ", header)
			log.Printf("missing key has len(k)=%d\n", len(header))
		}
	}
	w.Flush()
}

// NewHeaders returns a set of headers from the filename.
func NewHeaders(data string) Headers {
	var h = new(Headers)

	tmp := strings.Split(string(data), ",")
	for _, s := range tmp {
		myLen := 0
		var tmp []rune
		for _, r := range s {
			if unicode.IsLetter(r) || r == '_' || unicode.IsDigit(r) {
				tmp = append(tmp, r)
				myLen++
			}
		}
		s = string(tmp)
		// s = strings.TrimSpace(s)
		d(func() { fmt.Printf("NewHeaders: len of header %s: %d\n", s, len(s)) })
		*h = append(*h, s)
	}

	return *h
}

// String satisfies the fmt.Stringer interface for Headers
func (h Headers) String() string {
	b := new(bytes.Buffer)
	for i, field := range h {
		fmt.Fprintf(b, "\t%d:\t%s\n", i+1, field)
	}
	return b.String()
}

// d provides a switch for debug printing
func d(f func()) {
	if debug {
		f()
	}
}
