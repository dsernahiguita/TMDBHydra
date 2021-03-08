/**
* Config tests
* @author  Diana Lucia Serna Higuita
* Test funcion of all packages
 */
package main

import (
	"testing"

	Config "github.com/TMDBHydra/CliGo/pkg/config"
)

/*****************************************************************************
*  TEST FUNCTIONS CONFIG PACKAGE
*****************************************************************************/

/**
* Test config file
 */
func TestConfig(t *testing.T) {
	if Config.LogErrors != false && Config.LogErrors != true {
		t.Errorf("Error the config json with the variable LogErrors")
	}
}
