package types // import "github.com/davidwalter0/tools/util/types"

// VersionMap map[package name] version
type VersionMap map[string]string

// FinalDepMap map of final dependencies after overrides
type FinalDepMap map[string]string

type GlideDep struct {
	Package string
	Version string
}

type GlideLockDep struct {
	Name    string
	Version string
}

type GlideLockConfig struct {
	Imports []GlideLockDep `json:"imports"`
}

type GlideConfig struct {
	Package string     `json:"package"`
	Import  []GlideDep `json:"import"`
}
