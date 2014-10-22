package main

import (
	"encoding/xml"
	"regexp"
)

type Result struct {
	XMLName  xml.Name  `xml:"Projects"`
	Projects []Project `xml:"Project"`
}

func (r *Result) ShowAll() []string {
	list := make([]string, 0)

	for _, p := range r.Projects {
		list = append(list, p.DisplayText())
	}

	return list
}

func (r *Result) GetAll() []string {
	list := make([]string, 0)

	for _, p := range r.Projects {
		list = append(list, p.DisplayText())
	}

	return list
}

func (r *Result) ShowFiltered(filter string) {
	exp := regexp.MustCompile(filter)

	for _, p := range r.Projects {
		if exp.MatchString(p.Name) {
			p.PrintMinimal()
		}
	}
}
