package strings

import (
	"strings"
)

// UniqifySplitString has some utilities that split, strip of c and
// return an array, a string without c, or an elemment of the []string
// splits
type UniqifySplitString string

// Stringer for a UniqifySplitString
func (nc UniqifySplitString) String() string { return string(nc) }

// Uniq remove duplicates of c from UniqifySplitString
func (nc UniqifySplitString) Uniq(c rune) string {
	return strings.Join(nc.Split(c), string(c))
}

// Split remove duplicates of c from UniqifySplitString
func (nc UniqifySplitString) Split(c rune) []string {
	var has = strings.Split(string(nc), string(c))
	var split []string
	for _, segment := range has {
		if len(segment) > 0 {
			split = append(split, segment)
		}
	}
	return split
}

// FirstUniq returns the last element after splitting on c
func (nc UniqifySplitString) LastUniq(c rune) string {
	var split []string = nc.Split(c)
	if len(split) > 0 {
		return split[len(split)-1]
	}
	return ""
}

// FirstUniq returns the first element after splitting on c
func (nc UniqifySplitString) FirstUniq(c rune) string {
	var split []string = nc.Split(c)
	if len(split) > 0 {
		return split[0]
	}
	return ""
}

// NthUniq returns the nth element after splitting on c
func (nc UniqifySplitString) NthUniq(c rune, nth int) string {
	var split []string = nc.Split(c)
	if len(split) > nth-1 {
		return split[nth-1]
	}
	return ""
}
