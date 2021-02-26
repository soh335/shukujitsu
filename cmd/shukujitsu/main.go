package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/soh335/shukujitsu"
)

var (
	quiet       bool
	date        string
	dateFormat  = "2006-01-02"
	showVersion bool
	version     = "current"
)

func init() {
	flag.BoolVar(&quiet, "quiet", false, "quiet")
	flag.StringVar(&date, "date", "", "date (format "+dateFormat+")")
	flag.BoolVar(&showVersion, "version", false, "show version")
}

func main() {
	flag.Parse()
	if showVersion {
		fmt.Println("shukujitsu", version)
		return
	}
	code := shukujitsu.RunCLI(&shukujitsu.CLIOption{
		Quiet:      quiet,
		Date:       date,
		DateFormat: dateFormat,
		Writer:     os.Stdout,
	})
	os.Exit(code)
}
