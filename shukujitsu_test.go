package shukujitsu

import (
	"testing"
	"time"
)

func TestIsShukujitsu(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Error(err)
	}

	type spec struct {
		Input  time.Time
		Expect bool
	}

	specs := []spec{
		{
			Input:  time.Date(2020, 1, 13, 0, 0, 0, 0, loc),
			Expect: true,
		},
		{
			Input:  time.Date(2020, 1, 14, 0, 0, 0, 0, loc),
			Expect: false,
		},
		{
			Input:  time.Date(2019, 5, 4, 0, 0, 0, 0, loc),
			Expect: true,
		},
	}

	for _, s := range specs {
		if IsShukujitsu(s.Input) != s.Expect {
			t.Errorf("IsShukujitsu(%s) expect %v", s.Input, s.Expect)
		}
	}
}
