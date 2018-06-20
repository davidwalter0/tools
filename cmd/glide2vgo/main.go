// +build !windows

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/davidwalter0/go-cfg"
	. "github.com/davidwalter0/tools/util"
	. "github.com/davidwalter0/tools/util/types"

	yaml "gopkg.in/yaml.v2"
)

/*
// GlideConfigMap map[name]version
type GlideConfigMap map[string]string

// GlideLockMap map[name]version
type GlideLockMap map[string]string
*/
// // map[name]version
var glideConfigMap = make(VersionMap)

// // map[name]version
var glideLockMap = make(VersionMap)

const FauxID = "tools"

type Config struct {
	Verbose      bool   `json:"verbose"                              doc:"write status info" default:"true"`
	Overwrite    bool   `json:"overwrite"                            doc:"replace an existing file with new output" default:"false"`
	VGoDivertExt string `json:"vgo-divert-ext" yaml:"vgo-divert-ext" doc:"use empty flag use package name, else name go.mod.+'package-name' e.g: go.mod.tools or go.mod.+'vgo-divert-ext'" default:"tools"`
}

var app = &Config{}

func init() {
	var err error
	if err = cfg.Process("", app); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

// GlideConfigFromYAML returns an instance of GlideConfig from YAML
func GlideConfigFromYAML(yml []byte) (Package string, m VersionMap, e error) {
	cfg := &GlideConfig{}

	if e = yaml.Unmarshal([]byte(yml), &cfg); e != nil {
		log.Println(e)
		return "", nil, e
	}

	var version string
	var ok bool
	var dep GlideDep
	m = make(VersionMap)
	for _, dep = range cfg.Import {
		if version, ok = m[dep.Package]; ok {
			version = FixVersionPrefix(version)
			dep.Version = FixVersionPrefix(dep.Version)
			Duplicates(dep.Package, dep.Version, version)
			m[dep.Package] = version
		} else {
			dep.Version = FixVersionPrefix(dep.Version)
			m[dep.Package] = dep.Version
		}
	}
	return cfg.Package, m, e
}

// GlideLockConfigFromYAML returns an instance of GlideConfig from YAML
func GlideLockConfigFromYAML(yml []byte) (m VersionMap, e error) {
	cfg := &GlideLockConfig{}
	if e = yaml.Unmarshal([]byte(yml), &cfg); e != nil {
		log.Println(e)
		return nil, e
	}

	var version string
	var ok bool
	var dep GlideLockDep
	m = make(VersionMap)
	for _, dep = range cfg.Imports {
		if version, ok = m[dep.Name]; ok {
			version = FixVersionPrefix(version)
			dep.Version = FixVersionPrefix(dep.Version)
			Duplicates(dep.Name, dep.Version, version)
			m[dep.Name] = version
		} else {
			dep.Version = FixVersionPrefix(dep.Version)
			m[dep.Name] = dep.Version
		}
	}
	return m, e
}

// LoadGlideConfig file return GlideConfig
func LoadGlideConfig(filename string) (Package string, c VersionMap) {
	var err error
	var b = LoadFile(filename)
	if Package, c, err = GlideConfigFromYAML(b); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return Package, c
}

// LoadGlideLock file return GlideLockConfig
func LoadGlideLock(filename string) (c VersionMap) {
	var err error
	var b = LoadFile(filename)
	if c, err = GlideLockConfigFromYAML(b); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return c
}

func main() {
	var goModText = ""
	var Package, c = LoadGlideConfig("glide.yaml")
	var l = LoadGlideLock("glide.lock")
	var final = Merge(c, l)

	goModText += fmt.Sprintf("module %s\n", Package)
	goModText += fmt.Sprintf("require (\n")

	var name string
	var version string

	var first = true
	for name, version = range final {
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
