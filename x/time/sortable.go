package time // import "github.com/davidwalter0/toolsx/time"

import (
	// "fmt"
	"sort"
	"time"
)

type TimeSlice []time.Time

func (s TimeSlice) Ascending() TimeSlice {
	sort.Sort(s)
	return s
}
func (s TimeSlice) Descending() TimeSlice {
	sort.Sort(sort.Reverse(s))
	return s
}

func (s TimeSlice) Less(i, j int) bool      { return s[i].Before(s[j]) }
func (s TimeSlice) Swap(i, j int)           { s[i], s[j] = s[j], s[i] }
func (s TimeSlice) Len() int                { return len(s) }
func (s *TimeSlice) Add(datetime time.Time) { *s = append(*s, datetime) }

func SignalBackupFilename(datetime time.Time) string {
	var layout = "2006-01-02-15-04-05"
	var filename = "signal-" + datetime.Format(layout) + ".backup"
	return filename
}

func String2DateTimeTz(datetime string) (time.Time, error) {
	//	       2018-06-10-19-05-55
	layout := "2006-01-02-15-04-05.000Z"
	// str := "2014-11-12T11:45:26.371Z"

	// layout := "2006-01-02T15:04:05.000Z"
	// str := "2014-11-12T11:45:26.371Z"
	return time.Parse(layout, datetime)
}

func String2DateTime(datetime string) (time.Time, error) {
	//	       2018-06-10-19-05-55
	layout := "2006-01-02-15-04-05"
	// str := "2014-11-12T11:45:26.371Z"

	// layout := "2006-01-02T15:04:05.000Z"
	// str := "2014-11-12T11:45:26.371Z"
	return time.Parse(layout, datetime)
}

/*
var past = time.Date(2010, time.May, 18, 23, 0, 0, 0, time.Now().Location())
var present = time.Now()
var future = time.Now().Add(24 * time.Hour)

var dateSlice TimeSlice = []time.Time{present, future, past}

func main() {

	// fmt.Println("Past : ", past)
	// fmt.Println("Present : ", present)
	// fmt.Println("Future : ", future)

	fmt.Println("Before sorting : ")
	for i, date := range dateSlice {
		fmt.Printf("%3d %s\n", i, date)
	}

	// sort.Sort(dateSlice)
	dateSlice.Ascending()
	fmt.Println("After sorting : ")
	for i, date := range dateSlice {
		fmt.Printf("%3d %s\n", i, date)
	}

	// sort.Sort(sort.Reverse(dateSlice))
	dateSlice.Descending()
	fmt.Println("After REVERSE sorting : ")
	for i, date := range dateSlice {
		fmt.Printf("%3d %s\n", i, date)
	}

}
*/
