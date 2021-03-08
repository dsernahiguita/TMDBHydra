package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	Api "github.com/TMDBHydra/CliGo/pkg/api"
	Config "github.com/TMDBHydra/CliGo/pkg/config"
	Errors "github.com/TMDBHydra/CliGo/pkg/errors"
)

/**
* Welcome
 */
func welcome() {
	fmt.Println("---- Welcome to TMDB Hydra CLI ----")
	fmt.Println("|       Search TV Series          |")
	fmt.Println("-----------------------------------")
}

/**
* Ask tv serie
* ask the user to enter a tv series. It will prompt until the user has
* entered a non-empty value
* @return string tvSerieSearch
 */
func askTvSerie() string {
	reader := bufio.NewReader(os.Stdin)
	exit := false
	var tvSerieSearch string
	for exit == false {
		fmt.Print("-> Please enter a tv serie: ")
		tvSerie, err := reader.ReadString('\n')
		if err != nil {
			Errors.Error.HandlingErrors(err, true, Errors.ErrorEnterTVSerie)
		}
		/* convert CRLF to LF */
		tvSerie = strings.Replace(tvSerie, "\n", "", -1)
		if len(tvSerie) != 0 {
			tvSerieSearch = tvSerie
			exit = true
		}
	}
	return tvSerieSearch
}

/**
* Search and display tv series
* @param string tvSerie
* @param int page
* @return Api.TVSeries results
 */
func searchAndDisplayTVSeries(tvSerie string, page int) Api.TVSeries {
	/* Get the tvSeries using the backendForFrontend page 1 */
	var results Api.TVSeries
	var err error
	results, err = Api.GetTVSeries(tvSerie, page)
	if err != nil {
		Errors.Error.HandlingErrors(err, true, Errors.ErrorGetTVSerie)
		return results
	}
	/* Display the tvSeries results */
	fmt.Printf("\n%s\n", "RESULTS")
	fmt.Println("Total TV series found: " + strconv.Itoa(results.TotalResults))
	fmt.Println("Total Pages: " + strconv.Itoa(results.TotalPages))
	fmt.Println("Results page: " + strconv.Itoa(results.Page))
	fmt.Println("-----------------------------------")

	fmt.Printf("%v | %s | %s \n", "SERIE ID", "NAME", "ORIGINAL NAME")
	for i := range results.Results {
		p := results.Results[i]
		fmt.Printf("%v | %s | %s \n", i+1, p.Name, p.OriginalName)
	}
	return results
}

/**
* Ask page and display results
* @param string tvSerie
* @param int totalPages
 */
func askPageAndDisplayResults(tvSerie string, totalPages int, results Api.TVSeries) Api.TVSeries {
	exit := false
	for exit == false {
		page := 1
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("-> If you want to see more results, please type the page number, otherwise press enter: " + string(page))
		nextPage, err := reader.ReadString('\n')
		if err != nil {
			Errors.Error.HandlingErrors(err, true, Errors.ErrorGetNextPage)
		} else {
			nextPage = strings.Replace(nextPage, "\n", "", -1)
			if len(nextPage) == 0 {
				exit = true
			} else {
				page, err = strconv.Atoi(nextPage)
				if err != nil {
					fmt.Println("****** Please enters valid page between 1 and " + strconv.Itoa(totalPages))
				}
				if page > totalPages || page <= 0 {
					fmt.Println("****** Please enters a page between 1 and " + strconv.Itoa(totalPages))
				} else {
					results = searchAndDisplayTVSeries(tvSerie, page)
					totalPages = results.TotalPages
				}
			}
		}
	}
	return results
}

/**
* Select tv serie
* @param Api.TVSeries results
* @return int tvSerieId
 */
func selectTvSerie(results Api.TVSeries) int {
	reader := bufio.NewReader(os.Stdin)
	exit := false
	var tvSerieIdSearch int
	for exit == false {
		fmt.Print("-> Please enter the SERIE ID you wants to see: ")
		tvSerieId, err := reader.ReadString('\n')
		if err != nil {
			Errors.Error.HandlingErrors(err, true, Errors.ErrorSelectSerieId)
		}
		/* convert CRLF to LF */
		tvSerieId = strings.Replace(tvSerieId, "\n", "", -1)
		if len(tvSerieId) != 0 {
			tvSerieIdSearch, err = strconv.Atoi(tvSerieId)
			if err != nil {
				fmt.Println("****** Please enters valid serie Id ")
			}
			if tvSerieIdSearch > 0 {
				exit = true
			}
		}
	}
	/* return the tvSerieId located in the position tvSerieIdSearch-1 */
	return results.Results[tvSerieIdSearch-1].Id
}

/**
* Display seasons
* @param int tvSerieId
* @return Api.TVSerieSeasons results
 */
func displaySeasons(tvSerieId int) Api.TVSerieSeasons {
	var results Api.TVSerieSeasons
	results, err := Api.GetSeasons(tvSerieId)
	if err != nil {
		Errors.Error.HandlingErrors(err, true, Errors.ErrorGetSeasons)
		return results
	}

	/* Display the tvSeries results */
	fmt.Printf("\n%s\n", "RESULTS")
	fmt.Println("TV serie: " + results.Name)
	fmt.Println("Total seasons: " + strconv.Itoa(results.NumberOfSeasons))
	fmt.Println("-----------------------------------")

	fmt.Printf("%v | %s  \n", "SEASON ID", "NAME")
	for i := range results.Seasons {
		p := results.Seasons[i]
		fmt.Printf("%v | %s \n", p.SeasonNumber, p.Name)
	}

	return results
}

/**
* Select season
* @param int totalSeasons
* @param Api.TVSerieSeasons tvSerieSeasons
* @return int seasonId
 */
func selectSeason(totalSeasons int, tvSerieSeasons Api.TVSerieSeasons) int {
	reader := bufio.NewReader(os.Stdin)
	exit := false
	var seasonIdSearch int
	for exit == false {
		fmt.Print("-> Please enter the SEASON ID you wants to see: ")
		seasonId, err := reader.ReadString('\n')
		if err != nil {
			Errors.Error.HandlingErrors(err, true, Errors.ErrorSelectSeasonId)
		}
		/* convert CRLF to LF */
		seasonId = strings.Replace(seasonId, "\n", "", -1)
		if len(seasonId) != 0 {
			seasonIdSearch, err = strconv.Atoi(seasonId)
			if err != nil {
				fmt.Println("****** Please enters valid serie Id ")
			}
			/* search if the seasonId are int the seasons array */
			foundSeasonId := false
			for i := range tvSerieSeasons.Seasons {
				if tvSerieSeasons.Seasons[i].SeasonNumber == seasonIdSearch {
					foundSeasonId = true
				}
			}
			if foundSeasonId == true {
				exit = true
			}
		}
	}
	return seasonIdSearch
}

/**
* Display seasons
* @param int tvSerieId
* @param int season
* @param Api.TVSerieSeasons tvSerieSeasons
* @return Api.EpisodesSeason results
 */
func displayEpisodes(tvSerieId int, season int, tvSerieSeasons Api.TVSerieSeasons) Api.EpisodesSeason {
	var results Api.EpisodesSeason
	results, err := Api.GetEpisodes(tvSerieId, season)
	if err != nil {
		Errors.Error.HandlingErrors(err, true, Errors.ErrorGetEpisodes)
		return results
	}

	/* Display the tvSeries results */
	fmt.Printf("\n%s\n", "RESULTS")
	fmt.Println("TV serie: " + tvSerieSeasons.Name)
	fmt.Println("Season: " + strconv.Itoa(season))
	fmt.Println("Total episodes: " + strconv.Itoa(len(results.Episodes)))
	fmt.Println("-----------------------------------")

	fmt.Printf("%v | %s  \n", "EPISODE NUMBER", "NAME")
	for i := range results.Episodes {
		p := results.Episodes[i]
		fmt.Printf("%v | %s \n", p.EpisodeNumber, p.Name)
	}

	return results
}

/**
* Select episode
 */
func selectEpisode(episodes Api.EpisodesSeason) int {
	reader := bufio.NewReader(os.Stdin)
	exit := false
	var episodeIdSearch int
	for exit == false {
		fmt.Print("-> Please enter the EPISODE ID you wants to see: ")
		episodeId, err := reader.ReadString('\n')
		if err != nil {
			Errors.Error.HandlingErrors(err, true, Errors.ErrorSelectEpisodeId)
		}
		/* convert CRLF to LF */
		episodeId = strings.Replace(episodeId, "\n", "", -1)
		if len(episodeId) != 0 {
			episodeIdSearch, err = strconv.Atoi(episodeId)
			if err != nil {
				fmt.Println("****** Please enters valid episode Id ")
			}
			if episodeIdSearch > 0 {
				exit = true
			}
		}
	}

	return episodeIdSearch - 1
}

/**
* Display episode
* @param  Api.EpisodesSeason episodes
* @param int episodeId
 */
func displayEpisode(episodes Api.EpisodesSeason, episodeId int, tvSerieSeasons Api.TVSerieSeasons) {
	fmt.Println("")
	fmt.Println("-----------------------------------")
	fmt.Printf("%s\n", "Results")
	fmt.Println("TV serie: " + tvSerieSeasons.Name)
	fmt.Println("Season: " + strconv.Itoa(episodes.SeasonNumber))
	fmt.Println("episode: " + strconv.Itoa(episodeId+1))
	fmt.Println("Name: " + episodes.Episodes[episodeId].Name)
	fmt.Println("Overview: " + episodes.Episodes[episodeId].Overview)
	fmt.Println("-----------------------------------")
}

func main() {
	/* Read config file, this file contains information such local storages */
	/* the config file is located in the same folder where the application
	is runnig config/config.json */
	Config.ReadConfigFile("config/config.json")
	var results Api.TVSeries
	welcome()

	/* 1. Ask for the tv serie */
	tvSerieSearch := askTvSerie()

	/* 2. Display the first page with the results */
	results = searchAndDisplayTVSeries(tvSerieSearch, 1)
	totalPages := results.TotalPages

	/* 3. If the results has severals pages, ask the user if he wants to see
	more pages */
	if totalPages > 1 {
		results = askPageAndDisplayResults(tvSerieSearch, totalPages, results)
	}

	/* 4. Ask the user to select a tvserie entering the SERIE ID */
	tvSerieId := selectTvSerie(results)

	/* 5. Display the seasons of the selected tvserie */
	tvSerieSeasons := displaySeasons(tvSerieId)
	totalSeasons := tvSerieSeasons.NumberOfSeasons

	/* 6. Select season */
	seasonId := selectSeason(totalSeasons, tvSerieSeasons)

	/* 7. Display the episodes of the selected seasonId */
	episodes := displayEpisodes(tvSerieId, seasonId, tvSerieSeasons)

	/* 8. Select episode */
	episodeId := selectEpisode(episodes)

	/* 9. Display episode */
	displayEpisode(episodes, episodeId, tvSerieSeasons)
}
