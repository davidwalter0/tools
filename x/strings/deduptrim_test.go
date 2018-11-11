package strings_test

import (
	"fmt"
	"testing"

	"github.com/davidwalter0/tools/x/strings"
)

var char = ' '
var uniqStrings = [][]string{
	//                                        first  last   nth(n==1)
	[]string{"abc def xyz", "abc   def  xyz", "abc", "xyz", "def"},
	[]string{"abc def ghi", "abc def  ghi  ", "abc", "ghi", "def"},
	[]string{"abc def ghi", "   abc def ghi", "abc", "ghi", "def"},
	[]string{"abc def ghi", "  abc def ghi ", "abc", "ghi", "def"},
	[]string{"abc", "  abc                 ", "abc", "abc", ""},
	[]string{"abc def", "abc def           ", "abc", "def", "def"},
}

func verbose() bool {
	return testing.Verbose()
}

func TupleParts(tuple []string) (first, last, nth string) {
	first, last, nth = tuple[2], tuple[3], tuple[4]
	return
}

func Nth(t *testing.T, tuple []string, n int) string {
	if n < len(tuple) {
		// 	t.Fatalf("Illegal nth element, n (%d) > len(tuple) %d", n, len(tuple))
		// }
		return tuple[n]
	}
	return ""
}

func InterfaceTupleParts(tuple []string) (first, last, nth strings.DedupTrimmer) {
	first, last, nth =
		(*strings.DedupTrim)(&tuple[2]),
		(*strings.DedupTrim)(&tuple[3]),
		(*strings.DedupTrim)(&tuple[4])

	return
}
func InterfaceLhsRhs(tuple []string) (lhs, rhs strings.DedupTrimmer) {
	lhs, rhs =
		(*strings.DedupTrim)(&tuple[0]),
		(*strings.DedupTrim)(&tuple[1])

	return
}

// Uniq remove duplicates of c from strings.DedupTrim
func TestUniq(t *testing.T) {
	for _, tuple := range uniqStrings {
		lhs, rhs := tuple[0], strings.DedupTrim(tuple[1])
		if lhs != rhs.Uniq(char) {
			t.Fatalf(`Error wanted: "%s" got "%s"`, lhs, rhs)
		}
	}
}

// Split remove duplicates of c from strings.DedupTrim
func TestSplit(t *testing.T) {
	for _, tuple := range uniqStrings {
		lhs, rhs := strings.DedupTrim(tuple[0]).Split(char), strings.DedupTrim(tuple[1]).Split(char)
		if len(lhs) != len(rhs) {
			t.Fatalf(`Error wanted: "%s" got "%s"`, lhs, rhs)
		}
		for i, v := range lhs {
			if v != rhs[i] {
				t.Fatalf(`Error wanted: "%s" got "%s"`, v, rhs[i])
			}
		}
	}
}

// Last returns the last element after splitting on c
func TestLast(t *testing.T) {
	for _, tuple := range uniqStrings {
		lhs, rhs := strings.DedupTrim(tuple[0]).Last(char), strings.DedupTrim(tuple[1]).Last(char)
		_ /*first*/, last, _ /*nth*/ := TupleParts(tuple)
		if verbose() {
			fmt.Printf(`want: "%s" got "%s"
`, lhs, rhs)
		}
		if lhs != rhs {
			t.Fatalf(`Error wanted: "%s" got "%s"`, lhs, rhs)
		}
		if verbose() {
			fmt.Printf(`last want: "%s" got "%s"
`, lhs, last)
		}
		if lhs != rhs {
			t.Fatalf(`Error last wanted: "%s" got "%s"
`, lhs, last)
		}
	}
}

// First returns the first element after splitting on c
func TestFirst(t *testing.T) {
	for _, tuple := range uniqStrings {
		lhs, rhs := strings.DedupTrim(tuple[0]).First(char), strings.DedupTrim(tuple[1]).First(char)

		first, _ /*last*/, _ /*nth*/ := TupleParts(tuple)
		if verbose() {
			fmt.Printf(`want: "%s" got "%s"
`, lhs, rhs)
		}
		if lhs != rhs {
			t.Fatalf(`Error wanted: "%s" got "%s"`, lhs, rhs)
		}
		if verbose() {
			fmt.Printf("first wanted: %s, got %s\n", lhs, first)
		}
		if lhs != rhs {
			t.Fatalf("Error first wanted: %s, got %s\n", lhs, first)
		}
	}
}

// Nth returns the nth element after splitting on c
func TestNth(t *testing.T) {
	for _, tuple := range uniqStrings {
		lhs, rhs := strings.DedupTrim(tuple[0]).Nth(char, 1), strings.DedupTrim(tuple[1]).Nth(char, 1)

		_ /*first*/, _ /*last*/, nth := TupleParts(tuple)
		if verbose() {
			fmt.Printf(`want: "%s" got "%s"
`, lhs, rhs)
		}
		if lhs != rhs {
			t.Fatalf(`Error wanted: "%s" got "%s"`, lhs, rhs)
		}

		if verbose() {
			fmt.Printf(`nth %s Nth(t,tuple,1) %s %v wanted: "%s" got "%s"
`, nth, Nth(t, tuple, 1), tuple, lhs, nth)
		}
		if nth != Nth(t, strings.DedupTrim(tuple[1]).Split(char), 1) {
			t.Fatalf(`Error wanted: "%s" got "%s"`, lhs, nth)
		}
	}
}

// Uniq remove duplicates of c from strings.DedupTrim
func TestIfaceDedupTrimmerUniq(t *testing.T) {
	for _, tuple := range uniqStrings {
		var split strings.DedupTrimmer = (*strings.DedupTrim)(&tuple[1])
		lhs, rhs := tuple[0], split
		if lhs != rhs.Uniq(char) {
			t.Fatalf(`Error wanted: "%s" got "%s"`, lhs, rhs)
		}
	}
}

// Split remove duplicates of c from strings.DedupTrim
func TestIfaceDedupTrimmerSplit(t *testing.T) {
	for _, tuple := range uniqStrings {
		var split strings.DedupTrimmer = (*strings.DedupTrim)(&tuple[1])
		lhs, rhs := strings.DedupTrim(tuple[0]).Split(char), split.Split(char)
		if len(lhs) != len(rhs) {
			t.Fatalf(`Error wanted: "%s" got "%s"`, lhs, rhs)
		}
		for i, v := range lhs {
			if v != rhs[i] {
				t.Fatalf(`Error wanted: "%s" got "%s"`, v, rhs[i])
			}
		}
	}
}

// Last returns the last element after splitting on c
func TestIfaceDedupTrimmerLast(t *testing.T) {
	for _, tuple := range uniqStrings {
		l, r := InterfaceLhsRhs(tuple)
		lhs, rhs := l.Last(char), r.Last(char)
		_ /*first*/, last, _ /*nth*/ := InterfaceTupleParts(tuple)
		if verbose() {
			fmt.Printf(`want: "%s" got "%s"
`, lhs, rhs)
		}
		if lhs != rhs {
			t.Fatalf(`Error wanted: "%s" got "%s"`, lhs, rhs)
		}
		if verbose() {
			fmt.Printf(`last want: "%s" got "%s"
`, lhs, last)
		}
		if lhs != rhs {
			t.Fatalf(`Error last wanted: "%s" got "%s"
`, lhs, last)
		}
	}
}

// First returns the first element after splitting on c
func TestIfaceDedupTrimmerFirst(t *testing.T) {
	for _, tuple := range uniqStrings {
		l, r := InterfaceLhsRhs(tuple)
		lhs, rhs := l.First(char), r.First(char)

		first, _ /*last*/, _ /*nth*/ := InterfaceTupleParts(tuple)
		if verbose() {
			fmt.Printf(`want: "%s" got "%s"
`, lhs, rhs)
		}
		if lhs != rhs {
			t.Fatalf(`Error wanted: "%s" got "%s"`, lhs, rhs)
		}
		if verbose() {
			fmt.Printf("first wanted: %s, got %s\n", lhs, first)
		}
		if lhs != rhs {
			t.Fatalf("Error first wanted: %s, got %s\n", lhs, first)
		}
	}
}

// Nth returns the nth element after splitting on c
func TestIfaceDedupTrimmerNth(t *testing.T) {
	for _, tuple := range uniqStrings {
		l, r := InterfaceLhsRhs(tuple)
		lhs, rhs := l.Nth(char, 1), r.Nth(char, 1)

		_ /*first*/, _ /*last*/, nth := InterfaceTupleParts(tuple)
		if verbose() {
			fmt.Printf(`want: "%s" got "%s"
`, lhs, rhs)
		}
		if lhs != rhs {
			t.Fatalf(`Error wanted: "%s" got "%s"`, lhs, rhs)
		}

		if verbose() {
			fmt.Printf(`nth %s Nth(t,tuple,1) %s %v wanted: "%s" got "%s"
`, nth, Nth(t, tuple, 1), tuple, lhs, nth)
		}
		if nth.String() != Nth(t, strings.DedupTrim(tuple[1]).Split(char), 1) {
			t.Fatalf(`Error wanted: "%s" got "%s"`, lhs, nth)
		}
	}
}
