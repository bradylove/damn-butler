package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v1"
	"os"
)

var (
	app = kingpin.New("damn-butler", "A command line Jenkins build monitoring app")

	allCmd = app.Command("all", "Show the status for all available projects")

	filterCmd = app.Command("filter", "Show the status for projects whose names start with given string")
	filterStr = filterCmd.Arg("filter", "String to filter projects by").Required().String()

	hostCmd     = app.Command("host", "Host operations")
	hostAddCmd  = hostCmd.Command("add", "Adds a new Jenkins host to monitor")
	hostAddStr  = hostAddCmd.Arg("cc", "The URL to the Jenkins CC file used for monitoring projects").Required().String()
	hostListCmd = hostCmd.Command("list", "Lists all Jenkins hosts that are being monitored")

	config *Config
)

func main() {
	var err error
	config, err = LoadConfig()
	if err != nil {
		panic(err)
	}

	results, _ := FetchAllResults(config.Hosts)

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case allCmd.FullCommand():
		RequireHost()
		results.ShowAll()
	case filterCmd.FullCommand():
		RequireHost()
		results.ShowFiltered(*filterStr)
	case hostAddCmd.FullCommand():
		config.AddHost(*hostAddStr)
	case hostListCmd.FullCommand():
		config.PrintHostList()
	}
}

func RequireHost() {
	if !config.HasHosts() {
		fmt.Println("This command requires that you first add a host")
		os.Exit(1)
	}
}
