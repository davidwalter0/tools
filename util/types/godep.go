package types // import "github.com/davidwalter0/tools/util/types"

type Godep struct {
	Name     string
	Version  string
	Revision string
}

type GodepLock Godep

type GodepLockConfig struct {
	Imports []GodepLock `json:"imports"`
}

type GodepConfig struct {
	Package string  `json:"package"`
	Import  []Godep `json:"import"`
}
