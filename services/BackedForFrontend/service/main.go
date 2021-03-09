/**
* Main
* @author  Diana Lucia Serna Higuita
 */

package main

import (
	API "github.com/TMDBHydra/BackedForFrontend/pkg/api"
	Config "github.com/TMDBHydra/BackedForFrontend/pkg/config"
	Errors "github.com/TMDBHydra/BackedForFrontend/pkg/errors"
)

/**
* Main
* Entry point to the application
 */
func main() {
	/* the config file is located in the same folder where the application
	is runnig config/config.json */
	Config.ReadConfigFile("config/config.json")
	if Errors.LoadDatabaseCredentials() == false {
		Config.LogErrors = false
	}

	API.Main()
}
