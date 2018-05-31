package util // import "github.com/davidwalter0/tools/util"

import (
	"fmt"

	"encoding/json"
	"github.com/davidwalter0/transform"
	yaml "gopkg.in/yaml.v2"
)

// JSONify an object
func JSONify(data interface{}) string {
	var err error
	data, err = transform.TransformData(data)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	s, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	return string(s)
}

// YAMLify object to yaml string
func YAMLify(data interface{}) string {
	data, err := transform.TransformData(data)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	s, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	return string(s)
}
