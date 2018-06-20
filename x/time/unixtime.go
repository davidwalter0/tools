package time // import "github.com/davidwalter0/toolsx/time"

import (
	"fmt"
	"strconv"
	"time"
)

/*
const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)
    These are predefined layouts for use in Time.Format and time.Parse. The
    reference time used in the layouts is the specific time:

*/

const (
	DateTimeFormat = "2006-01-02-15-04-05"
)

type UnixTime int64

// ToTime returns time.Time from UnixTime
func (u UnixTime) ToTime() time.Time {
	return time.Unix(int64(u)/1000, (int64(u)%1000)*1000000)
}

// ToTime converts to time.Time and Formats using layout
func (u UnixTime) Format(layout string) string {
	tm := time.Unix(int64(u)/1000, (int64(u)%1000)*1000000)
	return tm.Format(layout)
}

// ToUnixTimeString converts to time.Time and Formats using the
// time.UnixDate constant = "Mon Jan _2 15:04:05 MST 2006"
func (u UnixTime) ToUnixTimeFormat() string {
	return u.Format(time.UnixDate)
}

// ToDateTimeString converts to time.Time and Formats using the
// DateTimeFormat = "2006-01-02-15-04-05"
func (u UnixTime) ToDateTimeString() string {
	return u.Format(DateTimeFormat)
}

// ToUnixTime takes seconds + milli seconds returning UnixTime
func ToUnixTime(t, m int64) UnixTime {
	return UnixTime(t*1000 + m)
}

// UnixTimeWithMilli create a time.Time from unix time t ( second resolution ) +
// milli seconds
func UnixTimeWithMilli(t, m int64) time.Time {
	return time.Unix(t, m*1000000)
}

// UnixTimeParseString unix string time parse
func UnixTimeParseString(ut string) (time, nano string) {
	return ut[:10], ut[10:]
}

// UnixTimeMsResolutionStr2Int unix string time parse
func UnixTimeMsResolutionStr2Int(ut string) (tv, ms int64) {
	t, n := ut[:10], ut[10:]
	tv, _ = strconv.ParseInt(t, 10, 64)
	ms, _ = strconv.ParseInt(n, 10, 64)
	return tv, ms
}

// UnixTimeStringWithMsToPrintable unix string time parse
func UnixTimeStringWithMsToPrintable(ut string) string {
	return fmt.Sprintf("%s", time.Unix(UnixTimeMsResolutionStr2Int(ut)).Format(time.RFC1123))
}
