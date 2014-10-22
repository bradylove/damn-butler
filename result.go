package main

import (
	"encoding/xml"
	"regexp"
)

type Result struct {
	XMLName  xml.Name  `xml:"Projects"`
	Projects []Project `xml:"Project"`
}

func (r *Result) ShowAll() {
	for _, p := range r.Projects {
		p.PrintMinimal()
	}
}

func (r *Result) ShowFiltered(filter string) {
	exp := regexp.MustCompile(filter)

	for _, p := range r.Projects {
		if exp.MatchString(p.Name) {
			p.PrintMinimal()
		}
	}
}
