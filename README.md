damn-butler (0.1.0)
============================================

damn-butler is a CLI tool for monitoring Jenkins build jobs using the Jenkins
CC file (usually found at http://<yourdomain>/cc.xml)

### Building
====================

With Go 1.2 or newer installed run the following to install damn-butler

    go get github.com/bradylove/damn-butler
    go install github.com/bradylove/damn-butler

### Usage
====================

Get help!

    $ damn-butler help

    usage: damn-butler <command> [<flags>] [<args> ...]

    A command line Jenkins build monitoring app

    Flags:
      --help  Show help.

    Commands:
      help [<command>]
        Show help for a command.

    all
      Show the status for all available projects

    filter [<flags>] <filter>
      Show the status for projects whose names start with given string

    host add <URL to cc file>
      Adds a new Jenkins host to monitor

    host list
      Lists all Jenkins hosts that are being monitored

Add a new Jenkins CC file to watch:

    damn-butler host add http://<yourdomain>/cc.xml

List all watched hosts:

    damn-butler host list

Currently removing a host is not supported, however you can manually remove a host
by modifying the JSON settings file (`.damn-butler`) which can be found in your
users home directory.

Show all projects statuses:

    damn-butler all

Show filtered list of projects statuses:

    damn-butler filter <Regex for filtering>

Continuously monitor filtered list of projects (current refresh time is every 5 seconds):

    damn-butler filter -w <Regex for filtering>

### License
==============================
MIT License
