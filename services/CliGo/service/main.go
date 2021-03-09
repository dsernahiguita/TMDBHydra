package main

import (
	"bufio"
	"fmt"
	"io"
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
* End
 */
func end() {
	fmt.Println("----------------------------------")
	fmt.Println("|  Thanks for using our services  |")
	fmt.Println("-----------------------------------")
}

/**
* Read string
* @param ioReader
* @return string value
* @return error err
 */
func readString(stdin io.Reader) (string, error) {
	reader := bufio.NewReader(stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	/* convert CRLF to LF */
	value = strings.Replace(value, "\n", "", -1)
	return value, nil
}

/**
* Ask tv serie
* ask the user to enter a tv series. It will prompt until the user has
* entered a non-empty value
* @return string tvSerieSearch
 */
func askTvSerie() string {
	var tvSerieSearch string
	for {
		fmt.Print("-> Please enter a tv serie: ")
		tvSerie, err := readString(os.Stdin)
		if err != nil {
			Errors.Error.HandlingErrors(err, true, Errors.ErrorEnterTVSerie)
		} else {
			if len(tvSerie) != 0 {
				tvSerieSearch = tvSerie
				break
			}
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
* @param Api.TVSeries results
* @return Api.TVSeries results
 */
func askPageAndDisplayResults(tvSerie string, totalPages int, results Api.TVSeries) Api.TVSeries {
	for {
		page := 1
		fmt.Print("-> If you want to see more results, please type the page number, otherwise press enter: " + string(page))
		nextPage, err := readString(os.Stdin)
		if err != nil {
			Errors.Error.HandlingErrors(err, true, Errors.ErrorGetNextPage)
		} else {
			if len(nextPage) == 0 {
				break
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
* @param io.Reader stdin
* @return int tvSerieId
 */
func selectTvSerie(results Api.TVSeries, stdin io.Reader) int {
	var tvSerieIdSearch int
	for {
		fmt.Print("-> Please enter the SERIE ID you wants to see: ")
		tvSerieId, err := readString(stdin)
		if err != nil {
			Errors.Error.HandlingErrors(err, true, Errors.ErrorSelectSerieId)
		}
		if len(tvSerieId) != 0 {
			tvSerieIdSearch, err = strconv.Atoi(tvSerieId)
			if err != nil {
				fmt.Println("****** Please enters valid serie Id ")
			}
			if tvSerieIdSearch > 0 {
				break
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
* Find season
* search if the seasonId are in the seasons array
* @param Api.TVSerieSeasons tvSerieSeasons
* @param int seasonIdSearch
* @return bool
 */
func findSeason(tvSerieSeasons Api.TVSerieSeasons, seasonIdSearch int) bool {
	foundSeasonId := false
	for i := range tvSerieSeasons.Seasons {
		if tvSerieSeasons.Seasons[i].SeasonNumber == seasonIdSearch {
			foundSeasonId = true
		}
	}
	return foundSeasonId
}

/**
* Select season
* @param int totalSeasons
* @param Api.TVSerieSeasons tvSerieSeasons
* @param o.Reader stdin
* @return int seasonId
 */
func selectSeason(totalSeasons int, tvSerieSeasons Api.TVSerieSeasons, stdin io.Reader) int {
	var seasonIdSearch int
	for {
		fmt.Print("-> Please enter the SEASON ID you wants to see: ")
		seasonId, err := readString(stdin)
		if err != nil {
			Errors.Error.HandlingErrors(err, true, Errors.ErrorSelectSeasonId)
		}
		if len(seasonId) != 0 {
			seasonIdSearch, err = strconv.Atoi(seasonId)
			if err != nil {
				fmt.Println("****** Please enters valid season Id ")
			} else {
				/* search if the seasonId are in the seasons array */
				if findSeason(tvSerieSeasons, seasonIdSearch) == true {
					break
				}
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
* @param Api.EpisodesSeason episodes
* @param o.Reader stdin
* @return int episodeIdSearch
 */
func selectEpisode(episodes Api.EpisodesSeason, stdin io.Reader) int {
	var episodeIdSearch int
	for {
		fmt.Print("-> Please enter the EPISODE ID you wants to see: ")
		episodeId, err := readString(stdin)
		if err != nil {
			Errors.Error.HandlingErrors(err, true, Errors.ErrorSelectEpisodeId)
		}
		if len(episodeId) != 0 {
			episodeIdSearch, err = strconv.Atoi(episodeId)
			if err != nil {
				fmt.Println("****** Please enters valid episode Id ")
			}
			if episodeIdSearch > 0 && episodeIdSearch <= len(episodes.Episodes) {
				break
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
	/* the config file is located in the same folder where the application
	is runnig config/config.json */
	Config.ReadConfigFile("config/config.json")
	var results Api.TVSeries
	var tvSerieSearch string
	welcome()

	for {
		/* 1. Ask for the tv serie */
		tvSerieSearch = askTvSerie()

		/* 2. Display the first page with the results */
		results = searchAndDisplayTVSeries(tvSerieSearch, 1)
		if len(results.Results) > 0 {
			break
		} else {
			fmt.Println("****** The search has no results")
		}
	}

	totalPages := results.TotalPages

	/* 3. If the results has severals pages, ask the user if he wants to see
	more pages */
	if totalPages > 1 {
		results = askPageAndDisplayResults(tvSerieSearch, totalPages, results)
	}

	/* 4. Ask the user to select a tvserie entering the SERIE ID */
	tvSerieId := selectTvSerie(results, os.Stdin)

	/* 5. Display the seasons of the selected tvserie */
	tvSerieSeasons := displaySeasons(tvSerieId)
	totalSeasons := tvSerieSeasons.NumberOfSeasons

	/* 6. Select season */
	seasonId := selectSeason(totalSeasons, tvSerieSeasons, os.Stdin)

	/* 7. Display the episodes of the selected seasonId */
	episodes := displayEpisodes(tvSerieId, seasonId, tvSerieSeasons)

	/* 8. Select episode */
	episodeId := selectEpisode(episodes, os.Stdin)

	/* 9. Display episode */
	displayEpisode(episodes, episodeId, tvSerieSeasons)

	/* 10. End */
	end()
}
