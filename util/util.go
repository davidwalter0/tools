// -*- mode: go -*-
package util // import "github.com/davidwalter0/tools/util"

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/blang/semver"

	. "github.com/davidwalter0/tools/util/types"
)

const (
	// Latest version tag
	Latest = "latest"
	// VGoMod vgo default file versioning name
	VGoMod = "go.mod"
)

var firstDuplicate bool

func init() {
	firstDuplicate = true
}

// LoadFile from string return []byte
func LoadFile(filename string) []byte {
	var err error
	var b []byte

	if b, err = ioutil.ReadFile(filename); err != nil {
		log.Println(err)
	}
	return b
}

// FixVersionPrefix remove prefix chars
func FixVersionPrefix(version string) string {
	var v semver.Version
	var err error

	version = strings.TrimLeft(version, "^@>=~")
	if version != Latest && !IsHexHash(version) {
		v, err = semver.ParseTolerant(strings.TrimLeft(version, "^@>=~"))
		if err == nil {
			if len(version) > 0 && version[0] != 'v' && version != Latest {
				version = "v" + v.String()
			}
		} else {
			if strings.Index(err.Error(), "Version string empty") != 0 {
				version = Latest
			}
		}
	}
	return version
}

// Duplicates reports different versions proposed between glide.yaml
// and glide.lock for the same package name
func Duplicates(name, version, priorVersion string) {
	if firstDuplicate {
		fmt.Println("Duplicates")
		fmt.Printf("%-40.40s %-12.12s %-12.12s\n", "Package", "File Version", "Override")
		firstDuplicate = false
	}
	fmt.Printf("%-40.40s %-12.12s %-12.12s\n", name, version, priorVersion)
}

func IsHexHash(hash string) bool {
	var err error
	var rex *regexp.Regexp
	const hashRegex = "^[0-9a-f]{5,40}$"
	var match bool
	match, err = regexp.MatchString(hashRegex, hash)

	if rex, err = regexp.Compile(hashRegex); err != nil {
		log.Fatal(fmt.Sprintf("error %s %v", hashRegex, err))
	}
	var text = rex.Find([]byte(hash))
	// log.Println(len(text) == len(hash) && match, len(text), len(hash), match)
	return len(text) == len(hash) && match
}

// Merge versions return final map
func Merge(lock VersionMap, rhs VersionMap) (m VersionMap) {
	var duplicates bool
	var k string
	var v string
	var nv string
	var ok bool
	m = make(VersionMap)
	for k, v = range rhs {
		if nv, ok = lock[k]; ok {
			duplicates = true
			Duplicates(k, v, nv)
			v = nv
		}
		m[k] = FixVersionPrefix(v)
	}
	for k, v = range lock {
		m[k] = FixVersionPrefix(v)
	}
	if duplicates {
		fmt.Printf("\n\n")
	}
	return
}

func Empty() {
}

func WriteModule(ext, text, divertFilename, defaultFilename string, overwrite bool) {
	var err error
	// if app.VGoDivertExt == FauxID {
	if len(ext) > 0 {
		divertFilename = fmt.Sprintf("go.mod.%s", ext)
	}
	text += `
// local variables:
// mode: vgo
// end:
`
	if ext != "" {
		if err = ioutil.WriteFile(divertFilename, []byte(text), 0666); err != nil {
			log.Fatal(err)
		}
	} else if _, err := os.Stat(defaultFilename); os.IsNotExist(err) || overwrite {
		// If go.mod doesn't exist or overwrite, ok to write it
		if err = ioutil.WriteFile(defaultFilename, []byte(text), 0666); err != nil {
			log.Fatal(err)
		}
	} else if _, err := os.Stat(defaultFilename); err == nil {
		log.Fatal(fmt.Sprintf("File %s present and overwrite not set. Refusing to overwrite\n", defaultFilename))
	}
}

// PackageFromCwd builds a path using GOPATH assumption of
func PackageFromCwd() string {
	var Package string
	var err error

	if Package, err = os.Getwd(); err != nil {
		log.Fatal(err)
	}

	var gopath string
	var found bool
	if gopath, found = os.LookupEnv("GOPATH"); found {
		Package = strings.Replace(Package, gopath+"/src/", "", 1)
	}
	return Package
}

// local variables:
// mode: go
// end:
