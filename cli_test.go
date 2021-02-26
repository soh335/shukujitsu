package shukujitsu

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

var cliTestCases = []struct {
	quiet  bool
	date   string
	output string
	code   int
}{
	{
		quiet:  false,
		date:   "2021-01-01",
		output: "2021-01-01 元日",
		code:   0,
	},
	{
		quiet:  false,
		date:   "2021-01-02",
		output: "2021-01-02 ",
		code:   1,
	},
	{
		quiet:  false,
		date:   "2021-01-0x",
		output: `parsing time "2021-01-0x" as "2006-01-02": cannot parse "0x" as "02"`,
		code:   2,
	},
	{
		quiet:  true,
		date:   "2021-01-01",
		output: "",
		code:   0,
	},
	{
		quiet:  true,
		date:   "2021-01-02",
		output: "",
		code:   1,
	},
	{
		quiet:  true,
		date:   "2021-01-0x",
		output: `parsing time "2021-01-0x" as "2006-01-02": cannot parse "0x" as "02"`,
		code:   2,
	},
}

func TestCLI(t *testing.T) {
	stdout := os.Stdout
	defer func() {
		os.Stdout = stdout
	}()
	for _, c := range cliTestCases {
		var b []byte
		buf := bytes.NewBuffer(b)
		code := RunCLI(&CLIOption{
			Quiet:      c.quiet,
			Date:       c.date,
			DateFormat: "2006-01-02",
			Writer:     buf,
		})
		if code != c.code {
			t.Errorf("unexpected code %d %v", code, c)
		}
		output := strings.TrimRight(buf.String(), "\n")
		if output != c.output {
			t.Errorf("unexpected output %s %v", output, c)
		}
	}
}
