/**
* Config tests
* @author  Diana Lucia Serna Higuita
* Test all the functions: $ go test ./...
 */
package config

import (
	"testing"
)

/**
*  Test load config file
 */
func TestConfig(t *testing.T) {
	ReadConfigFile("../../config/config.json")
	if len(PortRestAPI) == 0 {
		t.Errorf("Error the config json with the variable PortRestAPI")
	}

	if LogErrors != false && LogErrors != true {
		t.Errorf("Error the config json with the variable LogErrors")
	}
}
