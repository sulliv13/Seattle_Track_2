// Creates a more accurate JSON file from the rawDesciptions text file
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const debug = true
const dataRoot = `.`
const headings = `AIS_Proximity_Filtered_Data_Headings.csv`
const rawDescriptions = `rawDescriptions.txt`
const jsonDescriptions = `AIS_Headings_Descriptions.json`

type headers []string
type aisField struct {
	Name        string
	Description string
}

func main() {
	// open the headers csv
	filename := filepath.Join(dataRoot, headings)
	d(func() { fmt.Printf("Opening file: %s\n", filename) })
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var h headers
	h = strings.Split(string(data), ",")
	fmt.Printf("Header file contains: %d columns\n", len(h))
	fmt.Println(h)

	// Open the raw descriptions
	raw, err := os.Open(filepath.Join(dataRoot, rawDescriptions))
	if err != nil {
		panic(err)
	}
	defer raw.Close()
	scanner := bufio.NewScanner(raw)
	var descriptions []string
	for scanner.Scan() {
		descriptions = append(descriptions, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	err = cleanDescriptions(&descriptions)
	if err != nil {
		panic(err)
	}

	// Create a slice of aisField for each record and description
	if len(h) != len(descriptions) {
		panic(fmt.Errorf("len(h): %d != len(descriptions): %d", len(h), len(descriptions)))
	}
	var fields []aisField
	for i, header := range h {
		fields = append(fields, aisField{
			Name:        header,
			Description: descriptions[i],
		})
	}
	d(func() { fmt.Println(fields) })

	// Write the struct to a the headers descriptions json file
	err = writeJSON(fields)
	if err != nil {
		panic(err)
	}

}

func cleanDescriptions(s *[]string) error {
	const r = `".*":"(.*)"`
	re, err := regexp.Compile(r)
	if err != nil {
		return err
	}
	var cleaned []string
	for _, desc := range *s {
		matches := re.FindStringSubmatch(desc)
		if len(matches) == 0 {
			return fmt.Errorf("no regex match in line: %s", desc)
		}
		cleaned = append(cleaned, matches[1])
	}
	*s = cleaned

	return nil // success!
}

func writeJSON(fields []aisField) error {
	b, err := json.MarshalIndent(fields, "", "    ")
	if err != nil {
		return err
	}
	d(func() {
		fmt.Println("Writing JSON Headings")
		os.Stdout.Write(b)
		fmt.Println()
	})

	f, err := os.Create(filepath.Join(dataRoot, jsonDescriptions))
	if err != nil {
		return err
	}
	defer f.Close()
	n, err := f.Write(b)
	if err != nil {
		return err
	}
	fmt.Printf("Wrote %d bytes to: %s", n, filepath.Join(dataRoot, jsonDescriptions))

	return nil // success
}

func d(f func()) {
	if debug {
		f()
	}
}
