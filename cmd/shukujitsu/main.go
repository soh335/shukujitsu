package main

import (
	"flag"
	"os"

	"github.com/soh335/shukujitsu"
)

var (
	quiet      bool
	date       string
	dateFormat = "2006-01-02"
)

func init() {
	flag.BoolVar(&quiet, "quiet", false, "quiet")
	flag.StringVar(&date, "date", "", "date (format "+dateFormat+")")
}

func main() {
	flag.Parse()
	code := shukujitsu.RunCLI(&shukujitsu.CLIOption{
		Quiet:      quiet,
		Date:       date,
		DateFormat: dateFormat,
		Writer:     os.Stdout,
	})
	os.Exit(code)
}
