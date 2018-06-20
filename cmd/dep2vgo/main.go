// build -tags netgo -ldflags '-w -s'
// install -tags netgo -ldflags '-w -s'
// Tomljson reads TOML and converts to JSON.
//
// Usage:
//   cat file.toml | tomljson > file.json
//   tomljson file1.toml > file.json
package main // import "github.com/davidwalter0/toolscmd/dep2vgo"

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/davidwalter0/go-cfg"
	. "github.com/davidwalter0/toolsutil"
	. "github.com/davidwalter0/toolsutil/types"
)

// type _ignore_ util.VersionMap

const (
	DepTomlFilename = "Gopkg.toml"
	DepLockFilename = "Gopkg.lock"
)

var FauxID = "dep2vgo"

/*
currently ignore
  branch and packages

# Gopkg.lock
[[projects]]
  branch = "master"
  name = "github.com/hanwen/go-fuse"
  packages = [
    "fuse",
    "fuse/nodefs",
    "fuse/pathfs",
    "splice"
  ]
  revision = "a9ddcb8a4b609500fc59c89ccc9ee05f00a5fefd"

---
# Gopkg.toml
[[constraint]]
  branch = "master"
  name = "golang.org/x/crypto"


*/

var firstDuplicate bool

// DepMap map[name]version
type DepMap VersionMap

// godepMap map[name]version
var godepMap = make(DepMap)

type Config struct {
	// Ignored for now
	// Output       string    `json:"-" doc:"ignore"`
	// Filename     string    `json:"filename"`
	Verbose      bool   `json:"verbose" doc:"write status info" default:"true"`
	Overwrite    bool   `json:"overwrite" doc:"replace an existing file with new output" default:"false"`
	VGoDivertExt string `json:"vgo-divert-ext" doc:"use empty flag use package name, else name go.mod.+'package-name' e.g: go.mod.dep2vgo or go.mod.+'vgo-divert-ext'" default:"dep2vgo"`
}

var app = &Config{}

func init() {
	firstDuplicate = true
	var err error
	if err = cfg.Process("", app); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func main() {
	var TomlKey = "constraint"
	var LockKey = "projects"
	var TomlFile = DepTomlFilename
	var LockFile = DepLockFilename

	var lhs = LoadToml(LockKey, LockFile)
	var rhs = LoadToml(TomlKey, TomlFile)
	// fmt.Println(YAMLify(lhs))
	// fmt.Println(YAMLify(rhs))
	var final = Merge(lhs, rhs)
	// fmt.Println(YAMLify(final))

	var name string
	var version string
	var first = true
	var goModText = ""
	var Package = PackageFromCwd()
	goModText += fmt.Sprintf("module %s\n\n", Package)
	goModText += fmt.Sprintf("require (\n")
	var keys []string
	var key string
	for key, _ = range final {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, name = range keys {
		version = final[name]
		version = FixVersionPrefix(version)
		goModText += fmt.Sprintf("\t%s %s\n", name, version)
		if app.Verbose {
			if first {
				fmt.Println("Final")
				fmt.Printf("%-40.40s %-12.12s\n", "Package", "Version")
				first = false
			}
			fmt.Printf("%-40.40s %-12.12s\n", name, version)
		}
	}

	goModText += fmt.Sprintf(")\n\n")

	var packageName = filepath.Base(Package)
	var VGoModDivertName = fmt.Sprintf("go.mod.%s", packageName)

	var ext = app.VGoDivertExt
	if ext == "" || ext == FauxID {
		ext = ""
	}
	WriteModule(ext, goModText, VGoModDivertName, VGoMod, app.Overwrite)
}
