package util // import "github.com/davidwalter0/tools/util"

import (
	"log"

	"github.com/pelletier/go-toml"

	. "github.com/davidwalter0/tools/util/types"
)

func NameVersion(pkg interface{}) (name, version string, ok bool) {
	var pkgMap map[string]interface{}
	if pkgMap, ok = pkg.(map[string]interface{}); ok {
		if name, ok = pkgMap["name"].(string); ok {
			if version, ok = pkgMap["version"].(string); !ok {
				version, ok = pkgMap["revision"].(string)
			}
		}
	}
	return
}

// ContentArrayFromToml
func ContentArrayFromToml(key, filename string) (ContentArray, error) {
	var err error
	var ok bool
	var ca ContentArray

	var tree *toml.Tree
	if tree, err = toml.LoadFile(filename); err != nil {
		return nil, err
	}

	var treeMap map[string]interface{}
	if treeMap = tree.ToMap(); treeMap != nil {
		if ca, ok = treeMap[key].([]interface{}); ok {
			// log.Printf("%v %T\n", ca, ca)
			return ca, nil
		}
	}
	return nil, nil
}

// LoadDepToml file return version map
func LoadToml(key, filename string) (m VersionMap) {
	var name string
	var version string
	var ok bool
	var err error
	var ca ContentArray
	var pkg interface{}

	if ca, err = ContentArrayFromToml(key, filename); err != nil {
		log.Fatal(err)
	}
	// log.Printf("%T %v\n", ca, ca)

	m = VersionMap{}
	for _, pkg = range ca {
		if name, version, ok = NameVersion(pkg); ok {
			m[name] = version
		}
	}
	return m
}
