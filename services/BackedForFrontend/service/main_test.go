/**
* Config tests
* @author  Diana Lucia Serna Higuita
* Test funcion of all packages
 */
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	API "github.com/TMDBHydra/BackedForFrontend/pkg/api"
	Config "github.com/TMDBHydra/BackedForFrontend/pkg/config"
)

/*****************************************************************************
*  TEST FUNCTIONS CONFIG PACKAGE
*****************************************************************************/

/**
* Test config file
 */
func TestConfig(t *testing.T) {
	Config.ReadConfigFile("config/config.json")
	if len(Config.PortRestAPI) == 0 {
		t.Errorf("Error the config json with the variable PortRestAPI")
	}

	if Config.LogErrors != false && Config.LogErrors != true {
		t.Errorf("Error the config json with the variable LogErrors")
	}
}

/*****************************************************************************
*  TEST FUNCTIONS API PACKAGE
*****************************************************************************/

/**
* Test get tv series good request
* Test the endpoint GetTVSeries by a good formed request
* The request must have the parameter "query" in the body
 */
func TestGetTVSeriesGoodRequest(t *testing.T) {
	request, err := http.NewRequest("GET", "/api/frontend/tvserie?query=Modern", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(API.GetTVSeries)
	handler.ServeHTTP(response, request)
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

/**
* Test get tv series bad request
* Test the endpoint GetTVSeries by a bad request, in this case the
* error status StatusBadRequest must be returned
 */
func TestGetTVSeriesBadRequest(t *testing.T) {
	request, err := http.NewRequest("GET", "/api/frontend/tvserie?query1=Modern", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(API.GetTVSeries)
	handler.ServeHTTP(response, request)
	if status := response.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

/**
* Test get seasons good request
* Test the endpoint GetSeasons by a good formed request
* The request must have the parameter "tvSerieId" in the body
 */
func TestGetSeasonsGoodRequest(t *testing.T) {
	request, err := http.NewRequest("GET", "/api/frontend/seasons?tvSerieId=1421", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(API.GetSeasons)
	handler.ServeHTTP(response, request)
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

/**
* Test get seasons bad request
* Test the endpoint GetSeasons by a bad request, in this case the
* error status StatusBadRequest must be returned
 */
func TestGetSeasonsBadRequest(t *testing.T) {
	request, err := http.NewRequest("GET", "/api/frontend/seasons?tvSerieId1=1421", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(API.GetSeasons)
	handler.ServeHTTP(response, request)
	if status := response.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

//--

/**
* Test get episodes good request
* Test the endpoint GetEpisodes by a good formed request
* The request must have the parameter "tvSerieId" and "season" in the body
 */
func TestGetEpisodesGoodRequest(t *testing.T) {
	request, err := http.NewRequest("GET", "/api/frontend/episodes?tvSerieId=1421&season=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(API.GetEpisodes)
	handler.ServeHTTP(response, request)
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

/**
* Test get episodes bad request
* Test the endpoint GetEpisodes by a bad request, in this case the
* error status StatusBadRequest must be returned
 */
func TestGetEpisodesBadRequest(t *testing.T) {
	request, err := http.NewRequest("GET", "/api/frontend/episodes?tvSerieId1=1421", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(API.GetEpisodes)
	handler.ServeHTTP(response, request)
	if status := response.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
