package main

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"time"
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

func (r *Result) ShowFiltered(filter string, watch bool) {
	exp := regexp.MustCompile(filter)

	i := 0

	for _, p := range r.Projects {
		if exp.MatchString(p.Name) {
			i += p.PrintMinimal()
		}
	}

	if !watch {
		return
	}

	time.Sleep(time.Second * 5)

	Backspace(i)
}

func Backspace(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("\b")
	}
}
