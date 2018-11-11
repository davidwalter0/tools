package strings

import (
	"strings"
)

// DedupTrimmer (dedup delimiters and trim) has utilities that split,
// strip of c and return an array, a string without c, or an elemment
// of the []string splits, IOW compress duplicate delims and remove
// from begin and end
type DedupTrimmer interface {
	String() string
	Split(rune) []string
	Uniq(rune) string
	Last(rune) string
	First(rune) string
	Nth(c rune, nth int) string
}

type DedupTrim string

// Stringer for a DedupTrim
func (nc DedupTrim) String() string { return string(nc) }

// Uniq remove duplicates of c from DedupTrim
func (nc DedupTrim) Uniq(c rune) string {
	return strings.Join(nc.Split(c), string(c))
}

// Split remove duplicates of c from DedupTrim
func (nc DedupTrim) Split(c rune) []string {
	var has = strings.Split(string(nc), string(c))
	var split []string
	for _, segment := range has {
		if len(segment) > 0 {
			split = append(split, segment)
		}
	}
	return split
}

// Last returns the last element after split & trimming on c
func (nc DedupTrim) Last(c rune) string {
	var split []string = nc.Split(c)
	if len(split) > 0 {
		return split[len(split)-1]
	}
	return ""
}

// First returns the first element after split & trimming on c
func (nc DedupTrim) First(c rune) string {
	var split []string = nc.Split(c)
	if len(split) > 0 {
		return split[0]
	}
	return ""
}

// Nth returns the nth element after split & trimming on c
func (nc DedupTrim) Nth(c rune, nth int) string {
	var split []string = nc.Split(c)
	if len(split) > nth {
		return split[nth]
	}
	return ""
}
