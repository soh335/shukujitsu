package shukujitsu

import (
	"fmt"
	"io"
	"time"
)

// CLIOption represents an option for CLI.
type CLIOption struct {
	Quiet      bool
	Date       string
	DateFormat string
	Writer     io.Writer
}

// RunCLI runs shukujitsu CLI.
// return values
// 0 : is Shukujitsu
// 1 : is not Shukujitsu
// 2 : parse error of -d (date)
func RunCLI(opt *CLIOption) int {
	var (
		now time.Time
		err error
	)
	if opt.Date == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(opt.DateFormat, opt.Date)
		if err != nil {
			fmt.Fprintln(opt.Writer, err)
			return 2
		}
	}

	name, ok := dates[now.Format("2006/1/2")]
	if !opt.Quiet {
		fmt.Fprintln(opt.Writer, now.Format(opt.DateFormat), name)
	}
	if !ok {
		return 1
	}
	return 0
}
