/**
* Config
* Read the config file config.json
* @author  Diana Lucia Serna Higuita
 */

package config

import (
	"encoding/json"
	"io/ioutil"

	Errors "github.com/TMDBHydra/BackedForFrontend/pkg/errors"
)

type Config struct {
	LogErrors          bool   `json:"logErrors"`
	PortRestAPI        string `json:"portRestAPI"`
	BackendServiceTMDB string `json:"backendServiceTMDB"`
	ApiKeyTMDB         string `json:"apiKeyTMDB"`
}

/* Log errors: the errors will be save into the path logs */
var LogErrors bool

/* Port rest api: port where services can be accessed */
var PortRestAPI string

/* End point backend service The movie DB */
var BackendServiceTMDB string

/* End point backend service The movie DB */
var ApiKeyTMDB string

/**
* Read config file
* @param string configFile: path to the config file
 */
func ReadConfigFile(configFile string) {
	/* Read config file */
	jsonFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		Errors.Fatal.HandlingErrors(err, true, Errors.ErrorConfigFileNotFound)
	}
	var config Config

	/* we unmarshal our byteArray which contains our
	jsonFile's content into 'config' which we defined above */
	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		Errors.Fatal.HandlingErrors(err, true, Errors.ErrorConfigFileUnreadable)
	}
	LogErrors = config.LogErrors
	PortRestAPI = config.PortRestAPI
	BackendServiceTMDB = config.BackendServiceTMDB
	ApiKeyTMDB = config.ApiKeyTMDB
}
