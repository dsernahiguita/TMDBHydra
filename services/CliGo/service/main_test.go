/**
* Config tests
* @author  Diana Lucia Serna Higuita
* Test funcion of all packages
 */
package main

import (
	"bytes"
	"testing"

	Api "github.com/TMDBHydra/CliGo/pkg/api"
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

/**
* Test function readString
 */
func TestReadString(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("TestSerie\n"))

	result, _ := readString(&stdin)
	if result != "TestSerie" {
		t.Errorf("Function readString error: result: got %v want %v",
			result, "TestSerie")
	}
}

/**
* Test function selectTvSerie
* This function must return the tvSerie.Id
 */
func TestSelectTvSerie(t *testing.T) {
	var stdin bytes.Buffer
	var results Api.TVSeries
	var tvSerie Api.ResultsTVSeries

	tvSerie.Id = 50
	tvSerie.Name = "Test"
	tvSerie.OriginalName = "TestOriginalName"
	tvSerie.Overview = ""

	results.Page = 1
	results.TotalPages = 10
	results.TotalResults = 99
	results.Results = append(results.Results, tvSerie)

	stdin.Write([]byte("1\n"))
	result := selectTvSerie(results, &stdin)

	if result != tvSerie.Id {
		t.Errorf("Function selectTvSerie error: result: got %v want %v",
			result, tvSerie.Id)
	}
}

/**
* Test function find season
* This function must return true if the season exists in the array seasons
 */
func TestFindSeason(t *testing.T) {
	var serieSeasons Api.TVSerieSeasons
	var season Api.Season

	serieSeasons.Name = "Moderne Family"
	serieSeasons.NumberOfSeasons = 10
	serieSeasons.NumberOfEpisodes = 100

	season.Id = 1
	season.Name = "Season1"
	season.Overview = ""
	season.SeasonNumber = 1

	serieSeasons.Seasons = append(serieSeasons.Seasons, season)

	/* Test existing season */
	foundSeason := findSeason(serieSeasons, 1)
	if foundSeason != true {
		t.Errorf("Function findSeason error: result: got %v want %v",
			foundSeason, "true")
	}
	/* Test not existing season */
	foundSeason = findSeason(serieSeasons, 10)
	if foundSeason != false {
		t.Errorf("Function findSeason error: result: got %v want %v",
			foundSeason, "false")
	}
}

/**
* Test function select seasons
* This function must return the season Id selected
 */
func TestSelectSeason(t *testing.T) {
	var stdin bytes.Buffer
	var serieSeasons Api.TVSerieSeasons
	var season Api.Season

	serieSeasons.Name = "Moderne Family"
	serieSeasons.NumberOfSeasons = 10
	serieSeasons.NumberOfEpisodes = 100

	season.Id = 1
	season.Name = "Season1"
	season.Overview = ""
	season.SeasonNumber = 1
	serieSeasons.Seasons = append(serieSeasons.Seasons, season)

	stdin.Write([]byte("1\n"))
	result := selectSeason(1, serieSeasons, &stdin)

	if result != 1 {
		t.Errorf("Function selectSeason error: result: got %v want %v",
			result, "1")
	}
}

/**
* Test function select episode
* This function returns the index where the selected episode is located.
 */
func TestSelectEpisode(t *testing.T) {
	var stdin bytes.Buffer
	var episodeSeason Api.EpisodesSeason
	var episode Api.Episode

	episodeSeason.Id = 1
	episodeSeason.Name = "Season 1"
	episodeSeason.Overview = ""
	episodeSeason.SeasonNumber = 1

	episode.Id = 1
	episode.Name = "Episode 1"
	episode.Overview = ""
	episode.EpisodeNumber = 1

	episodeSeason.Episodes = append(episodeSeason.Episodes, episode)

	stdin.Write([]byte("1\n"))
	result := selectEpisode(episodeSeason, &stdin)

	if result != 0 {
		t.Errorf("Function selectEpisode error: result: got %v want %v",
			result, "1")
	}
}
