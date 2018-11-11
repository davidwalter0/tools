package time // import "github.com/davidwalter0/tools/x/time"

import (
	"fmt"
	"testing"
)

var debug = true

type UnixDateSplit struct {
	tstr string
	nstr string
	t    int64
	n    int64
}

type Table struct {
	UnixDate    string
	RFC1123     string
	DefaultDate string
	UnixDateSplit
	UnixTime
}

// Table of test inputs and expected values
var Tables = []Table{
	{
		"1461621630014",
		"2016-04-25 18:00:30 -0400 EDT",
		"Mon, 25 Apr 2016 18:00:30 EDT",
		UnixDateSplit{
			"1461621630",
			"014",
			1461621630,
			14,
		},
		1461621630014,
	},
	{
		"1461621630014",
		"2016-04-25 18:00:30 -0500 CDT",
		"Mon, 25 Apr 2016 18:00:30 CDT",
		UnixDateSplit{
			"1461621630",
			"014",
			1461621630,
			14,
		},
		1461621630014,
	},
}

func Test_UnixTimeFunctions(t *testing.T) {
	var u UnixDateSplit
	var unixTime UnixTime

	for _, itm := range Tables {
		u = itm.UnixDateSplit
		if debug {
			fmt.Println(u)
		}
		unixTime = UnixTime(itm.UnixTime.ToTime().Unix()*1000 + u.n)
		if unixTime != itm.UnixTime {
			t.Errorf("UnixTime() was incorrect, want %d got: %d.", itm.UnixTime, unixTime)
		}

		unixTime = UnixTime(itm.UnixTime.ToTime().Unix())
		if unixTime != UnixTime(itm.UnixDateSplit.t) {
			t.Errorf("UnixTime() was incorrect, want %d got: %d.", itm.UnixDateSplit.t, unixTime)
		}
		u = itm.UnixDateSplit
		unixTime = UnixTime(UnixTimeWithMilli(u.t, u.n).Unix()*1000 + u.n)
		if unixTime != itm.UnixTime {
			t.Errorf("UnixTime() was incorrect, want %d got: %d.", itm.UnixTime, unixTime)
		}
		if debug {
			fmt.Println(UnixTimeParseString(itm.UnixDate))
		}
		{
			l, r := UnixTimeMsResolutionStr2Int(itm.UnixDate)
			if debug {
				fmt.Printf(">>>  %10d   :%03d\n", l, r)
				fmt.Printf("time %10d:ms:%03d\n", l, r)
			}
			timeString := UnixTimeStringWithMsToPrintable(itm.UnixDate)
			ln := len(timeString)
			if timeString[:ln-4] != itm.DefaultDate[:ln-4] {
				t.Errorf("UnixTime() was incorrect, want %s got: %s.", itm.DefaultDate[:ln-4], timeString[:ln-4])
			}
			if debug {
				fmt.Println(UnixTimeStringWithMsToPrintable(itm.UnixDate))
			}
		}

	}
}
