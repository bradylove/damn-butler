package main

import (
	"fmt"
	"github.com/buger/goterm"
)

const (
	BUILDING_CHAR string = string(0x27F2)
	FAILURE_CHAR  string = string(0x2717)
	UNKNOWN_CHAR  string = string(0x26A0)
	SUCCESS_CHAR  string = string(0x2714)
)

type Project struct {
	WebURL        string `xml:"webUrl,attr"`
	Name          string `xml:"name,attr"`
	Label         string `xml:"lastBuildLabel,attr"`
	LastBuildTime string `xml:"lastBuildTime,attr"`
	Status        string `xml:"lastBuildStatus,attr"`
	Activity      string `xml:"activity,attr"`
}

func (p *Project) PrintMinimal() {
	fmt.Println(p.StatusText(), p.Name)
}

func (p *Project) StatusText() string {
	return goterm.Color(p.GetStatusIcon()+" "+p.GetStatusText(), p.GetStatusColor())
}

func (p *Project) GetStatusText() string {
	if p.Activity == "Building" {
		return p.Activity
	} else {
		return p.Status
	}
}

func (p *Project) GetStatusColor() int {
	switch p.Status {
	case "Success":
		return goterm.GREEN
	case "Failure":
		return goterm.RED
	default:
		return goterm.YELLOW
	}
}

func (p *Project) GetStatusIcon() string {
	switch p.GetStatusText() {
	case "Success":
		return SUCCESS_CHAR
	case "Failure":
		return FAILURE_CHAR
	case "Building":
		return BUILDING_CHAR
	default:
		return UNKNOWN_CHAR
	}
}
