package main

import (
	"fmt"
	// "github.com/buger/goterm"
	"github.com/hoffoo/termboxutil"
	"github.com/nsf/termbox-go"
	"gopkg.in/alecthomas/kingpin.v1"
	"os"
	"time"
)

var (
	app = kingpin.New("damn-butler", "A command line Jenkins build monitoring app")

	allCmd          = app.Command("all", "Show the status for all available projects")
	allWatch        = allCmd.Flag("watch", "Continue to watch projects").Short('w').Bool()
	allWatchSeconds = allCmd.Flag("refresh", "Status refresh frequency").Short('s').Default("5s").Duration()

	filterCmd          = app.Command("filter", "Show the status for projects whose names start with given string")
	filterStr          = filterCmd.Arg("filter", "String to filter projects by").Required().String()
	filterWatch        = filterCmd.Flag("watch", "Continue to watch projects").Short('w').Bool()
	filterWatchSeconds = filterCmd.Flag("refresh", "Status refresh frequency").Short('s').Default("5s").Duration()

	hostCmd     = app.Command("host", "Host operations")
	hostAddCmd  = hostCmd.Command("add", "Adds a new Jenkins host to monitor")
	hostAddStr  = hostAddCmd.Arg("URL to cc file", "The URL to the Jenkins CC file used for monitoring projects").Required().String()
	hostListCmd = hostCmd.Command("list", "Lists all Jenkins hosts that are being monitored")

	config *Config
)

func main() {
	var err error
	config, err = LoadConfig()
	if err != nil {
		panic(err)
	}

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case allCmd.FullCommand():
		if *allWatch {
			ShowAllWatch()
		} else {
			ShowAll()
		}
	case filterCmd.FullCommand():
		if *filterWatch {
			ShowFilteredWatch()
		} else {
			ShowFiltered()
		}
	case hostAddCmd.FullCommand():
		config.AddHost(*hostAddStr)
	case hostListCmd.FullCommand():
		config.PrintHostList()
	}
}

func ShowAll() {
	RequireHost()

	results, err := FetchAllResults(config.Hosts)
	if err != nil {
		fmt.Println("Failed to fetch results:", err)
	}

	WithTermbox(results.GetAll)
}

func ShowAllWatch() {
	RequireHost()

	for {
		results, err := FetchAllResults(config.Hosts)
		if err != nil {
			fmt.Println("Failed to fetch results:", err)
		}

		results.ShowAll()

		time.Sleep(*filterWatchSeconds)
	}
}

func ShowFiltered() {
	RequireHost()

	results, err := FetchAllResults(config.Hosts)
	if err != nil {
		fmt.Println("Failed to fetch results:", err)
	}

	results.ShowFiltered(*filterStr)
}

func ShowFilteredWatch() {
	RequireHost()

	for {
		results, err := FetchAllResults(config.Hosts)
		if err != nil {
			fmt.Println("Failed to fetch results:", err)
		}

		results.ShowFiltered(*filterStr)

		time.Sleep(*filterWatchSeconds)
	}
}

func RequireHost() {
	if !config.HasHosts() {
		fmt.Println("This command requires that you first add a host")
		os.Exit(1)
	}
}

func WithTermbox(action func() []string) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	screen := termboxutil.Screen{}
	window := screen.NewWindow(
		termbox.ColorWhite,
		termbox.ColorDefault,
		termbox.ColorBlack,
		termbox.ColorGreen,
	)
	window.Scrollable(true)

	go func() {
		window.Draw(action())
		screen.Focus(&window)
		termbox.Flush()

		// TODO: sleep for refresh rate duration
	}()

	window.CatchEvent = func(event termbox.Event) {
		if event.Ch == 'j' || event.Key == termbox.KeyArrowDown {
			window.NextRow()
		} else if event.Ch == 'k' || event.Key == termbox.KeyArrowUp {
			window.PrevRow()
		} else if event.Ch == 'q' || event.Key == termbox.KeyEsc {
			termbox.Close()
			os.Exit(0)
		}

		window.Redraw()
		termbox.Flush()
	}

	screen.Loop()
}
