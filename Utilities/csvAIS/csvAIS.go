// Package csvAIS is a package to manipulate csv file of AIS data for
// HACKtheMACHINE Seattle.
package csvAIS

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

// CsvField holds data element descriptions for the columns in csv files.
type CsvField struct {
	Name        string
	Description string
}

// Debug is a global debug switch used in the D function
var Debug = false

// Headers are the column names of a csv data file
type Headers []string

// NewHeaders returns a set of headers from the filename.
func NewHeaders(data string) Headers {
	var h = new(Headers)

	tmp := strings.Split(string(data), ",")
	for _, s := range tmp {
		myLen := 0
		var tmp []rune
		for _, r := range s {
			if unicode.IsLetter(r) || r == '_' {
				tmp = append(tmp, r)
				myLen++
			}
		}
		s = string(tmp)
		// s = strings.TrimSpace(s)
		D(func() { fmt.Printf("NewHeaders: len of header %s: %d\n", s, len(s)) })
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

// D provides a switch for debug printing
func D(f func()) {
	if Debug {
		f()
	}
}
