package main

import (
	"encoding/xml"
	"fmt"
	"github.com/buger/goterm"
	"gopkg.in/alecthomas/kingpin.v1"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

const (
	BUILDING_CHAR string = string(0x27F2)
	FAILURE_CHAR  string = string(0x2717)
	UNKNOWN_CHAR  string = string(0x26A0)
	SUCCESS_CHAR  string = string(0x2714)
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

var (
	app = kingpin.New("damn-butler", "A command line Jenkins build monitoring app")

	allCmd = app.Command("all", "Show the status for all available projects")

	filterCmd = app.Command("filter", "Show the status for projects whose names start with given string")
	filterStr = filterCmd.Arg("filter", "String to filter projects by").Required().String()
)

func main() {
	results := FetchResults("<INSERT URL>")

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case allCmd.FullCommand():
		results.ShowAll()
	case filterCmd.FullCommand():
		results.ShowFiltered(*filterStr)
	}
}

func FetchResults(url string) Result {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	results := Result{}
	err = xml.Unmarshal(data, &results)
	if err != nil {
		panic(err)
	}

	return results
}
